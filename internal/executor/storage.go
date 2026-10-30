package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"strconv"
	"strings"
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

func (s *Storage) SetCode(code int) {
	strCode := strconv.Itoa(code)
	s.vars["code"] = strCode
}

func (s *Storage) SetStatus(status string) {
	s.vars["status"] = status
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
	switch {
	case strings.HasPrefix(what, "json"):
		return s.getJson(what)
	case strings.HasPrefix(what, "headers"):
		return s.getHeader(what)
	default:
		return s.getVar(what)
	}
}

func (s *Storage) getJson(what string) (string, error) {
	if !strings.HasPrefix(what, "json") {
		return "", errors.New("what is not json")
	}
	var bytes []byte
	var err error
	if what == "json" {
		bytes, err = json.Marshal(s.json)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	parts := strings.Split(what, ".")
	cur := s.json
	for _, key := range parts[1 : len(parts)-1] {
		something, ok := s.json[key]
		if !ok {
			return "", fmt.Errorf("%s not found", what)
		}
		asMap, ok := something.(map[string]any)
		if !ok {
			return "", fmt.Errorf("%s not found", what)
		}
		cur = asMap
	}
	lastKey := parts[len(parts)-1]
	el, ok := cur[lastKey]
	if !ok {
		return "", fmt.Errorf("%s not found", what)
	}
	bytes, err = json.Marshal(el)
	if err != nil {
		return "", err
	}
	value := string(bytes)
	value = strings.Trim(value, "\"")
	return value, nil
}

func (s *Storage) getHeader(what string) (string, error) {
	if !strings.HasPrefix(what, "headers") {
		return "", errors.New("what is not headers")
	}
	if what == "headers" {
		result := make([]string, 0, len(s.headers))
		for key, headers := range s.headers {
			values := strings.Join(headers, "; ")
			row := fmt.Sprintf("%s: %s", key, values)
			result = append(result, row)
		}
		// 0x1E - RECORD SEPARATOR
		return strings.Join(result, "0x1E"), nil
	}
	parts := strings.Split(what, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("%s not found", what)
	}
	key := parts[1]
	values, ok := s.headers[key]
	if !ok {
		return "", fmt.Errorf("%s not found", what)
	}
	return fmt.Sprintf("%s: %s", key, strings.Join(values, "; ")), nil
}

func (s *Storage) getVar(what string) (string, error) {
	if what == "vars" {
		val := fmt.Sprintf("%v", s.vars)
		return val, nil
	}
	val, ok := s.vars[what]
	if !ok {
		return "", fmt.Errorf("%s not found", what)
	}
	return val, nil
}
