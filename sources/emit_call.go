package sources

import "github.com/ericr/solanalyzer/parser"

// EmitCall represents an emit call in Solidity.
type EmitCall struct {
	Expression *Expression
	Arguments  *FunctionCallArguments
}

// NewEmitCall returns a new instance of Emit.
func (s *Source) NewEmitCall() *EmitCall {
	emitCall := &EmitCall{}
	s.AddNode(emitCall)

	return emitCall
}

// Visit is called by a visitor.
func (ec *EmitCall) Visit(s *Source, ctx *parser.FunctionCallContext) {
	expr := s.NewExpression()
	expr.Visit(s, ctx.Expression().(*parser.ExpressionContext))

	args := s.NewFunctionCallArguments()
	args.Visit(s, ctx.FunctionCallArguments().(*parser.FunctionCallArgumentsContext))

	ec.Expression = expr
	ec.Arguments = args
}
