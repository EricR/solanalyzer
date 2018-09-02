package sources

import "github.com/ericr/solanalyzer/parser"

// EmitCall represents an emit call in Solidity.
type EmitCall struct {
	Expression *Expression
	Arguments  *FunctionCallArguments
}

// NewEmit returns a new instance of Emit.
func NewEmitCall() *EmitCall {
	return &EmitCall{}
}

// Visit is called by a visitor.
func (ec *EmitCall) Visit(ctx *parser.FunctionCallContext) {
	expr := NewExpression()
	expr.Visit(ctx.Expression().(*parser.ExpressionContext))

	args := NewFunctionCallArguments()
	args.Visit(ctx.FunctionCallArguments().(*parser.FunctionCallArgumentsContext))

	ec.Expression = expr
	ec.Arguments = args
}
