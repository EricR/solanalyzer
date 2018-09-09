package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// ForStatement represents a for statement in Solidity.
type ForStatement struct {
	Tokens
	Initialization *SimpleStatement
	Condition      *Expression
	Increment      *Expression
	Body           *Statement
}

// NewForStatement returns a new instance of ForStatement.
func NewForStatement() *ForStatement {
	return &ForStatement{}
}

// Visit is called by a visitor.
func (fs *ForStatement) Visit(ctx *parser.ForStatementContext) {
	fs.Start = ctx.GetStart()
	fs.Stop = ctx.GetStop()

	if ctx.SimpleStatement() != nil {
		simpleStatement := NewSimpleStatement()
		simpleStatement.Visit(ctx.SimpleStatement().(*parser.SimpleStatementContext))

		fs.Initialization = simpleStatement
	}

	if ctx.Expression(0) != nil {
		expr := NewExpression()
		expr.Visit(ctx.Expression(0).(*parser.ExpressionContext))

		fs.Condition = expr
	}

	if ctx.Expression(1) != nil {
		expr := NewExpression()
		expr.Visit(ctx.Expression(1).(*parser.ExpressionContext))

		fs.Increment = expr
	}

	body := NewStatement()
	body.Visit(ctx.Statement().(*parser.StatementContext))

	fs.Body = body
}

func (fs *ForStatement) String() string {
	return fmt.Sprintf("for (%s; %s; %s) %s",
		fs.Initialization, fs.Condition, fs.Increment, fs.Body)
}
