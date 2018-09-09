package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// ReturnStatement represents a return statement in Solidity.
type ReturnStatement struct {
	Tokens
	Expression *Expression
}

// NewReturnStatement returns a new instance of ReturnStatement.
func NewReturnStatement() *ReturnStatement {
	return &ReturnStatement{}
}

// Visit is called by a visitor.
func (rs *ReturnStatement) Visit(ctx *parser.ReturnStatementContext) {
	rs.Start = ctx.GetStart()
	rs.Stop = ctx.GetStop()

	expr := NewExpression()
	expr.Visit(ctx.Expression().(*parser.ExpressionContext))

	rs.Expression = expr
}

func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("return %s;", rs.Expression)
}
