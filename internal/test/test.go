package test

import "github.com/bushmv/H/internal/test/directives"

type Test struct {
	Request        *directives.Request
	Body           *directives.Body
	Prints         []*directives.Print
	Headers        []*directives.Header
	Captures       []*directives.Capture
	Asserts        []directives.AssertableSubstitutable
	Queries        []*directives.Query
	SubstituteList []directives.Substituteable
	Passed         bool
}

func NewTest(request *directives.Request) *Test {
	return &Test{
		Request:        request,
		Body:           nil,
		Prints:         make([]*directives.Print, 0),
		Headers:        make([]*directives.Header, 0),
		Captures:       make([]*directives.Capture, 0),
		Asserts:        make([]directives.AssertableSubstitutable, 0),
		Queries:        make([]*directives.Query, 0),
		SubstituteList: make([]directives.Substituteable, 0),
		Passed:         true,
	}
}

func (t *Test) AddPrint(print *directives.Print) {
	t.Prints = append(t.Prints, print)
}

func (t *Test) AddHeader(header *directives.Header) {
	t.Headers = append(t.Headers, header)
}

func (t *Test) AddAssert(assert directives.AssertableSubstitutable) {
	t.Asserts = append(t.Asserts, assert)
}

func (t *Test) AddCapture(capture *directives.Capture) {
	t.Captures = append(t.Captures, capture)
}

func (t *Test) AddSubstitutable(substituteable directives.Substituteable) {
	t.SubstituteList = append(t.SubstituteList, substituteable)
}
