package directives

type Print struct {
	What string
	Source
	substitutor
}

func (p *Print) Substitute(env map[string]string) error {
	value, err := p.substitute(p.What, env)
	if err != nil {
		return err
	}
	p.What = value
	return nil
}

func NewPrint(what string, source Source) *Print {
	return &Print{
		What:   what,
		Source: source,
	}
}
