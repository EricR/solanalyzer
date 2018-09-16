package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalStateVariableDeclaration(stateVar *sources.StateVariable) *Variable {
	defer e.Recover(stateVar.Tokens)

	variable := NewVariable(stateVar.Identifier, stateVar.TypeName, "storage")

	if stateVar.Expression != nil {
		values := e.Eval(stateVar.Expression)

		if len(values) != 1 {
			panic("State variable expression must return one value")
		}
		variable.Value = values[0]
	}

	return variable
}
