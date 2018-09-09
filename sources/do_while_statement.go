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
func NewDoWhileStatement() *DoWhileStatement {
	return &DoWhileStatement{}
}

// Visit is called by a visitor.
func (dws *DoWhileStatement) Visit(ctx *parser.DoWhileStatementContext) {
	dws.Start = ctx.GetStart()
	dws.Stop = ctx.GetStop()

	bodyStmt := NewStatement()
	bodyStmt.Visit(ctx.Statement().(*parser.StatementContext))

	dws.Body = bodyStmt

	condition := NewExpression()
	condition.Visit(ctx.Expression().(*parser.ExpressionContext))

	dws.Condition = condition
}

func (dws *DoWhileStatement) String() string {
	return fmt.Sprintf("do %s while (%s)", dws.Body, dws.Condition)
}
