package directives

type Header struct {
	Key   string
	Value string
	Source
	substitutor
}

func NewHeader(key string, value string, source Source) *Header {
	return &Header{
		Key:    key,
		Value:  value,
		Source: source,
	}
}

func (h *Header) Substitute(env map[string]string) error {
	keyValue, err := h.substitute(h.Key, env)
	if err != nil {
		return err
	}
	valueValue, err := h.substitute(h.Value, env)
	if err != nil {
		return err
	}
	h.Key = keyValue
	h.Value = valueValue
	return nil
}
