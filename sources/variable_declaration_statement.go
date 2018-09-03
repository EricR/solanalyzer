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
func (s *Source) NewVariableDeclarationStatement() *VariableDeclarationStatement {
	stmt := &VariableDeclarationStatement{
		Identifiers: []string{},
	}
	s.AddNode(stmt)

	return stmt
}

// Visit is called by a visitor.
func (vds *VariableDeclarationStatement) Visit(s *Source, ctx *parser.VariableDeclarationStatementContext) {
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
		varDec := s.NewVariableDeclaration()
		varDec.Visit(s, ctx.VariableDeclaration().(*parser.VariableDeclarationContext))

		vds.VariableDeclaration = varDec

	case ctx.VariableDeclarationList() != nil:
		varDecList := ctx.VariableDeclarationList().(*parser.VariableDeclarationListContext)
		for _, varDecCtx := range varDecList.AllVariableDeclaration() {
			varDec := s.NewVariableDeclaration()
			varDec.Visit(s, varDecCtx.(*parser.VariableDeclarationContext))

			vds.VariableDeclarationList = append(vds.VariableDeclarationList, varDec)
		}
	default:
		panic("Unknown type of variable declaration")
	}

	if ctx.Expression() != nil {
		expr := s.NewExpression()
		expr.Visit(s, ctx.Expression().(*parser.ExpressionContext))

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
		str += fmt.Sprintf(" %s", vds.VariableDeclaration)
	case len(vds.VariableDeclarationList) > 0:
		for _, varDec := range vds.VariableDeclarationList {
			str += fmt.Sprintf(" %s", varDec)
		}
	}

	if vds.Expression != nil {
		str += fmt.Sprintf(" %s", vds.Expression)
	}

	return str
}
