package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Function represents a Solidity function.
type Function struct {
	Tokens
	Contract   *Contract
	Identifier string
	Parameters []*Parameter
	Modifiers  *ModifierList
	Returns    []*Parameter
	Block      *Block
}

// NewFunction returns a new instance of Function.
func (s *Source) NewFunction() *Function {
	fn := &Function{
		Parameters: []*Parameter{},
		Modifiers:  s.NewModifierList(),
		Returns:    []*Parameter{},
	}
	s.AddNode(fn)

	return fn
}

// Visit is called by a visitor.
func (f *Function) Visit(s *Source, ctx *parser.FunctionDefinitionContext) {
	f.Start = ctx.GetStart()
	f.Stop = ctx.GetStop()
	f.Identifier = ctx.Identifier().GetText()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		param := s.NewParameter()
		param.Visit(s, paramCtx.(*parser.ParameterContext))

		f.Parameters = append(f.Parameters, param)
	}

	modifiers := s.NewModifierList()
	modifiers.Visit(s, ctx.ModifierList().(*parser.ModifierListContext))

	f.Modifiers = modifiers

	if ctx.ReturnParameters() != nil {
		returnParams := ctx.ReturnParameters().(*parser.ReturnParametersContext)
		returnList := returnParams.ParameterList().(*parser.ParameterListContext)

		for _, paramCtx := range returnList.AllParameter() {
			param := s.NewParameter()
			param.Visit(s, paramCtx.(*parser.ParameterContext))

			f.Returns = append(f.Returns, param)
		}
	}

	if ctx.Block() != nil {
		block := s.NewBlock()
		block.Visit(s, ctx.Block().(*parser.BlockContext))

		f.Block = block
	}
}

// ShortSignature returns an abbreviated version of String().
func (f *Function) ShortSignature() string {
	return fmt.Sprintf("%s(%s)", f.Identifier, paramsToString(f.Parameters))
}

func (f *Function) String() string {
	str := fmt.Sprintf("%s(%s)", f.Identifier, paramsToString(f.Parameters))

	if len(f.Modifiers.String()) > 0 {
		str = fmt.Sprintf("%s %s", str, f.Modifiers)
	}

	if len(f.Returns) > 0 {
		str = fmt.Sprintf("%s returns(%s)", str, paramsToString(f.Returns))
	}

	return str
}
