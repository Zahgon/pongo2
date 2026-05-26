package pongo2

// tagExtendsNode represents the {% extends %} tag.
//
// The extends tag indicates that this template extends a parent template.
// It must be the first tag in the template if used. The child template can
// override blocks defined in the parent template using {% block %} tags.
//
// Usage:
//
//	{% extends "base.html" %}
//	{% block content %}
//	    My custom content that overrides the parent's content block.
//	{% endblock %}
//
// Example parent template (base.html):
//
//	<!DOCTYPE html>
//	<html>
//	<head><title>{% block title %}Default{% endblock %}</title></head>
//	<body>
//	    <header>{% block header %}Default Header{% endblock %}</header>
//	    <main>{% block content %}{% endblock %}</main>
//	    <footer>{% block footer %}Default Footer{% endblock %}</footer>
//	</body>
//	</html>
//
// Example child template (page.html):
//
//	{% extends "base.html" %}
//	{% block title %}My Page{% endblock %}
//	{% block content %}
//	    <h1>Welcome to my page!</h1>
//	{% endblock %}
//
// Note: Only one extends tag is allowed per template, and it must be at the root level.
type tagExtendsNode struct {
	filename string
}

// Execute is a no-op for extends nodes. The inheritance relationship is
// established at parse time; execution happens through the parent template.
func (node *tagExtendsNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"

	// tagExtendsParser parses the {% extends %} tag. It requires a string filename
	// argument and establishes the parent-child template relationship at parse time.
	return nil
}

func tagExtendsParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Already one parent

// prepared, static template

// Get parent's filename

// Parse the parent

// Keep track of things

func init() {
	mustRegisterTag("extends", tagExtendsParser)
}
