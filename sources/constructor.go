package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Constructor represents a constructor in Solidity.
type Constructor struct {
	Tokens
	Parameters []*Parameter
	Modifiers  *ModifierList
	Block      *Block
}

// NewConstructor returns a new instance of Constructor.
func (s *Source) NewConstructor() *Constructor {
	constructor := &Constructor{
		Parameters: []*Parameter{},
		Modifiers:  s.NewModifierList(),
	}
	s.AddNode(constructor)

	return constructor
}

// Visit is called by a visitor.
func (c *Constructor) Visit(s *Source, ctx *parser.ConstructorDefinitionContext) {
	c.Start = ctx.GetStart()
	c.Stop = ctx.GetStop()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		param := s.NewParameter()
		param.Visit(s, paramCtx.(*parser.ParameterContext))

		c.Parameters = append(c.Parameters, param)
	}

	modifiers := s.NewModifierList()
	modifiers.Visit(s, ctx.ModifierList().(*parser.ModifierListContext))

	c.Modifiers = modifiers

	if ctx.Block() != nil {
		block := s.NewBlock()
		block.Visit(s, ctx.Block().(*parser.BlockContext))

		c.Block = block
	}
}

func (c *Constructor) String() string {
	return fmt.Sprintf("constructor(%s)", paramsToString(c.Parameters))
}
