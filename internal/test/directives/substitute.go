package directives

import (
	"fmt"
	"strings"
)

type substitutor struct{}

func (s substitutor) substitute(template string, vars map[string]string) (string, error) {
	if !strings.Contains(template, "{{") {
		return template, nil
	}
	var sb strings.Builder
	sb.Grow(len(template))
	for {
		start := strings.Index(template, "{{")
		if start == -1 {
			sb.WriteString(template)
			break
		}
		sb.WriteString(template[:start])
		end := strings.Index(template[start+2:], "}}")
		if end == -1 {
			return "", fmt.Errorf("No close brackets for open brackets")
		}
		end += start + 2
		name := template[start+2 : end]
		value, ok := vars[name]
		if !ok {
			return "", fmt.Errorf("%s not found", name)
		}
		sb.WriteString(value)
		template = template[end+2:]
	}
	return sb.String(), nil
}
