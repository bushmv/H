package directives

type Query struct {
	Key   string
	Value string
	Source
	substitutor
}

func (q *Query) Substitute(env map[string]string) error {
	keyValue, err := q.substitute(q.Key, env)
	if err != nil {
		return err
	}
	valueValue, err := q.substitute(q.Value, env)
	if err != nil {
		return err
	}
	q.Key = keyValue
	q.Value = valueValue
	return nil
}
