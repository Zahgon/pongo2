package pongo2

// tagIncludeNode represents the {% include %} tag.
//
// The include tag renders another template and inserts its output at the
// current location. The included template has access to the current context.
//
// Basic usage:
//
//	{% include "header.html" %}
//	<main>Content here</main>
//	{% include "footer.html" %}
//
// Using "if_exists" to silently skip missing templates:
//
//	{% include "optional_sidebar.html" if_exists %}
//
// Passing additional context with "with":
//
//	{% include "user_card.html" with username=user.name avatar=user.avatar %}
//
// Using "only" to exclude parent context (included template only sees with variables):
//
//	{% include "widget.html" with title="My Widget" only %}
//
// Dynamic template names (lazy evaluation):
//
//	{% include template_name %}
//	{% include "partials/"|add:partial_name|add:".html" %}
//
// The "only" keyword must come after all with pairs:
//
//	{% include "card.html" with title="Hello" subtitle="World" only %}
//
// Note: Static filenames (strings) are parsed at compile time for better
// performance. Dynamic filenames are resolved at runtime.
type tagIncludeNode struct {
	tpl               *Template
	filenameEvaluator IEvaluator
	lazy              bool
	only              bool
	filename          string
	withPairs         map[string]IEvaluator
	ifExists          bool
}

// Execute renders the included template with the appropriate context.
// For lazy includes, the filename is evaluated at runtime; otherwise
// the pre-parsed template is executed directly.
func (node *tagIncludeNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	// Building the context for the template
	return nil
}

// Fill the context with all data from the parent

// Put all custom with-pairs into the context

// Execute the template

// Evaluate the filename

// Get include-filename

// if this is ReadFile error, and "if_exists" flag is enabled

// Template is already parsed with static filename

// tagIncludeEmptyNode is a placeholder node returned when a static include
// with "if_exists" references a non-existent file at parse time.
type tagIncludeEmptyNode struct{}

// Execute is a no-op for empty include nodes (missing template with if_exists).
func (node *tagIncludeEmptyNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"

	// tagIncludeParser parses the {% include %} tag. It supports static or dynamic
	// filenames, "if_exists" flag, "with" context pairs, and "only" isolation.
	return nil
}

func tagIncludeParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// prepared, static template

// "if_exists" flag

// Get include-filename

// Check if this template is currently being parsed (recursive include)
// If so, we must use lazy evaluation to avoid infinite recursion at parse time

// Recursive include detected - use lazy evaluation

// Create a simple evaluator that returns the static filename

// Parse the included template

// if this is ReadFile error, and "if_exists" token presents we should create and empty node

// No String, then the user wants to use lazy-evaluation (slower, but possible)

// "if_exists" flag

// After having parsed the filename we're gonna parse the with+only options

// We have at least one key=expr pair (because of starting "with")

// Only?

// stop parsing arguments because it's the last option

// Support "only" without "with" (Django allows {% include "file" only %})

func init() {
	mustRegisterTag("include", tagIncludeParser)
}
