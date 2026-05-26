package pongo2

// FilterFunction is the type filter functions must fulfil
type FilterFunction func(in *Value, param *Value) (out *Value, err error)

var builtinFilters = make(map[string]FilterFunction)

// copyFilters creates a shallow copy of a filter map.
func copyFilters(src map[string]FilterFunction) map[string]FilterFunction {
	_ = "STUB: not implemented"
	return nil
}

// BuiltinFilterExists returns true if the given filter is a built-in filter.
// Use TemplateSet.FilterExists to check filters in a specific template set.
func BuiltinFilterExists(name string) bool { _ = "STUB: not implemented"; return false }

// BuiltinTagExists returns true if the given tag is registered in builtinTags.
// Use TemplateSet.TagExists to check tags in a specific template set.
func BuiltinTagExists(name string) bool { _ = "STUB: not implemented"; return false }

// registerFilterBuiltin registers a new filter to the global filter map.
// This is used during package initialization to register builtin filters.
func registerFilterBuiltin(name string, fn FilterFunction) error {
	_ = "STUB: not implemented"
	return nil
}

// MustApplyFilter behaves like ApplyFilter, but panics on an error.
// This function uses builtinFilters. Use TemplateSet.MustApplyFilter for set-specific filters.
func MustApplyFilter(name string, value *Value, param *Value) *Value {
	_ = "STUB: not implemented"
	return nil
}

// ApplyFilter applies a built-infilter to a given value using the given
// parameters. Returns a *pongo2.Value or an error. Use TemplateSet.ApplyFilter
// for set-specific filters.
func ApplyFilter(name string, value *Value, param *Value) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make sure param is a *Value

type filterCall struct {
	token *Token

	name      string
	parameter IEvaluator

	filterFunc FilterFunction
}

func (fc *filterCall) Execute(v *Value, ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter = IDENT | IDENT ":" FilterArg | IDENT "|" Filter
func (p *Parser) parseFilter() (*filterCall, error) { _ = "STUB: not implemented"; return nil, nil }

// Check filter ident

// Check sandbox filter restriction

// Get the appropriate filter function and bind it

// Check for filter-argument (2 tokens needed: ':' ARG)

// Get filter argument expression
