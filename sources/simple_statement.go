package sources

import (
	"github.com/ericr/solanalyzer/parser"
)

const (
	// SimpleStatementVarDec is a simple statement with a variable declaration.
	SimpleStatementVarDec = iota
	// SimpleStatementExpr is a simple statement with an expression.
	SimpleStatementExpr
)

// SimpleStatement represents a simple statement in Solidity.
type SimpleStatement struct {
	Tokens
	SubType             int
	VariableDeclaration *VariableDeclarationStatement
	Expression          *Expression
}

// NewSimpleStatement returns a new instance of SimpleStatement.
func (s *Source) NewSimpleStatement() *SimpleStatement {
	stmt := &SimpleStatement{}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (ss *SimpleStatement) Visit(s *Source, ctx *parser.SimpleStatementContext) {
	ss.Start = ctx.GetStart()
	ss.Stop = ctx.GetStop()

	if ctx.VariableDeclarationStatement() != nil {
		varDecStmtCtx := ctx.VariableDeclarationStatement()
		varDecStmt := s.NewVariableDeclarationStatement()
		varDecStmt.Visit(s, varDecStmtCtx.(*parser.VariableDeclarationStatementContext))

		ss.SubType = SimpleStatementVarDec
		ss.VariableDeclaration = varDecStmt
		return
	}

	exprStmtCtx := ctx.ExpressionStatement().(*parser.ExpressionStatementContext)

	expr := s.NewExpression()
	expr.Visit(s, exprStmtCtx.Expression().(*parser.ExpressionContext))

	ss.SubType = SimpleStatementExpr
	ss.Expression = expr
}

func (ss *SimpleStatement) String() string {
	if ss.SubType == SimpleStatementVarDec {
		return ss.VariableDeclaration.String()
	}
	return ss.Expression.String()
}
