package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// UsingFor represents a using for declaration in Solidity.
type UsingFor struct {
	Tokens
	Contract   *Contract
	Identifier string
	TypeName   *TypeName
}

// NewUsingFor returns a new instance of UsingFor.
func (s *Source) NewUsingFor() *UsingFor {
	uf := &UsingFor{}
	s.AddNode(uf)

	return uf
}

// Visit is called by a visitor.
func (uf *UsingFor) Visit(s *Source, ctx *parser.UsingForDeclarationContext) {
	uf.Start = ctx.GetStart()
	uf.Stop = ctx.GetStop()
	uf.Identifier = ctx.Identifier().GetText()

	if ctx.TypeName() != nil {
		typeName := s.NewTypeName()
		typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

		uf.TypeName = typeName
	}
}

func (uf *UsingFor) String() string {
	if uf.TypeName == nil {
		return fmt.Sprintf("using %s for *", uf.Identifier)
	}
	return fmt.Sprintf("using %s for %s", uf.Identifier, uf.TypeName)
}
