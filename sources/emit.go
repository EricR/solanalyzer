package sources

import "github.com/ericr/solanalyzer/parser"

// Emit represents an emit in Solidity.
type Emit struct {
	EmitCall *EmitCall
}

// NewEmit returns a new instance of Emit.
func (s *Source) NewEmit() *Emit {
	emit := &Emit{}
	s.AddNode(emit)

	return emit
}

// Visit is called by a visitor.
func (e *Emit) Visit(s *Source, ctx *parser.EmitStatementContext) {
	call := s.NewEmitCall()
	call.Visit(s, ctx.FunctionCall().(*parser.FunctionCallContext))

	e.EmitCall = call
}
