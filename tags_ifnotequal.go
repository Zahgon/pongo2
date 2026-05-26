package pongo2

// tagIfNotEqualNode represents the {% ifnotequal %} tag.
//
// DEPRECATED: This tag is considered legacy in Django. Use {% if %} with
// comparison operators instead: {% if var1 != var2 %}
//
// The ifnotequal tag compares two values and renders the block if they are NOT equal.
//
// Basic usage:
//
//	{% ifnotequal user.name "Admin" %}
//	    Welcome, regular user!
//	{% endifnotequal %}
//
// Preferred alternative using {% if %}:
//
//	{% if user.name != "Admin" %}
//	    Welcome, regular user!
//	{% endif %}
//
// Deprecated: Use {% if var1 != var2 %} instead.
type tagIfNotEqualNode struct {
	var1, var2  IEvaluator
	thenWrapper *NodeWrapper
	elseWrapper *NodeWrapper
}

// Execute compares two values and renders the then block if NOT equal,
// otherwise renders the else block (if present).
func (node *tagIfNotEqualNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// tagIfNotEqualParser parses the {% ifnotequal %} tag. It requires exactly
// two expression arguments to compare for inequality.
func tagIfNotEqualParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Parse two expressions

// Wrap then/else-blocks

// if there's an else in the if-statement, we need the else-Block as well

func init() {
	mustRegisterTag("ifnotequal", tagIfNotEqualParser)
}
