package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// Enum represents an enum in Solidity.
type Enum struct {
	Tokens
	Contract   *Contract
	Identifier string
	Values     []string
}

// NewEnum returns a new instance of Enum.
func (s *Source) NewEnum() *Enum {
	enum := &Enum{}
	s.AddNode(enum)

	return enum
}

// Visit is called by a visitor.
func (e *Enum) Visit(s *Source, ctx *parser.EnumDefinitionContext) {
	e.Start = ctx.GetStart()
	e.Stop = ctx.GetStop()
	e.Identifier = ctx.Identifier().GetText()

	for _, valueCtx := range ctx.AllEnumValue() {
		value := valueCtx.(*parser.EnumValueContext)
		e.Values = append(e.Values, value.Identifier().GetText())
	}
}

func (e *Enum) String() string {
	return fmt.Sprintf("enum %s {%s}", e.Identifier, strings.Join(e.Values, ","))
}
