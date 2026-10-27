package executor

type TemplateProcessor struct {
}

func NewTemplateProcessor() *TemplateProcessor {
	return &TemplateProcessor{}
}

func (t *TemplateProcessor) Substitute(template string, vars map[string]string) (string, error) {
	return "", nil
}
