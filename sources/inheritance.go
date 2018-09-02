package sources

import "github.com/ericr/solanalyzer/parser"

// Inheritance represents inheritance in Solidity.
type Inheritance struct {
	Tokens
	TypeName    *UserDefinedTypeName
	Expressions []*Expression
}

// NewInheritance returns a new instance of Inheritance.
func NewInheritance() *Inheritance {
	return &Inheritance{}
}

// Visit is called by a visitor.
func (i *Inheritance) Visit(ctx *parser.InheritanceSpecifierContext) {
	i.Start = ctx.GetStart()
	i.Stop = ctx.GetStop()
}
