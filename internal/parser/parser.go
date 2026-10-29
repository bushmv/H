package parser

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/bushmv/H/internal/test"
	"github.com/bushmv/H/internal/test/directives"
)

type Parser struct {
	tests          []*test.Test
	filename       string
	n              int
	line           string
	substituteable bool
	raw            string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) prepare(filename string) {
	p.tests = make([]*test.Test, 0)
	p.filename = filename
	p.n = 1
}

func (p *Parser) clear() {
	p.tests = nil
	p.filename = ""
	p.n = 1
}

func (p *Parser) Parse(reader io.Reader, filename string) ([]*test.Test, error) {
	p.prepare(filename)
	defer func() {
		p.clear()
	}()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			p.n++
			continue
		}
		p.line = line
		if err := p.parseLine(); err != nil {
			return nil, err
		}
		p.n++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return p.tests, nil
}

func (p *Parser) parseLine() error {
	p.raw = strings.TrimSpace(p.line)
	prefix, rest, ok := strings.Cut(p.line, " ")
	if !ok {
		return errors.New("no command specified")
	}
	var err error = nil
	p.substituteable = strings.Contains(rest, "{{")
	switch prefix {
	case "GET", "POST":
		err = p.handleRequest(prefix, rest)
	case ">":
		err = p.handlePrint(rest)
	case "+":
		err = p.handleHeader(rest)
	case "!":
		err = p.handleAssert(rest)
	case "{}":
		err = p.handleJsonBody(rest)
	case "&":
		err = p.handleCapture(rest)
	case "@":
		err = p.handleAuth(rest)
	default:
		return NewParsingError(fmt.Errorf("unknown prefix: %s", prefix), p.Source())
	}
	return err
}

func (p *Parser) handleRequest(method string, rest string) error {
	URL := strings.TrimSpace(rest)
	request := directives.NewRequest(method, URL, p.Source())
	t := test.NewTest(request)
	p.tests = append(p.tests, t)
	if p.substituteable {
		p.current().AddSubstitutable(request)
	}
	return nil
}

func (p *Parser) handlePrint(rest string) error {
	what := strings.TrimSpace(rest)
	print := directives.NewPrint(what, p.Source())
	p.current().AddPrint(print)
	if p.substituteable {
		p.current().AddSubstitutable(print)
	}
	return nil
}

func (p *Parser) handleHeader(rest string) error {
	parts := strings.Split(rest, ":")
	if len(parts) != 2 {
		err := fmt.Errorf("invalid format for header, use: <key>: <value>")
		return NewParsingError(err, p.Source())
	}
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	header := directives.NewHeader(key, value, p.Source())
	p.current().AddHeader(header)
	if p.substituteable {
		p.current().AddSubstitutable(header)
	}
	return nil

}

func (p *Parser) handleAssert(rest string) error {
	var assert directives.AssertableSubstitutable
	var err error
	switch {
	case strings.Contains(rest, " == "):
		what, expected, _ := strings.Cut(rest, " == ")
		what = strings.TrimSpace(what)
		expected = strings.TrimSpace(expected)
		assert = directives.NewEqualAssert(what, expected, p.Source())
	case strings.Contains(rest, " != "):
		// not equal assert
	case strings.Contains(rest, " ~= "):
		// regexp equal assert
	case strings.Contains(rest, " in "):
		// in assert
	case strings.Contains(rest, " in range "):
		// in range assert
	default:
		err = NewParsingError(fmt.Errorf("unknown assert type"), p.Source())
	}
	if err == nil {
		p.current().AddAssert(assert)
		if p.substituteable {
			p.current().AddSubstitutable(assert)
		}
	}
	return err
}

func (p *Parser) handleJsonBody(rest string) error {
	if !json.Valid([]byte(rest)) {
		return NewParsingError(fmt.Errorf("json not valid"), p.Source())
	}
	p.current().Body = &directives.Body{
		ContentType: "json",
		Content:     bytes.NewBufferString(rest),
	}
	header := &directives.Header{
		Key:   "Content-Type",
		Value: "application-json",
	}
	p.current().AddHeader(header)
	return nil
}

func (p *Parser) handleCapture(rest string) error {
	name, what, ok := strings.Cut(rest, " = ")
	if !ok {
		err := fmt.Errorf("invalid format for capture, use: <name> = <object>")
		return NewParsingError(err, p.Source())
	}
	capture := directives.NewCapture(name, what, p.Source())
	p.current().AddCapture(capture)
	return nil
}

func (p *Parser) handleAuth(rest string) error {
	scheme, credentials, found := strings.Cut(rest, " ")
	if !found {
		err := fmt.Errorf("invalid format for auth, use: @ SCHEMA CREDENTIALS")
		return NewParsingError(err, p.Source())
	}
	switch strings.ToLower(scheme) {
	case "basic":
		credentials = strings.TrimSpace(credentials)
		if !strings.Contains(credentials, ":") {
			err := fmt.Errorf("invalid format for basic auth, use: @ basic login:password")
			return NewParsingError(err, p.Source())
		}
		value := "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials))
		header := directives.NewHeader("Authorization", value, p.Source())
		p.current().AddHeader(header)
	default:
		err := fmt.Errorf("unknown auth type")
		return NewParsingError(err, p.Source())
	}
	return nil
}

func (p *Parser) current() *test.Test {
	if len(p.tests) <= 0 {
		panic("no test was added")
	}
	return p.tests[len(p.tests)-1]
}

func (p *Parser) Source() directives.Source {
	return directives.Source{
		Raw:        p.raw,
		LineNumber: p.n,
		Filename:   p.filename,
	}
}
