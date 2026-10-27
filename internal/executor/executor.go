package executor

import "github.com/bushmv/H/internal/test"

type Executor struct {
	tp *TemplateProcessor
}

func NewExecutor(tp *TemplateProcessor) *Executor {
	return &Executor{
		tp: tp,
	}
}
func (e *Executor) Run(tests []*test.Test, storage *Storage) ([]*test.TestResult, error) {
	return nil, nil
}
