package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// DoWhileStatement represents a do while statement in Solidity.
type DoWhileStatement struct {
	Tokens
	Body      *Statement
	Condition *Expression
}

// NewDoWhileStatement returns a new instance of DoWhileStatement.
func (s *Source) NewDoWhileStatement() *DoWhileStatement {
	stmt := &DoWhileStatement{}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (dws *DoWhileStatement) Visit(s *Source, ctx *parser.DoWhileStatementContext) {
	dws.Start = ctx.GetStart()
	dws.Stop = ctx.GetStop()

	bodyStmt := s.NewStatement()
	bodyStmt.Visit(s, ctx.Statement().(*parser.StatementContext))

	dws.Body = bodyStmt

	condition := s.NewExpression()
	condition.Visit(s, ctx.Expression().(*parser.ExpressionContext))

	dws.Condition = condition
}

func (dws *DoWhileStatement) String() string {
	return fmt.Sprintf("do %s while (%s)", dws.Body, dws.Condition)
}
