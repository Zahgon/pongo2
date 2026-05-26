package pongo2

import (
	"reflect"
)

const (
	varTypeInt = iota
	varTypeIdent
	varTypeSubscript
	varTypeArray
	varTypeNil
)

var (
	typeOfValuePtr   = reflect.TypeFor[*Value]()
	typeOfExecCtxPtr = reflect.TypeFor[*ExecutionContext]()
)

type variablePart struct {
	typ       int
	s         string
	i         int
	subscript IEvaluator
	isNil     bool

	isFunctionCall bool
	callingArgs    []functionCallArgument // needed for a function call, represents all argument nodes (INode supports nested function calls)
}

func (p *variablePart) String() string { _ = "STUB: not implemented"; return "" }

type functionCallArgument interface {
	Evaluate(*ExecutionContext) (*Value, error)
}

// TODO: Add location tokens
type stringResolver struct {
	locationToken *Token
	val           string
}

type intResolver struct {
	locationToken *Token
	val           int
}

type floatResolver struct {
	locationToken *Token
	val           float64
}

type boolResolver struct {
	locationToken *Token
	val           bool
}

type variableResolver struct {
	locationToken *Token

	parts []*variablePart
}

type nodeFilteredVariable struct {
	locationToken *Token

	resolver    IEvaluator
	filterChain []*filterCall
}

type nodeVariable struct {
	locationToken *Token
	expr          IEvaluator
}

type executionCtxEval struct{}

// executeEvaluator is a helper that evaluates and writes a value to the writer.
func executeEvaluator(e IEvaluator, ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *nodeFilteredVariable) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (vr *variableResolver) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stringResolver) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *intResolver) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *floatResolver) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *boolResolver) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *nodeFilteredVariable) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (vr *variableResolver) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (s *stringResolver) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (i *intResolver) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (f *floatResolver) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (b *boolResolver) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (s *stringResolver) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *intResolver) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *floatResolver) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *boolResolver) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stringResolver) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (i *intResolver) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (f *floatResolver) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (b *boolResolver) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (nv *nodeVariable) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (nv *nodeVariable) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// apply escape filter

func (executionCtxEval) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vr *variableResolver) FilterApplied(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (vr *variableResolver) String() string { _ = "STUB: not implemented"; return "" }

func (vr *variableResolver) resolve(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	// Handle in-template array definition
	return nil, nil
}

// Unpack *Value if needed

// Resolve interface to concrete value

// Handle function call

// resolveArrayDefinition handles in-template array definitions like [a, b, c].
func (vr *variableResolver) resolveArrayDefinition(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookupInitialValue looks up the first part of the variable in the context.
func (vr *variableResolver) lookupInitialValue(ctx *ExecutionContext) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// unpackValue unpacks a *Value if the current value is of that type.
func (vr *variableResolver) unpackValue(current reflect.Value, isSafe bool) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// resolveNextPart resolves the next part of a variable path from the current value.
// Returns (resolved value, isNil, error).
func (vr *variableResolver) resolveNextPart(
	ctx *ExecutionContext,
	current reflect.Value,
	part *variablePart,
) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	// Check for method call first
	return *new(reflect.Value), false, nil
}

// Resolve pointer

// resolvePartByType resolves a variable part based on its type.
func (vr *variableResolver) resolvePartByType(
	ctx *ExecutionContext,
	current reflect.Value,
	part *variablePart,
) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// resolveIntIndex resolves an integer index access on a slice/array/string.
func (vr *variableResolver) resolveIntIndex(current reflect.Value, part *variablePart) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// For strings, return the character (rune) at the index (Django-compatible behavior)

// resolveIdentifier resolves a field or map key access by name.
func (vr *variableResolver) resolveIdentifier(current reflect.Value, part *variablePart) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// resolveSubscript resolves a subscript access (e.g., foo[bar]).
func (vr *variableResolver) resolveSubscript(
	ctx *ExecutionContext,
	current reflect.Value,
	part *variablePart,
) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}

// For strings, return the character (rune) at the index (Django-compatible behavior)

// callResult holds the result of a function call resolution.
type callResult struct {
	value  reflect.Value
	isSafe bool
}

// handleFunctionCall processes a function call on the current value and returns the result.
func (vr *variableResolver) handleFunctionCall(
	ctx *ExecutionContext,
	current reflect.Value,
	part *variablePart,
) (*callResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If an implicit ExecCtx is needed

// Validate input argument count

// Validate output argument count

// Evaluate and prepare parameters

// Execute the function call

// prepareCallParameters evaluates arguments and prepares them for function call.
func (vr *variableResolver) prepareCallParameters(
	ctx *ExecutionContext,
	t reflect.Type,
	currArgs []functionCallArgument,
) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate all parameters

// getFnArgType returns the expected type for a function argument at the given index.
func (vr *variableResolver) getFnArgType(t reflect.Type, idx, numArgs int, isVariadic bool) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// convertArgToParam converts an evaluated Value to a reflect.Value suitable for function call.
func (vr *variableResolver) convertArgToParam(pv *Value, fnArg reflect.Type, idx int, isVariadic bool) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Check type compatibility

// executeCall performs the actual function call and processes the result.
func (vr *variableResolver) executeCall(
	current reflect.Value,
	t reflect.Type,
	parameters []reflect.Value,
) (*callResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for error return value

func (vr *variableResolver) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *nodeFilteredVariable) FilterApplied(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *nodeFilteredVariable) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// "[" [expr {, expr}] "]"
func (p *Parser) parseArray() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// We consume '['

// We allow an empty list, so check for a closing bracket.

// parsing an array declaration with at least one expression

// No closing bracket, so we're parsing an expression

// If there's a closing bracket after an expression, we will stop parsing the arguments

// If there's NO closing bracket, there MUST be an comma

func (p *Parser) parseNumberLiteral(sign int, numToken *Token, locToken *Token) (IEvaluator, error) {
	_ = "STUB: not implemented"
	// One exception to the rule that we don't have float64 literals is at the beginning
	// of an expression (or a variable name). Since we know we started with an integer
	// which can't obviously be a variable name, we can check whether the first number
	// is followed by dot (and then a number again). If so we're converting it to a float64.
	return *new(IEvaluator), nil
}

// IDENT | IDENT.(IDENT|NUMBER)... | IDENT[expr]... | "[" [ expr {, expr}] "]"
//
//nolint:gocyclo,cyclop,funlen // parser for variable expressions handles many token types
func (p *Parser) parseVariableOrLiteral() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Is first part a number or a string, there's nothing to resolve (because there's only to return the value then)

// Parsing an array literal [expr {, expr}]

// Negative number literal
// consume '-'

// consume the number

// First part of a variable MUST be an identifier

// we consumed the first identifier of the variable name

// Next variable part (can be either NUMBER or IDENT)

// consume: IDENT

// consume: NUMBER

// consume: NIL

// EOF

// Variable subscript

// Function call
// FunctionName '(' Comma-separated list of expressions ')'

// No closing bracket, so we're parsing an expression

// If there's a closing bracket after an expression, we will stop parsing the arguments

// If there's NO closing bracket, there MUST be an comma

// We got a closing bracket, so stop parsing arguments

// We're done parsing the function call, next variable part

// No dot, subscript or function call? Then we're done with the variable parsing

func (p *Parser) parseVariableOrLiteralWithFilter() (*nodeFilteredVariable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse the variable name

// Parse all the filters

// Parse one single filter

// Check sandbox filter restriction

func (p *Parser) parseVariableElement() (INode, error) {
	_ = "STUB: not implemented"
	return *new(INode), nil
}

// consume '{{'
