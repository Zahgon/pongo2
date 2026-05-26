package pongo2

type NodeWrapper struct {
	Endtag string
	nodes  []INode
}

func (wrapper *NodeWrapper) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}
