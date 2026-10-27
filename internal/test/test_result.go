package test

type AssertResult struct {
	What     string
	Expected string
	Actual   string
	Passed   bool
	Details  string
}

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
