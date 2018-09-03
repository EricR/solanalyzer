package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// VariableDeclaration represents a variable declaration in Solidity.
type VariableDeclaration struct {
	Tokens
	TypeName        *TypeName
	StorageLocation string
	Identifier      string
}

// NewVariableDeclaration returns a new instance of VariableDeclaration.
func (s *Source) NewVariableDeclaration() *VariableDeclaration {
	varDec := &VariableDeclaration{}
	s.AddNode(varDec)

	return varDec
}

// Visit is called by a visitor.
func (vd *VariableDeclaration) Visit(s *Source, ctx *parser.VariableDeclarationContext) {
	vd.Start = ctx.GetStart()
	vd.Stop = ctx.GetStop()

	typeName := s.NewTypeName()
	typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

	vd.TypeName = typeName

	if ctx.StorageLocation() != nil {
		vd.StorageLocation = ctx.StorageLocation().GetText()
	}

	vd.Identifier = ctx.Identifier().GetText()
}

func (vd *VariableDeclaration) String() string {
	if vd.StorageLocation == "" {
		return fmt.Sprintf("%s %s", vd.TypeName.String(), vd.Identifier)
	}
	return fmt.Sprintf("%s %s %s", vd.TypeName.String(), vd.StorageLocation, vd.Identifier)
}
