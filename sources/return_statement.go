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
func (s *Source) NewReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (rs *ReturnStatement) Visit(s *Source, ctx *parser.ReturnStatementContext) {
	rs.Start = ctx.GetStart()
	rs.Stop = ctx.GetStop()

	expr := s.NewExpression()
	expr.Visit(s, ctx.Expression().(*parser.ExpressionContext))

	rs.Expression = expr
}

func (rs *ReturnStatement) String() string {
	return fmt.Sprintf("return %s", rs.Expression)
}
