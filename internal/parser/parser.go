package parser

import (
	"io"

	"github.com/bushmv/H/internal/test"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(reader io.Reader, sourceName string) ([]*test.Test, error) {
	return nil, nil
}
