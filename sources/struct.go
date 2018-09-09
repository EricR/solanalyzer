package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Struct represents a struct in Solidity.
type Struct struct {
	Tokens
	Identifier           string
	VariableDeclarations []*VariableDeclaration
}

// NewStruct returns a new instance of Struct.
func NewStruct() *Struct {
	return &Struct{}
}

// Visit is called by a visitor.
func (s *Struct) Visit(ctx *parser.StructDefinitionContext) {
	s.Start = ctx.GetStart()
	s.Stop = ctx.GetStop()

	for _, varDecCtx := range ctx.AllVariableDeclaration() {
		varDec := NewVariableDeclaration()
		varDec.Visit(varDecCtx.(*parser.VariableDeclarationContext))

		s.VariableDeclarations = append(s.VariableDeclarations, varDec)
	}
}

func (s *Struct) String() string {
	str := fmt.Sprintf("struct %s {", s.Identifier)

	for _, varDeclaration := range s.VariableDeclarations {
		str += fmt.Sprintf("%s;", varDeclaration.String())
	}

	str += "}"

	return str
}
