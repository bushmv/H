package view

import (
	"fmt"

	"github.com/bushmv/H/internal/test"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

func (r *View) Show(results []*test.TestResult) {
	total, passed := 0, 0
	n := len(results)
	for i, result := range results {
		showRequestInfo(i+1, n, result)
		showCapture(result.Captures)
		showPrints(result.Prints)
		showAsserts(result.AssertResults)
		fmt.Println()
		total++
		if result.Passed {
			passed++
		}
	}
	showSummary(passed, total)
}

func showRequestInfo(i int, n int, result *test.TestResult) {
	color := green
	if !result.Passed {
		color = red
	}
	fmt.Printf("%s[%d/%d] %s %s %s(%s)%s\n", color, i, n, result.Method, result.URL, gray, result.Status, reset)
}

func showAsserts(assertResults []*test.AssertResult) {
	n := len(assertResults)
	for i, assert := range assertResults {
		if assert.Passed {
			fmt.Printf("    %s! [%d/%d] PASSED %s%s\n", green, i+1, n, assert.Message, reset)
		} else {
			fmt.Printf("    %s! [%d/%d] FAILED %s %s(%s)%s\n", red, i+1, n, assert.Message, gray, assert.Details, reset)
		}
	}
}

func showPrints(prints []string) {
	for _, print := range prints {
		fmt.Printf("    > %s\n", print)
	}
}

func showCapture(captures []string) {
	for _, capture := range captures {
		fmt.Printf("    %s& %s%s\n", gray, capture, reset)
	}
}

func showSummary(passed int, total int) {
	fmt.Printf("\ntotal: %d, passed: %d [%.2f%%]\n", total, passed, float64(passed)*100.0/float64(total))
}
