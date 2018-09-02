package sources

import "github.com/ericr/solanalyzer/parser"

// Emit represents an emit in Solidity.
type Emit struct {
	EmitCall *EmitCall
}

// NewEmit returns a new instance of Emit.
func NewEmit() *Emit {
	return &Emit{}
}

// Visit is called by a visitor.
func (e *Emit) Visit(ctx *parser.EmitStatementContext) {
	call := NewEmitCall()
	call.Visit(ctx.FunctionCall().(*parser.FunctionCallContext))

	e.EmitCall = call
}
