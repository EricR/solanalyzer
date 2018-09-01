package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Parameter represents a parameter in Solidity.
type Parameter struct {
	Identifier      string
	TypeName        *TypeName
	StorageLocation string
}

// NewParameter returns a new instance of Parameter.
func NewParameter() *Parameter {
	return &Parameter{}
}

// Visit is called by a visitor. See source.go for additional information on
// this pattern.
func (p *Parameter) Visit(ctx *parser.ParameterContext) {
	if ctx.Identifier() != nil {
		p.Identifier = ctx.Identifier().GetText()
	}

	if ctx.StorageLocation() != nil {
		p.StorageLocation = ctx.StorageLocation().GetText()
	}

	typeName := NewTypeName()
	typeName.Visit(ctx.TypeName().(*parser.TypeNameContext))

	p.TypeName = typeName
}

func (p *Parameter) String() string {
	if p.Identifier != "" {
		return fmt.Sprintf("%s %s", p.TypeName, p.Identifier)
	}
	return p.TypeName.String()
}
