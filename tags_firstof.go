package pongo2

// tagFirstofNode represents the {% firstof %} tag.
//
// Django difference: Django supports {% firstof var1 var2 as name %} to store
// the result in a variable. This is not yet supported in pongo2.
//
// The firstof tag outputs the first variable that is "true" (not empty, not zero,
// not nil, not false). If all variables are false, nothing is output.
// This is useful for displaying fallback values.
//
// Usage:
//
//	{% firstof var1 var2 var3 %}
//
// Example with fallback values:
//
//	{% firstof user.nickname user.username "Anonymous" %}
//
// If user.nickname is "Johnny", output: "Johnny"
// If user.nickname is empty but username is "john_doe", output: "john_doe"
// If both are empty, output: "Anonymous"
//
// Example in practice:
//
//	<p>Welcome, {% firstof user.display_name user.email "Guest" %}!</p>
//
// Note: Output is automatically HTML-escaped when autoescape is enabled.
// Use the |safe filter if you need unescaped output.
type tagFirstofNode struct {
	position *Token
	args     []IEvaluator
}

// Execute evaluates arguments in order and outputs the first truthy value.
// HTML escaping is applied when autoescape is enabled (unless |safe is used).
func (node *tagFirstofNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// tagFirstofParser parses the {% firstof %} tag. It requires at least one
// expression argument; all arguments are parsed as potential fallback values.
func tagFirstofParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Django requires at least one argument

func init() {
	mustRegisterTag("firstof", tagFirstofParser)
}
