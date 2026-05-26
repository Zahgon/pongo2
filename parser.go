package pongo2

// INode is the base interface for all executable template nodes.
// See INodeTag for template tags returned by tag parsers.
//
//nolint:iface // base interface used for polymorphism across node types
type INode interface {
	Execute(*ExecutionContext, TemplateWriter) error
}

type IEvaluator interface {
	INode
	GetPositionToken() *Token
	Evaluate(*ExecutionContext) (*Value, error)
	FilterApplied(name string) bool
}

// The parser provides you a comprehensive and easy tool to
// work with the template document and arguments provided by
// the user for your custom tag.
//
// The parser works on a token list which will be provided by pongo2.
// A token is a unit you can work with. Tokens are either of type identifier,
// string, number, keyword, HTML or symbol.
//
// (See Token's documentation for more about tokens)
type Parser struct {
	name      string
	idx       int
	tokens    []*Token
	lastToken *Token

	// if the parser parses a template document, here will be
	// a reference to it (needed to access the template through Tags)
	template *Template
}

// Creates a new parser to parse tokens.
// Used inside pongo2 to parse documents and to provide an easy-to-use
// parser for tag authors
func newParser(name string, tokens []*Token, template *Template) *Parser {
	_ = "STUB: not implemented"
	return nil
}

// Consume one token. It will be gone forever.
func (p *Parser) Consume() {
	_ = "STUB: not implemented"

	// Consume N tokens. They will be gone forever.
	return
}

func (p *Parser) ConsumeN(count int) {
	_ = "STUB: not implemented"

	// Returns the current token.
	return
}

func (p *Parser) Current() *Token { _ = "STUB: not implemented"; return nil }

// Returns the CURRENT token if the given type matches.
// Consumes this token on success.
func (p *Parser) MatchType(typ TokenType) *Token { _ = "STUB: not implemented"; return nil }

// Returns the CURRENT token if the given type AND value matches.
// Consumes this token on success.
func (p *Parser) Match(typ TokenType, val string) *Token { _ = "STUB: not implemented"; return nil }

// Returns the CURRENT token if the given type AND *one* of
// the given values matches.
// Consumes this token on success.
func (p *Parser) MatchOne(typ TokenType, vals ...string) *Token {
	_ = "STUB: not implemented"
	return nil
}

// Returns the CURRENT token if the given type matches.
// It DOES NOT consume the token.
func (p *Parser) PeekType(typ TokenType) *Token { _ = "STUB: not implemented"; return nil }

// Returns the CURRENT token if the given type AND value matches.
// It DOES NOT consume the token.
func (p *Parser) Peek(typ TokenType, val string) *Token { _ = "STUB: not implemented"; return nil }

// Returns the CURRENT token if the given type AND *one* of
// the given values matches.
// It DOES NOT consume the token.
func (p *Parser) PeekOne(typ TokenType, vals ...string) *Token {
	_ = "STUB: not implemented"
	return nil
}

// Returns the tokens[current position + shift] token if the
// given type AND value matches for that token.
// DOES NOT consume the token.
func (p *Parser) PeekN(shift int, typ TokenType, val string) *Token {
	_ = "STUB: not implemented"
	return nil
}

// Returns the tokens[current position + shift] token if the given type matches.
// DOES NOT consume the token for that token.
func (p *Parser) PeekTypeN(shift int, typ TokenType) *Token { _ = "STUB: not implemented"; return nil }

// Returns the UNCONSUMED token count.
func (p *Parser) Remaining() int { _ = "STUB: not implemented"; return 0 }

// Returns the total token count.
func (p *Parser) Count() int { _ = "STUB: not implemented"; return 0 }

// Returns tokens[i] or NIL (if i >= len(tokens))
func (p *Parser) Get(i int) *Token { _ = "STUB: not implemented"; return nil }

// Returns tokens[current-position + shift] or NIL
// (if (current-position + i) >= len(tokens))
func (p *Parser) GetR(shift int) *Token { _ = "STUB: not implemented"; return nil }

// Error produces a nice error message and returns an error-object.
// The 'token'-argument is optional. If provided, it will take
// the token's position information. If not provided, it will
// automatically use the CURRENT token's position information.
func (p *Parser) Error(msg string, token *Token) *Error {
	_ = "STUB: not implemented"

	// Set current token
	return nil
}

// Set to last token

// Wraps all nodes between starting tag and "{% endtag %}" and provides
// one simple interface to execute the wrapped nodes.
// It returns a parser to process provided arguments to the tag.
func (p *Parser) WrapUntilTag(names ...string) (*NodeWrapper, *Parser, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// New tag, check whether we have to stop wrapping here

// We've found a (!) end-tag

// We only process the tag if we've found an end tag

// Okay, endtag found.
// '{%' tagname

// Okay, end the wrapping here

// Otherwise process next element to be wrapped

// Skips all nodes between starting tag and "{% endtag %}"
func (p *Parser) SkipUntilTag(names ...string) error { _ = "STUB: not implemented"; return nil }

// New tag, check whether we have to stop wrapping here

// We've found an (!) end-tag

// We only process the tag if we've found an end tag

// Okay, endtag found.
// '{%' tagname

// Done skipping, exit.

// If we haven't encountered '%}', we consume whatever
// there might be.

// EOF encountered
