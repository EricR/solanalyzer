package sources

import "github.com/ericr/solanalyzer/parser"

// Expression represents a Solidity expression.
type Expression struct {
	Tokens
	Expressions []*Expression
}

// NewExpression returns a new instance of Expression.
func NewExpression() *Expression {
	return &Expression{}
}

// Visit is called by a visitor.
func (e *Expression) Visit(ctx *parser.ExpressionContext) {
}

func (e *Expression) String() string {
	return "TODO"
}
