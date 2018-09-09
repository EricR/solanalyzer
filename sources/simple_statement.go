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
func NewSimpleStatement() *SimpleStatement {
	return &SimpleStatement{}
}

// Visit is called by a visitor.
func (ss *SimpleStatement) Visit(ctx *parser.SimpleStatementContext) {
	ss.Start = ctx.GetStart()
	ss.Stop = ctx.GetStop()

	if ctx.VariableDeclarationStatement() != nil {
		varDecStmtCtx := ctx.VariableDeclarationStatement()
		varDecStmt := NewVariableDeclarationStatement()
		varDecStmt.Visit(varDecStmtCtx.(*parser.VariableDeclarationStatementContext))

		ss.SubType = SimpleStatementVarDec
		ss.VariableDeclaration = varDecStmt
		return
	}

	exprStmtCtx := ctx.ExpressionStatement().(*parser.ExpressionStatementContext)

	expr := NewExpression()
	expr.Visit(exprStmtCtx.Expression().(*parser.ExpressionContext))

	ss.SubType = SimpleStatementExpr
	ss.Expression = expr
}

func (ss *SimpleStatement) String() string {
	if ss.SubType == SimpleStatementVarDec {
		return ss.VariableDeclaration.String()
	}
	return ss.Expression.String()
}
