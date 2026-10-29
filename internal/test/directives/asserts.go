package directives

import "fmt"

const (
	AssertResultUnknown = iota
	AssertResultPassed
	AssertResultFailed
)

type Assert struct {
	Passed int
	Source
	substitutor
}

type EqualAssert struct {
	What     string
	Expected string
	Value    string
	Assert
}

func NewEqualAssert(what string, expected string, source Source) *EqualAssert {
	return &EqualAssert{
		What:     what,
		Expected: expected,
		Assert: Assert{
			Source: source,
			Passed: AssertResultUnknown,
		},
	}
}

func (a *EqualAssert) Substitute(env map[string]string) error {
	whatValue, err := a.substitute(a.What, env)
	if err != nil {
		return err
	}
	expectedValue, err := a.substitute(a.Expected, env)
	if err != nil {
		return err
	}
	a.What = whatValue
	a.Expected = expectedValue
	return nil
}

func (a *EqualAssert) Check() {
	if a.Value == a.Expected {
		a.Passed = AssertResultPassed
	} else {
		a.Passed = AssertResultFailed
	}
}

func (a *EqualAssert) Message() string {
	switch a.Passed {
	case AssertResultPassed:
		return fmt.Sprintf("%s == %s", a.What, a.Expected)
	case AssertResultFailed:
		return fmt.Sprintf("%s != %s", a.What, a.Expected)
	case AssertResultUnknown:
		return ""
	}
	panic("unreachable")
}

func (a *EqualAssert) Details() string {
	switch a.Passed {
	case AssertResultPassed:
		return ""
	case AssertResultFailed:
		return fmt.Sprintf("actual: %s", a.Value)
	case AssertResultUnknown:
		return ""
	}
	panic("unreachable")
}
