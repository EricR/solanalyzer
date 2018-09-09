package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// FunctionTypeName represents a function type name in Solidity.
type FunctionTypeName struct {
	Tokens
	Parameters      []*FunctionTypeParameter
	Internal        bool
	External        bool
	StateMutability *StateMutability
	Returns         []*FunctionTypeParameter
}

// NewFunctionTypeName returns a new instance of FunctionTypeName.
func NewFunctionTypeName() *FunctionTypeName {
	return &FunctionTypeName{}
}

// Visit is called by a visitor.
func (ftn *FunctionTypeName) Visit(ctx *parser.FunctionTypeNameContext) {
	ftn.Start = ctx.GetStart()
	ftn.Stop = ctx.GetStop()

	paramList := ctx.FunctionTypeParameterList(0).(*parser.FunctionTypeParameterListContext)

	for _, paramCtx := range paramList.AllFunctionTypeParameter() {
		param := NewFunctionTypeParameter()
		param.Visit(paramCtx.(*parser.FunctionTypeParameterContext))

		ftn.Parameters = append(ftn.Parameters, param)
	}

	ftn.Internal = ctx.InternalKeyword(0) != nil
	ftn.External = ctx.ExternalKeyword(0) != nil
	ftn.StateMutability = NewStateMutabilityFromCtxs(ctx.AllStateMutability())

	returnList := ctx.FunctionTypeParameterList(1).(*parser.FunctionTypeParameterListContext)

	for _, paramCtx := range returnList.AllFunctionTypeParameter() {
		param := NewFunctionTypeParameter()
		param.Visit(paramCtx.(*parser.FunctionTypeParameterContext))

		ftn.Returns = append(ftn.Returns, param)
	}
}

func (ftn *FunctionTypeName) String() string {
	paramStrs := []string{}
	returnStrs := []string{}

	for _, param := range ftn.Parameters {
		paramStrs = append(paramStrs, param.String())
	}
	for _, param := range ftn.Returns {
		returnStrs = append(returnStrs, param.String())
	}

	str := fmt.Sprintf("function(%s) %s", strings.Join(paramStrs, ", "),
		ftn.StateMutability)

	switch {
	case ftn.Internal:
		str += " internal"
	case ftn.External:
		str += " external"
	}

	return fmt.Sprintf("%s returns(%s)", str, strings.Join(returnStrs, ", "))
}
