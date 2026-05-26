package pongo2

// tagSSINode represents the {% ssi %} tag.
//
// DEPRECATED: This tag was removed from Django in version 1.10.
// Use {% include %} instead, which provides better functionality and security.
//
// See: https://code.djangoproject.com/ticket/24022
//
// The ssi (Server Side Include) tag includes the contents of another file
// into the template. It can include files as plain text or as parsed templates.
//
// Including a file as plain text (content is not parsed):
//
//	{% ssi "static/robots.txt" %}
//
// Output: Contents of robots.txt displayed as-is
//
// Including a file as a parsed template:
//
//	{% ssi "includes/header.html" parsed %}
//
// With "parsed", the file is treated as a template and has access to
// the current context variables.
//
// Deprecated: Use {% include %} instead.
type tagSSINode struct {
	filename string
	content  string
	template *Template
}

// Execute outputs the file content. If "parsed" was specified, the content
// is executed as a template with the current context; otherwise it's output as-is.
func (node *tagSSINode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil

	// Execute the template within the current context
}

// Just print out the content

// tagSSIParser parses the {% ssi %} tag. It requires a filename string and
// optionally accepts "parsed" to treat the file as a template.
func tagSSIParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// parsed

// plaintext - use the template loader to support virtual filesystems

func init() {
	mustRegisterTag("ssi", tagSSIParser)
}
