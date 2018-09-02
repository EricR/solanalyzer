package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// VariableDeclarationStatement represents a variable declaration statement in
// Solidity.
type VariableDeclarationStatement struct {
	Tokens
	Identifiers             []string
	VariableDeclaration     *VariableDeclaration
	VariableDeclarationList []*VariableDeclaration
	Expression              *Expression
}

// NewVariableDeclarationStatement returns a new instance of SimpleStatement.
func NewVariableDeclarationStatement() *VariableDeclarationStatement {
	return &VariableDeclarationStatement{
		Identifiers: []string{},
	}
}

// Visit is called by a visitor.
func (vds *VariableDeclarationStatement) Visit(ctx *parser.VariableDeclarationStatementContext) {
	vds.Start = ctx.GetStart()
	vds.Stop = ctx.GetStop()

	switch {
	case ctx.IdentifierList() != nil:
		identifierList := ctx.IdentifierList().(*parser.IdentifierListContext)
		for _, identifierCtx := range identifierList.AllIdentifier() {
			identifier := identifierCtx.(*parser.IdentifierContext).GetText()
			vds.Identifiers = append(vds.Identifiers, identifier)
		}

	case ctx.VariableDeclaration() != nil:
		varDec := NewVariableDeclaration()
		varDec.Visit(ctx.VariableDeclaration().(*parser.VariableDeclarationContext))

		vds.VariableDeclaration = varDec

	case ctx.VariableDeclarationList() != nil:
		varDecList := ctx.VariableDeclarationList().(*parser.VariableDeclarationListContext)
		for _, varDecCtx := range varDecList.AllVariableDeclaration() {
			varDec := NewVariableDeclaration()
			varDec.Visit(varDecCtx.(*parser.VariableDeclarationContext))

			vds.VariableDeclarationList = append(vds.VariableDeclarationList, varDec)
		}
	}

	if ctx.Expression() != nil {
		expr := NewExpression()
		expr.Visit(ctx.Expression().(*parser.ExpressionContext))

		vds.Expression = expr
	}
}

func (vds *VariableDeclarationStatement) String() string {
	str := "var"

	switch {
	case len(vds.Identifiers) > 0:
		for _, identifier := range vds.Identifiers {
			str += fmt.Sprintf(" %s", identifier)
		}
	case vds.VariableDeclaration != nil:
		str += fmt.Sprintf(" %s", vds.VariableDeclaration.String())
	case len(vds.VariableDeclarationList) > 0:
		for _, varDec := range vds.VariableDeclarationList {
			str += fmt.Sprintf(" %s", varDec.String())
		}
	}

	if vds.Expression != nil {
		str += fmt.Sprintf(" %s", vds.Expression.String())
	}

	return fmt.Sprintf("%s;", str)
}
