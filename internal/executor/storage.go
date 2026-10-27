package executor

import (
	"maps"
	"net/http"
)

type Storage struct {
	vars    map[string]string
	headers map[string][]string
	json    map[string]any
}

func NewStorage() *Storage {
	return &Storage{
		vars:    make(map[string]string),
		headers: make(map[string][]string),
		json:    make(map[string]any),
	}
}

func (s *Storage) SetVariable(key string, value string) {
	s.vars[key] = value
}

func (s *Storage) SetHeaders(headers http.Header) {
	clear(s.headers)
	maps.Copy(s.headers, headers)
}

func (s *Storage) SetJson(data map[string]any) {
	clear(s.json)
	s.json = data
}

func (s *Storage) Get(what string) (string, error) {
	return "", nil
}
