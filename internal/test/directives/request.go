package directives

type Request struct {
	Method string
	URL    string
	substitutor
	Source
}

func (r *Request) Substitute(env map[string]string) error {
	methodValue, err := r.substitute(r.Method, env)
	if err != nil {
		return err
	}
	r.Method = methodValue
	URLValue, err := r.substitute(r.URL, env)
	if err != nil {
		return err
	}
	r.URL = URLValue
	return nil
}

func NewRequest(method string, URL string, source Source) *Request {
	return &Request{
		Method: method,
		URL:    URL,
		Source: source,
	}
}
