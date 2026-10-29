package test

import (
	"github.com/bushmv/H/internal/test/directives"
)

type TestResult struct {
	Method        string
	URL           string
	Status        string
	Prints        []string
	AssertResults []*AssertResult
	Captures      []string
	Passed        bool
}

func NewTestResult(method string, URL string, status string) *TestResult {
	return &TestResult{
		Method:        method,
		URL:           URL,
		Status:        status,
		Prints:        make([]string, 0),
		AssertResults: make([]*AssertResult, 0),
		Captures:      make([]string, 0),
		Passed:        false,
	}
}

type AssertResult struct {
	Passed  bool
	Message string
	Details string
	directives.Source
}

func NewAssertResult(assert directives.Assertable) *AssertResult {
	switch v := assert.(type) {
	case *directives.EqualAssert:
		return &AssertResult{
			Passed:  v.Passed == directives.AssertResultPassed,
			Message: v.Message(),
			Details: v.Details(),
			Source:  v.Source,
		}
	}
	panic("unreachable")
}
