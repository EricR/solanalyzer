package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// Parameter represents a parameter in Solidity.
type Parameter struct {
	Tokens
	Identifier      string
	TypeName        *TypeName
	StorageLocation string
}

// NewParameter returns a new instance of Parameter.
func (s *Source) NewParameter() *Parameter {
	param := &Parameter{}
	s.AddNode(param)

	return param
}

// Visit is called by a visitor.
func (p *Parameter) Visit(s *Source, ctx *parser.ParameterContext) {
	p.Start = ctx.GetStart()
	p.Stop = ctx.GetStop()

	if ctx.Identifier() != nil {
		p.Identifier = ctx.Identifier().GetText()
	}

	if ctx.StorageLocation() != nil {
		p.StorageLocation = ctx.StorageLocation().GetText()
	}

	typeName := s.NewTypeName()
	typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

	p.TypeName = typeName
}

func (p *Parameter) String() string {
	if p.Identifier != "" {
		return fmt.Sprintf("%s %s", p.TypeName, p.Identifier)
	}
	return p.TypeName.String()
}

func paramsToString(params []*Parameter) string {
	strs := []string{}

	for _, param := range params {
		strs = append(strs, param.String())
	}

	return strings.Join(strs, ", ")
}
