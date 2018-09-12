package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalContractDefinition(contract *sources.Contract) {
	defer e.Recover(contract.Tokens)

	e.stateVars = map[string]*Variable{}

	for _, stateVar := range contract.StateVars {
		variable := NewVariable(stateVar.Identifier, stateVar.TypeName)

		if stateVar.Expression != nil {
			vals := e.Eval(stateVar.Expression)

			if len(vals) != 1 {
				panic("State variable expressions must return one value")
			}
			variable.Value = e.Eval(stateVar.Expression)[0]
		}

		e.stateVars[stateVar.Identifier] = variable
	}

	for _, function := range contract.Functions {
		e.localVars = map[string]*Variable{}
		e.evalFunctionDefinition(function)
	}
}
