package directives

import (
	"bytes"
)

type Body struct {
	ContentType string
	Content     *bytes.Buffer
	Source
	substitutor
}

func (b *Body) Substitute(env map[string]string) error {
	if b.ContentType == "json" {
		value, err := b.substitute(b.Content.String(), env)
		if err != nil {
			return err
		}
		b.Content = bytes.NewBufferString(value)
	}
	return nil
}
