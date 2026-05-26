package pongo2

// tagWithNode represents the {% with %} tag.
//
// The with tag creates a new scope with additional context variables.
// This is useful for caching complex expressions or simplifying nested
// attribute access.
//
// New style syntax (recommended):
//
//	{% with name=user.profile.full_name email=user.contact.email %}
//	    <p>{{ name }}</p>
//	    <p>{{ email }}</p>
//	{% endwith %}
//
// Old style syntax (Django-compatible):
//
//	{% with user.profile.full_name as name %}
//	    <p>{{ name }}</p>
//	{% endwith %}
//
// Multiple variables (new style):
//
//	{% with
//	    total=cart.items|length
//	    subtotal=cart.subtotal
//	    tax=cart.subtotal * 0.1
//	%}
//	    <p>Items: {{ total }}</p>
//	    <p>Subtotal: ${{ subtotal }}</p>
//	    <p>Tax: ${{ tax }}</p>
//	{% endwith %}
//
// Caching expensive computations:
//
//	{% with expensive_result=some_object.compute_heavy_operation %}
//	    <p>Result: {{ expensive_result }}</p>
//	    <p>Again: {{ expensive_result }}</p>
//	{% endwith %}
//
// Note: Variables defined in {% with %} are only available within the
// {% with %}...{% endwith %} block and do not affect the outer context.
type tagWithNode struct {
	withPairs map[string]IEvaluator
	wrapper   *NodeWrapper
}

// Execute creates a child context with the with-pairs and executes the
// wrapped content. The with-pairs are only visible within the block.
func (node *tagWithNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	// new context for block
	return nil
}

// Put all custom with-pairs into the context

// tagWithParser parses the {% with %} tag. It supports both new style (name=value)
// and old style (value as name) syntax for defining context variables.
func tagWithParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Scan through all arguments to see which style the user uses (old or new style).
// If we find any "as" keyword we will enforce old style; otherwise we will use new style.
// by default we're using the new_style

func init() {
	mustRegisterTag("with", tagWithParser)
}
