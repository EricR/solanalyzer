package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalContractDefinition(contract *sources.Contract) {
	defer e.Recover(contract.Tokens)

	e.Reset()

	for _, structDec := range contract.Structs {
		e.AddStructDeclaration(structDec)
	}

	for _, stateVarDec := range contract.StateVars {
		variable := e.evalStateVariableDeclaration(stateVarDec)
		e.SetVariable(variable)
	}

	for _, functionDef := range contract.Functions {
		e.AddFunctionDefinition(functionDef)
	}

	for _, functionDef := range contract.Functions {
		defer e.stack.Pop()

		e.stack.Push(contract, functionDef)
		e.evalFunctionDefinition(functionDef)
	}
}
