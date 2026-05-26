package pongo2

import (
	"strings"
)

// maxLoremCount limits the maximum number of lorem items to prevent abuse.
const maxLoremCount = 100000

// tagLoremParagraphs and tagLoremWords are pre-split lorem ipsum text.
var (
	tagLoremParagraphs = strings.Split(tagLoremText, "\n")
	tagLoremWords      = strings.Fields(tagLoremText)
)

// tagLoremNode represents the {% lorem %} tag.
//
// Verified against Django 4.2 with script (paragraph separator behavior).
// Django difference: pongo2 uses static pre-defined paragraphs while Django
// generates random text from a word list. Paragraph count and separator
// behavior ("\n\n") matches Django.
//
// The lorem tag generates placeholder "Lorem Ipsum" text, useful for
// prototyping and design mockups.
//
// Basic usage (1 paragraph, plain text):
//
//	{% lorem %}
//
// Output: "Lorem ipsum dolor sit amet, consectetur adipisici elit..."
//
// Specifying number of items:
//
//	{% lorem 3 %}      {# 3 paragraphs #}
//	{% lorem 5 w %}    {# 5 words #}
//	{% lorem 2 p %}    {# 2 HTML paragraphs #}
//
// Method options:
//   - b: Plain text paragraphs (default)
//   - w: Individual words
//   - p: HTML paragraphs wrapped in <p> tags
//
// Using random option (randomizes selection instead of sequential):
//
//	{% lorem 3 p random %}
//
// Examples with output:
//
//	{% lorem 5 w %}
//	{# Output: "Lorem ipsum dolor sit amet" #}
//
//	{% lorem 1 p %}
//	{# Output: "<p>Lorem ipsum dolor sit amet...</p>" #}
//
// Note: The maximum count is 100,000 to prevent abuse.
type tagLoremNode struct {
	position *Token
	count    int    // number of paragraphs
	method   string // w = words, p = HTML paragraphs, b = plain-text (default is b)
	random   bool   // does not use the default paragraph "Lorem ipsum dolor sit amet, ..."
}

// writeLoremItems writes items from the source slice with separator, prefix, and suffix.
func writeLoremItems(writer TemplateWriter, count int, source []string, sep, prefix, suffix string, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // G404: lorem ipsum generation, cryptographic randomness not needed

// Execute outputs lorem ipsum text according to the configured method
// (words, paragraphs, or HTML paragraphs) and count.
func (node *tagLoremNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// Django: "\n\n".join(paras)

// Django: "\n\n".join("<p>%s</p>" % p for p in paras)

// tagLoremParser parses the {% lorem %} tag. It accepts an optional count,
// method (w/p/b), and "random" flag.
func tagLoremParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

func init() {
	mustRegisterTag("lorem", tagLoremParser)
}

//nolint:dupword // standard lorem ipsum text naturally contains repeated Latin words
const tagLoremText = `Lorem ipsum dolor sit amet, consectetur adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.
Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi.
Nam liber tempor cum soluta nobis eleifend option congue nihil imperdiet doming id quod mazim placerat facer possim assum. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat.
Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis.
At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, At accusam aliquyam diam diam dolore dolores duo eirmod eos erat, et nonumy sed tempor et et invidunt justo labore Stet clita ea et gubergren, kasd magna no rebum. sanctus sea sed takimata ut vero voluptua. est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat.
Consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.`
