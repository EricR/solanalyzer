package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Modifier represents a modifier in Solidity.
type Modifier struct {
	Tokens
	Identifier string
	Parameters []*Parameter
	Block      *Block
}

// NewModifier returns a new instance of Modifier.
func (s *Source) NewModifier() *Modifier {
	modifier := &Modifier{
		Parameters: []*Parameter{},
	}
	s.AddNode(modifier)

	return modifier
}

// Visit is called by a visitor.
func (m *Modifier) Visit(s *Source, ctx *parser.FunctionDefinitionContext) {
	m.Start = ctx.GetStart()
	m.Stop = ctx.GetStop()
	m.Identifier = ctx.Identifier().GetText()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		param := s.NewParameter()
		param.Visit(s, paramCtx.(*parser.ParameterContext))

		m.Parameters = append(m.Parameters, param)
	}

	if ctx.Block() != nil {
		block := s.NewBlock()
		block.Visit(s, ctx.Block().(*parser.BlockContext))

		m.Block = block
	}
}

func (m *Modifier) String() string {
	return fmt.Sprintf("modifier %s(%s)", m.Identifier, paramsToString(m.Parameters))
}
