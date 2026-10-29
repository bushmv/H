package directives

type Capture struct {
	Name string
	What string
	Source
	substitutor
}

func (c *Capture) Substitute(env map[string]string) error {
	nameValue, err := c.substitute(c.Name, env)
	if err != nil {
		return err
	}
	whatValue, err := c.substitute(c.What, env)
	if err != nil {
		return err
	}
	c.Name = nameValue
	c.What = whatValue
	return nil
}

func NewCapture(name string, what string, source Source) *Capture {
	return &Capture{
		Name:   name,
		What:   what,
		Source: source,
	}
}
