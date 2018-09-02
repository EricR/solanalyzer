package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Modifier represents a modifier in Solidity.
type Modifier struct {
	Tokens
	Identifier string
	Parameters *ParameterList
	Block      *Block
}

// NewModifier returns a new instance of Modifier.
func NewModifier() *Modifier {
	return &Modifier{
		Parameters: NewParameterList(),
	}
}

// Visit is called by a visitor.
func (m *Modifier) Visit(ctx *parser.FunctionDefinitionContext) {
	m.Start = ctx.GetStart()
	m.Stop = ctx.GetStop()
	m.Identifier = ctx.Identifier().GetText()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		parameter := NewParameter()
		parameter.Visit(paramCtx.(*parser.ParameterContext))

		m.Parameters.Add(parameter)
	}

	if ctx.Block() != nil {
		block := NewBlock()
		block.Visit(ctx.Block().(*parser.BlockContext))

		m.Block = block
	}
}

func (m *Modifier) String() string {
	return fmt.Sprintf("modifier %s(%s)", m.Identifier, m.Parameters)
}
