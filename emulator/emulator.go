package emulator

import (
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
)

// Emulator represents a Solidity emulator.
type Emulator struct {
	source    *sources.Source
	stateVars map[string]*Variable
	localVars map[string]*Variable
}

// New returns a new instance of Emulator.
func New(source *sources.Source) *Emulator {
	return &Emulator{
		source:    source,
		stateVars: map[string]*Variable{},
		localVars: map[string]*Variable{},
	}
}

// Run runs the emulator.
func (e *Emulator) Run() {
	for _, contract := range e.source.Contracts {
		e.evalContractDefinition(contract)
	}
}

// Eval evaluates an expression, returning zero or more values.
func (e *Emulator) Eval(expr *sources.Expression) []*Value {
	defer e.Recover(expr.Tokens)

	switch expr.SubType {
	case sources.ExpressionPrimary:
		return e.evalPrimary(expr.Primary)

	case sources.ExpressionParentheses:
		return e.Eval(expr.SubExpression)

	case sources.ExpressionBinaryOperation:
		return e.evalBinaryOperation(expr.Operation, expr.LeftExpression, expr.RightExpression)
	}

	return []*Value{}
}

// Recover acts as a recovery strategy when the emulator panics.
func (e *Emulator) Recover(tokens sources.Tokens) {
	if r := recover(); r != nil {
		logrus.Errorf("Error evaluating %s:%d:%d: %s",
			e.source.FilePath, tokens.Start.GetLine(), tokens.Start.GetColumn(), r)
	}
}

// MustFindVariable must return a variable, otherwise the program panics.
func (e *Emulator) MustFindVariable(name string) *Variable {
	variable, ok := e.localVars[name]
	if !ok {
		panic("Undeclared variable")
	}

	return variable
}

// MustAssignVariable must find and assign a variable, otherwise the program
// panics.
func (e *Emulator) MustAssignVariable(name string, val *Value) {
	e.MustFindVariable(name).Value = val
}
