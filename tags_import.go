package pongo2

// tagImportNode represents the {% import %} tag.
//
// The import tag imports macros from another template file, making them
// available as callable functions in the current template.
//
// Basic usage (importing a single macro):
//
//	{% import "macros.html" input_field %}
//	{{ input_field("username", "Enter your name") }}
//
// Importing multiple macros:
//
//	{% import "forms/macros.html" input_field, textarea, select_box %}
//	{{ input_field("email", "Email address") }}
//	{{ textarea("bio", "Tell us about yourself", 5) }}
//
// Using aliases with "as":
//
//	{% import "macros.html" input_field as field, textarea as ta %}
//	{{ field("name", "Your name") }}
//	{{ ta("description", "Description", 3) }}
//
// The imported macros must be defined with "export" in the source template:
//
//	{# In macros.html #}
//	{% macro input_field(name, label) export %}
//	    <label>{{ label }}</label>
//	    <input type="text" name="{{ name }}">
//	{% endmacro %}
//
// Note: Only macros marked with "export" can be imported.
type tagImportNode struct {
	position *Token
	filename string
	macros   map[string]*tagMacroNode // alias/name -> macro instance
}

// Execute registers imported macros as callable functions in the private context.
// Each macro becomes available under its name (or alias) as a function.
func (node *tagImportNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// tagImportParser parses the {% import %} tag. It requires a filename string
// followed by one or more macro names to import, with optional "as" aliases.
func tagImportParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

// Compile the given template

func init() {
	mustRegisterTag("import", tagImportParser)
}
