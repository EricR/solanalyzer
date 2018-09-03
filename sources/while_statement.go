package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// WhileStatement represents a while statement in Solidity.
type WhileStatement struct {
	Tokens
	Condition *Expression
	Body      *Statement
}

// NewWhileStatement returns a new instance of WhileStatement.
func (s *Source) NewWhileStatement() *WhileStatement {
	stmt := &WhileStatement{}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (ws *WhileStatement) Visit(s *Source, ctx *parser.WhileStatementContext) {
	ws.Start = ctx.GetStart()
	ws.Stop = ctx.GetStop()

	cond := s.NewExpression()
	cond.Visit(s, ctx.Expression().(*parser.ExpressionContext))

	ws.Condition = cond

	bodyStmt := s.NewStatement()
	bodyStmt.Visit(s, ctx.Statement().(*parser.StatementContext))

	ws.Body = bodyStmt
}

func (ws *WhileStatement) String() string {
	return fmt.Sprintf("while (%s) %s", ws.Condition, ws.Body)
}
