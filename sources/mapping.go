package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Mapping represents a mapping in Solidity.
type Mapping struct {
	Elementary *ElementaryTypeName
	TypeName   *TypeName
}

// NewMapping returns a new instance of Mapping.
func NewMapping() *Mapping {
	return &Mapping{}
}

// Visit is called by a visitor. See source.go for additional information on
// this pattern.
func (m *Mapping) Visit(ctx *parser.MappingContext) {
	tn := NewTypeName()
	tn.Visit(ctx.TypeName().(*parser.TypeNameContext))

	m.Elementary = NewElementaryTypeName(ctx.ElementaryTypeName().GetText())
	m.TypeName = tn
}

func (m *Mapping) String() string {
	return fmt.Sprintf("mapping (%s=>%s)",
		m.Elementary, m.TypeName.String())
}
