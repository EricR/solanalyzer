package sources

import "github.com/ericr/solanalyzer/parser"

// Expression represents a Solidity expression.
type Expression struct {
	Tokens
	Text        string
	Expressions []*Expression
}

// NewExpression returns a new instance of Expression.
func NewExpression() *Expression {
	return &Expression{}
}

// NewExpressionFromCtx returns a new instance of Expression from a given parser
// context.
func NewExpressionFromCtx(ctx *parser.ExpressionContext) *Expression {
	expr := NewExpression()

	for _, child := range ctx.GetChildren() {
		expr.Text += getText(child)
	}

	return expr
}

func (e *Expression) String() string {
	return e.Text
}
