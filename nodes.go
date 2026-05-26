package pongo2

// The root document
type nodeDocument struct {
	Nodes []INode
}

func (doc *nodeDocument) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}
