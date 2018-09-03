package sources

import (
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

const (
	// FunctionCallArgsWithNameValues is a function call with name values.
	FunctionCallArgsWithNameValues = iota
	// FunctionCallArgsWithExprs is a function call with expressions.
	FunctionCallArgsWithExprs
)

// FunctionCallArguments represents a list of function call arguments in
// Solidity.
type FunctionCallArguments struct {
	Tokens
	SubType     int
	NameValues  []*NameValue
	Expressions []*Expression
}

// NewFunctionCallArguments returns a new instance of FunctionCallArguments.
func NewFunctionCallArguments() *FunctionCallArguments {
	return &FunctionCallArguments{
		Expressions: []*Expression{},
	}
}

// Visit is called by a visitor.
func (fca *FunctionCallArguments) Visit(ctx *parser.FunctionCallArgumentsContext) {
	fca.Start = ctx.GetStart()
	fca.Stop = ctx.GetStop()

	switch {
	case ctx.NameValueList() != nil:
		fca.SubType = FunctionCallArgsWithNameValues

		nvList := ctx.NameValueList().(*parser.NameValueListContext)

		for _, nvCtx := range nvList.AllNameValue() {
			nv := NewNameValue()
			nv.Visit(nvCtx.(*parser.NameValueContext))

			fca.NameValues = append(fca.NameValues, nv)
		}

	case ctx.ExpressionList() != nil:
		fca.SubType = FunctionCallArgsWithExprs

		exprList := ctx.ExpressionList().(*parser.ExpressionListContext)

		for _, exprCtx := range exprList.AllExpression() {
			expr := NewExpression()
			expr.Visit(exprCtx.(*parser.ExpressionContext))

			fca.Expressions = append(fca.Expressions, expr)
		}
	default:
		panic("Unknown type of function call arguments")
	}
}

func (fca *FunctionCallArguments) String() string {
	args := []string{}

	switch fca.SubType {
	case FunctionCallArgsWithNameValues:
		for _, nv := range fca.NameValues {
			args = append(args, nv.String())
		}

	case FunctionCallArgsWithExprs:
		for _, expr := range fca.Expressions {
			args = append(args, expr.String())
		}
	}

	return strings.Join(args, ", ")
}
