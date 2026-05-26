package pongo2

type Expression struct {
	// TODO: Add location token?
	expr1   IEvaluator
	expr2   IEvaluator
	opToken *Token
}

type relationalExpression struct {
	// TODO: Add location token?
	expr1   IEvaluator
	expr2   IEvaluator
	opToken *Token
}

type notExpression struct {
	expr IEvaluator
}

type simpleExpression struct {
	negativeSign bool
	term1        IEvaluator
	term2        IEvaluator
	opToken      *Token
}

type term struct {
	// TODO: Add location token?
	factor1 IEvaluator
	factor2 IEvaluator
	opToken *Token
}

type power struct {
	// TODO: Add location token?
	power1 IEvaluator
	power2 IEvaluator
}

func (expr *Expression) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (expr *relationalExpression) FilterApplied(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (expr *notExpression) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (expr *simpleExpression) FilterApplied(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (expr *term) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (expr *power) FilterApplied(name string) bool { _ = "STUB: not implemented"; return false }

func (expr *Expression) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *relationalExpression) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *notExpression) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *simpleExpression) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *term) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *power) GetPositionToken() *Token { _ = "STUB: not implemented"; return nil }

func (expr *Expression) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *relationalExpression) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *notExpression) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *simpleExpression) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *term) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *power) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (expr *Expression) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (expr *relationalExpression) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (expr *notExpression) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (expr *simpleExpression) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Result will be a string

// Result will be a float

// Result will be an integer

// Result will be a float

// Result will be an integer

func (expr *term) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Result will be float

// Result will be int

// Result will be float

// Result will be int

// Result will be int

func (expr *power) Evaluate(ctx *ExecutionContext) (*Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) parseFactor() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

func (p *Parser) parsePower() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Shortcut for faster evaluation

func (p *Parser) parseTerm() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Create new sub-term

// Shortcut for faster evaluation

func (p *Parser) parseSimpleExpression() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// New sub expr

// Shortcut for faster evaluation

func (p *Parser) parseRelationalExpression() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Shortcut for faster evaluation

func (p *Parser) parseNotExpression() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Support chained not: "not not x"

// parseAndExpression handles "and" / "&&" (higher precedence than or).
func (p *Parser) parseAndExpression() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Shortcut for faster evaluation

// ParseExpression handles "or" / "||" (lowest precedence among logical operators).
// Precedence: or < and < not < relational < additive < multiplicative < power < factor
func (p *Parser) ParseExpression() (IEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(IEvaluator), nil
}

// Shortcut for faster evaluation
