package pongo2

// tagCommentNode represents the {% comment %} tag.
//
// The comment tag ignores everything between {% comment %} and {% endcomment %}.
// This is useful for commenting out code or adding notes that should not appear
// in the rendered output. Unlike HTML comments, template comments are completely
// removed from the output.
//
// Usage:
//
//	{% comment %}
//	    This text will not appear in the output.
//	    You can write multiple lines here.
//	    {{ variables }} and {% tags %} are also ignored.
//	{% endcomment %}
//
// Example:
//
//	<p>Visible content</p>
//	{% comment %}
//	    TODO: Add more features here later
//	    {{ debug_var }}
//	{% endcomment %}
//	<p>More visible content</p>
//
// Output:
//
//	<p>Visible content</p>
//	<p>More visible content</p>
type tagCommentNode struct{}

// Execute is a no-op for comment nodes. The content between {% comment %}
// and {% endcomment %} is completely ignored and never rendered.
func (node *tagCommentNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"

	// tagCommentParser parses the {% comment %} tag. It skips all content
	// until {% endcomment %} and does not accept any arguments.
	return nil
}

func tagCommentParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// TODO: Process the endtag's arguments (see django 'comment'-tag documentation)

func init() {
	mustRegisterTag("comment", tagCommentParser)
}
