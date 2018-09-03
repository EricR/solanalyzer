package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// ImportDeclaration represents an import declaration in Solidity.
type ImportDeclaration struct {
	Tokens
	Identifier string
	As         string
}

// NewImportDeclaration returns a new instance of ImportDeclaration.
func (s *Source) NewImportDeclaration() *ImportDeclaration {
	dec := &ImportDeclaration{}
	s.AddNode(dec)

	return dec
}

// Visit is called by a visitor.
func (id *ImportDeclaration) Visit(s *Source, ctx *parser.ImportDeclarationContext) {
	id.Start = ctx.GetStart()
	id.Stop = ctx.GetStop()

	id.Identifier = ctx.Identifier(0).GetText()

	if ctx.Identifier(1) != nil {
		id.As = ctx.Identifier(1).GetText()
	}
}

func (id *ImportDeclaration) String() string {
	if id.As == "" {
		return id.Identifier
	}
	return fmt.Sprintf("%s as %s", id.Identifier, id.As)
}
