package view

import "github.com/bushmv/H/internal/test"

type View struct {
}

func NewView() *View {
	return &View{}
}

func (r *View) Show(testResults []*test.TestResult) {
	total, passed := 0, 0
	// n := len(testResults)
	for _, result := range testResults {
		// print prints
		// check asserts
		// capture captures
		total++
		if result.Passed {
			passed++
		}
	}
}
