package test

import "bytes"

type Body struct {
	LineNumber  int
	ContentType string
	Content     *bytes.Buffer
}

type Print struct {
	LineNumber int
	What       string
}

type Header struct {
	LineNumber int
	Key        string
	Value      string
}

type Capture struct {
	LineNumber int
	Name       string
	What       string
}

type Assert struct {
	LineNumber int
	What       string
	Expected   string
}

type Query struct {
	LineNumber int
	Key        string
	Value      string
}
type Test struct {
	Method   string
	URL      string
	Body     *Body
	Prints   []*Print
	Headers  []*Header
	Captures []*Capture
	Asserts  []*Assert
	Queries  []*Query
}

func NewTest(method string, url string) *Test {
	return &Test{
		Method:   method,
		URL:      url,
		Body:     nil,
		Prints:   make([]*Print, 0),
		Headers:  make([]*Header, 0),
		Captures: make([]*Capture, 0),
		Asserts:  make([]*Assert, 0),
		Queries:  make([]*Query, 0),
	}
}

func (t *Test) AddPrint(print *Print) {
	t.Prints = append(t.Prints, print)
}

func (t *Test) AddHeader(header *Header) {
	t.Headers = append(t.Headers, header)
}

func (t *Test) AddAssert(assert *Assert) {
	t.Asserts = append(t.Asserts, assert)
}

func (t *Test) AddCapture(capture *Capture) {
	t.Captures = append(t.Captures, capture)
}
