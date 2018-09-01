package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// StateVar represents a state variable in Solidity.
type StateVar struct {
	Tokens
}

// NewStateVar returns a new instance of StateVar.
func NewStateVar() *StateVar {
	return &StateVar{}
}

// Visit is called by a visitor. See source.go for additional information on
// this pattern.
func (sv *StateVar) Visit(ctx *parser.StateVariableDeclarationContext) {
	sv.Start = ctx.GetStart()
	sv.Stop = ctx.GetStop()
}

func (sv *StateVar) String() string {
	return fmt.Sprintf("TODO")
}
