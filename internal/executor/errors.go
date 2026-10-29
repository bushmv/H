package executor

import (
	"fmt"

	"github.com/bushmv/H/internal/test/directives"
)

type ExecutionError struct {
	Cause error
	directives.Source
}

func NewExecutionError(cause error, source directives.Source) *ExecutionError {
	return &ExecutionError{
		Cause:  cause,
		Source: source,
	}
}

func (e *ExecutionError) Error() string {
	return fmt.Sprintf("execution error: %v (%s:%d: %s)\n", e.Cause, e.Filename, e.LineNumber, e.Raw)
}
