package pongo2

type nodeHTML struct {
	token     *Token
	trimLeft  bool
	trimRight bool
}

func (n *nodeHTML) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}
