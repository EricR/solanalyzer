package sources

import "github.com/ericr/solanalyzer/parser"

// Expression represents a Solidity expression.
type Expression struct {
	Tokens
	Expressions []*Expression
}

// NewExpression returns a new instance of Expression.
func NewExpression() *Expression {
	return &Expression{
		Expressions: []*Expression{},
	}
}

// Visit is called by a visitor.
func (e *Expression) Visit(ctx *parser.ExpressionContext) {
	e.Start = ctx.GetStart()
	e.Stop = ctx.GetStop()
}

func (e *Expression) String() string {
	return "TODO"
}
