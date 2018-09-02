package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Function represents a Solidity function.
type Function struct {
	Tokens
	Identifier string
	Parameters *ParameterList
	Modifiers  *ModifierList
	Returns    *ParameterList
	Block      *Block
}

// NewFunction returns a new instance of Function.
func NewFunction() *Function {
	return &Function{
		Parameters: NewParameterList(),
		Modifiers:  NewModifierList(),
		Returns:    NewParameterList(),
	}
}

// Visit is called by a visitor.
func (f *Function) Visit(ctx *parser.FunctionDefinitionContext) {
	f.Start = ctx.GetStart()
	f.Stop = ctx.GetStop()
	f.Identifier = ctx.Identifier().GetText()

	paramList := ctx.ParameterList().(*parser.ParameterListContext)

	for _, paramCtx := range paramList.AllParameter() {
		parameter := NewParameter()
		parameter.Visit(paramCtx.(*parser.ParameterContext))

		f.Parameters.Add(parameter)
	}

	modifiers := NewModifierList()
	modifiers.Visit(ctx.ModifierList().(*parser.ModifierListContext))

	f.Modifiers = modifiers

	if ctx.ReturnParameters() != nil {
		returnParams := ctx.ReturnParameters().(*parser.ReturnParametersContext)
		returnList := returnParams.ParameterList().(*parser.ParameterListContext)

		for _, paramCtx := range returnList.AllParameter() {
			parameter := NewParameter()
			parameter.Visit(paramCtx.(*parser.ParameterContext))

			f.Returns.Add(parameter)
		}
	}

	if ctx.Block() != nil {
		block := NewBlock()
		block.Visit(ctx.Block().(*parser.BlockContext))

		f.Block = block
	}
}

func (f *Function) ShortSignature() string {
	return fmt.Sprintf("%s(%s)", f.Identifier, f.Parameters)
}

func (f *Function) String() string {
	str := fmt.Sprintf("%s(%s)", f.Identifier, f.Parameters)

	if len(f.Modifiers.String()) > 0 {
		str = fmt.Sprintf("%s %s", str, f.Modifiers)
	}

	if len(*f.Returns) > 0 {
		str = fmt.Sprintf("%s returns(%s)", str, f.Returns)
	}

	return str
}
