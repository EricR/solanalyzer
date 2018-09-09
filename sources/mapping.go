package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Mapping represents a mapping in Solidity.
type Mapping struct {
	Tokens
	Elementary *ElementaryTypeName
	TypeName   *TypeName
}

// NewMapping returns a new instance of Mapping.
func NewMapping() *Mapping {
	return &Mapping{}
}

// Visit is called by a visitor.
func (m *Mapping) Visit(ctx *parser.MappingContext) {
	m.Start = ctx.GetStart()
	m.Stop = ctx.GetStop()

	tn := NewTypeName()
	tn.Visit(ctx.TypeName().(*parser.TypeNameContext))

	etn := NewElementaryTypeName()
	etn.Visit(ctx.ElementaryTypeName().(*parser.ElementaryTypeNameContext))

	m.Elementary = etn
	m.TypeName = tn
}

func (m *Mapping) String() string {
	return fmt.Sprintf("mapping (%s=>%s)", m.Elementary, m.TypeName)
}
