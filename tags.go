package pongo2

/* Incomplete:
   -----------

   verbatim (only the "name" argument is missing for verbatim)

   Reconsideration:
   ----------------

   debug (reason: not sure what to output yet)
   regroup / Grouping on other properties (reason: maybe too python-specific; not sure how useful this would be in Go)

   Following built-in tags wont be added:
   --------------------------------------

   csrf_token (reason: web-framework specific)
   load (reason: python-specific)
   url (reason: web-framework specific)
*/

// INodeTag is a semantic interface for template tags returned by TagParser functions.
// While structurally identical to INode, it provides API clarity and future extensibility.
//
//nolint:iface // intentional semantic type for public API clarity and future extension
type INodeTag interface {
	INode
}

// This is the function signature of the tag's parser you will have
// to implement in order to create a new tag.
//
// 'doc' is providing access to the whole document while 'arguments'
// is providing access to the user's arguments to the tag:
//
//	{% your_tag_name some "arguments" 123 %}
//
// start_token will be the *Token with the tag's name in it (here: your_tag_name).
//
// Please see the Parser documentation on how to use the parser.
// See RegisterTag()'s documentation for more information about
// writing a tag as well.
type TagParser func(doc *Parser, start *Token, arguments *Parser) (INodeTag, error)

type tag struct {
	name   string
	parser TagParser
}

var builtinTags = make(map[string]*tag)

// copyTags creates a shallow copy of a tag map.
func copyTags(src map[string]*tag) map[string]*tag { _ = "STUB: not implemented"; return nil }

func mustRegisterTag(name string, parserFn TagParser) { _ = "STUB: not implemented"; return }

// registerTagGlobal registers a new tag to the global tag map.
// This is used during package initialization to register builtin tags.
func registerTagGlobal(name string, parserFn TagParser) error {
	_ = "STUB: not implemented"
	return nil
}

// Tag = "{%" IDENT ARGS "%}"
func (p *Parser) parseTagElement() (INodeTag, error) {
	_ = "STUB: not implemented"
	// consume "{%"
	return *new(INodeTag), nil
}

// Check for identifier

// Check sandbox tag restriction

// Check for the existing tag

// Does not exists

// Add token to args

// next token

// EOF?

// This is done to have nice EOF error messages
