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
func NewUsingFor() *UsingFor {
	return &UsingFor{}
}

// Visit is called by a visitor.
func (uf *UsingFor) Visit(ctx *parser.UsingForDeclarationContext) {
	uf.Start = ctx.GetStart()
	uf.Stop = ctx.GetStop()
	uf.Identifier = ctx.Identifier().GetText()

	if ctx.TypeName() != nil {
		typeName := NewTypeName()
		typeName.Visit(ctx.TypeName().(*parser.TypeNameContext))

		uf.TypeName = typeName
	}
}

func (uf *UsingFor) String() string {
	if uf.TypeName == nil {
		return fmt.Sprintf("using %s for *", uf.Identifier)
	}
	return fmt.Sprintf("using %s for %s", uf.Identifier, uf.TypeName)
}
