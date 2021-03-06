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
func NewWhileStatement() *WhileStatement {
	return &WhileStatement{}
}

// Visit is called by a visitor.
func (ws *WhileStatement) Visit(ctx *parser.WhileStatementContext) {
	ws.Start = ctx.GetStart()
	ws.Stop = ctx.GetStop()

	cond := NewExpression()
	cond.Visit(ctx.Expression().(*parser.ExpressionContext))

	ws.Condition = cond

	bodyStmt := NewStatement()
	bodyStmt.Visit(ctx.Statement().(*parser.StatementContext))

	ws.Body = bodyStmt
}

func (ws *WhileStatement) String() string {
	return fmt.Sprintf("while (%s) %s", ws.Condition, ws.Body)
}
