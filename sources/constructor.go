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
func NewConstructor() *Constructor {
	return &Constructor{
		Parameters: []*Parameter{},
		Modifiers:  NewModifierList(),
	}
}

// Visit is called by a visitor.
func (c *Constructor) Visit(ctx *parser.ConstructorDefinitionContext) {
	c.Start = ctx.GetStart()
	c.Stop = ctx.GetStop()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		param := NewParameter()
		param.Visit(paramCtx.(*parser.ParameterContext))

		c.Parameters = append(c.Parameters, param)
	}

	modifiers := NewModifierList()
	modifiers.Visit(ctx.ModifierList().(*parser.ModifierListContext))

	c.Modifiers = modifiers

	if ctx.Block() != nil {
		block := NewBlock()
		block.Visit(ctx.Block().(*parser.BlockContext))

		c.Block = block
	}
}

func (c *Constructor) String() string {
	return fmt.Sprintf("constructor(%s)", paramsToString(c.Parameters))
}
