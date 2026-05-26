package pongo2

// tagIfEqualNode represents the {% ifequal %} tag.
//
// DEPRECATED: This tag is considered legacy in Django. Use {% if %} with
// comparison operators instead: {% if var1 == var2 %}
//
// The ifequal tag compares two values and renders the block if they are equal.
//
// Basic usage:
//
//	{% ifequal user.name "John" %}
//	    Hello, John!
//	{% endifequal %}
//
// Preferred alternative using {% if %}:
//
//	{% if user.name == "John" %}
//	    Hello, John!
//	{% endif %}
//
// Deprecated: Use {% if var1 == var2 %} instead.
type tagIfEqualNode struct {
	var1, var2  IEvaluator
	thenWrapper *NodeWrapper
	elseWrapper *NodeWrapper
}

// Execute compares two values and renders the then block if equal,
// otherwise renders the else block (if present).
func (node *tagIfEqualNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// tagIfEqualParser parses the {% ifequal %} tag. It requires exactly
// two expression arguments to compare for equality.
func tagIfEqualParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Parse two expressions

// Wrap then/else-blocks

// if there's an else in the if-statement, we need the else-Block as well

func init() {
	mustRegisterTag("ifequal", tagIfEqualParser)
}
