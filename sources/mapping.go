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
func (s *Source) NewMapping() *Mapping {
	mapping := &Mapping{}
	s.AddNode(mapping)

	return mapping
}

// Visit is called by a visitor.
func (m *Mapping) Visit(s *Source, ctx *parser.MappingContext) {
	m.Start = ctx.GetStart()
	m.Stop = ctx.GetStop()

	tn := s.NewTypeName()
	tn.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

	etn := s.NewElementaryTypeName()
	etn.Visit(s, ctx.ElementaryTypeName().(*parser.ElementaryTypeNameContext))

	m.Elementary = etn
	m.TypeName = tn
}

func (m *Mapping) String() string {
	return fmt.Sprintf("mapping (%s=>%s)", m.Elementary, m.TypeName)
}
