package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// IfStatement represents an if statement in Solidity.
type IfStatement struct {
	Tokens
	If   *Expression
	Body *Statement
	Else *Statement
}

// NewIfStatement returns a new instance of IfStatement.
func (s *Source) NewIfStatement() *IfStatement {
	stmt := &IfStatement{}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (is *IfStatement) Visit(s *Source, ctx *parser.IfStatementContext) {
	is.Start = ctx.GetStart()
	is.Stop = ctx.GetStop()

	ifExpr := s.NewExpression()
	ifExpr.Visit(s, ctx.Expression().(*parser.ExpressionContext))

	is.If = ifExpr

	bodyStmt := s.NewStatement()
	bodyStmt.Visit(s, ctx.Statement(0).(*parser.StatementContext))

	is.Body = bodyStmt

	if ctx.Statement(1) != nil {
		elseStmt := s.NewStatement()
		elseStmt.Visit(s, ctx.Statement(1).(*parser.StatementContext))

		is.Else = elseStmt
	}
}

func (is *IfStatement) String() string {
	str := fmt.Sprintf("if (%s) %s", is.If, is.Body)

	if is.Else != nil {
		str += fmt.Sprintf(" else %s", is.Else)
	}

	return str
}
