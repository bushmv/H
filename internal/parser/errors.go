package parser

import (
	"fmt"

	"github.com/bushmv/H/internal/test/directives"
)

type ParsingError struct {
	Cause error
	directives.Source
}

func NewParsingError(cause error, source directives.Source) *ParsingError {
	return &ParsingError{
		Cause:  cause,
		Source: source,
	}
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("parsing error: %v\n    %s:%d: %s\n", e.Cause, e.Filename, e.LineNumber, e.Raw)
}
