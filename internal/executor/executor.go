package executor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/bushmv/H/internal/test"
	"github.com/bushmv/H/internal/test/directives"
)

type Executor struct {
}

func NewExecutor() *Executor {
	return &Executor{}
}
func (e *Executor) Run(tests []*test.Test) ([]*test.TestResult, error) {
	storage := NewStorage()
	client := http.Client{}
	results := make([]*test.TestResult, 0, len(tests))
	for _, t := range tests {
		if err := Substitute(t, storage.vars); err != nil {
			return nil, err
		}
		req, err := HttpRequestFromTest(t)
		if err != nil {
			return nil, err
		}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		result, err := handleResponse(res, t, storage)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
		storage.SetJson(nil)
	}
	return results, nil
}

func Substitute(t *test.Test, env map[string]string) error {
	for _, s := range t.SubstituteList {
		if err := s.Substitute(env); err != nil {
			return err
		}
	}
	return nil
}

func HttpRequestFromTest(test *test.Test) (*http.Request, error) {
	URL, err := url.Parse(test.Request.URL)
	if err != nil {
		err := fmt.Errorf("invalid url")
		return nil, NewExecutionError(err, test.Request.Source)
	}
	req := &http.Request{
		URL:    URL,
		Method: test.Request.Method,
		Header: make(http.Header, 0),
	}
	if test.Body != nil {
		req.Body = io.NopCloser(test.Body.Content)
	}
	for _, header := range test.Headers {
		req.Header.Add(header.Key, header.Value)
	}
	return req, nil
}

func handleResponse(res *http.Response, t *test.Test, storage *Storage) (*test.TestResult, error) {
	saveResponse(res, storage)
	result := test.NewTestResult(t.Request.Method, t.Request.URL, res.Status)
	prints, err := addPrintsToResult(t, storage)
	if err != nil {
		return nil, err
	}
	result.Prints = prints
	assertResults, err := CheckAsserts(t, storage)
	if err != nil {
		return nil, err
	}
	result.Passed = verdict(assertResults)
	result.AssertResults = assertResults
	captures, err := CaptureVariables(t, storage)
	if err != nil {
		return nil, err
	}
	result.Captures = captures
	return result, nil
}

func saveResponse(res *http.Response, storage *Storage) error {
	storage.SetCode(res.StatusCode)
	storage.SetStatus(res.Status)
	storage.SetHeaders(res.Header)
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if json.Valid(bodyBytes) {
		var jsonData map[string]any
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&jsonData); err != nil {
			return err
		}
		storage.SetJson(jsonData)
	}
	return nil
}

func addPrintsToResult(t *test.Test, storage *Storage) ([]string, error) {
	prints := make([]string, 0, len(t.Prints))
	for _, print := range t.Prints {
		what := print.What
		val, err := storage.Get(what)
		if err != nil {
			return nil, NewExecutionError(err, print.Source)
		}
		out := fmt.Sprintf("%s: %s", what, val)
		prints = append(prints, out)
	}
	return prints, nil
}

func CheckAsserts(t *test.Test, storage *Storage) ([]*test.AssertResult, error) {
	results := make([]*test.AssertResult, 0, len(t.Asserts))
	for _, assert := range t.Asserts {
		switch a := assert.(type) {
		case *directives.EqualAssert:
			value, err := storage.Get(a.What)
			if err != nil {
				return nil, NewExecutionError(err, a.Source)
			}
			a.Value = value
			a.Check()
			t.Passed = t.Passed && a.Passed == directives.AssertResultPassed
			result := test.NewAssertResult(a)
			results = append(results, result)
		}
	}
	return results, nil
}

func verdict(assertResults []*test.AssertResult) bool {
	for _, ar := range assertResults {
		if !ar.Passed {
			return false
		}
	}
	return true
}

func CaptureVariables(t *test.Test, storage *Storage) ([]string, error) {
	captures := make([]string, 0, len(t.Captures))
	for _, capture := range t.Captures {
		value, err := storage.Get(capture.What)
		if err != nil {
			return nil, NewExecutionError(err, capture.Source)
		}
		storage.SetVariable(capture.Name, value)
		cap := fmt.Sprintf("%s = %s", capture.Name, value)
		captures = append(captures, cap)
	}
	return captures, nil
}
