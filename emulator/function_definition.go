package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalFunctionDefinition(function *sources.Function) {
	defer e.Recover(function.Tokens)

	e.Emit("function_definition", &Event{
		Source:   e.source,
		Contract: e.stack.CurrentFrame().Contract,
		Function: function,
	})

	for _, param := range function.Parameters {
		e.setParamAsVariable(param)
	}

	for _, param := range function.Returns {
		e.setParamAsVariable(param)
	}

	for _, stmt := range function.Block.Statements {
		if stmt.SubType == sources.StatementSimple {
			e.evalSimpleStatement(stmt.SimpleStatement)
		}
	}
}

func (e *Emulator) setParamAsVariable(param *sources.Parameter) {
	// Parameters are stored in memory by default
	if param.StorageLocation == "" {
		param.StorageLocation = "memory"
	}

	variable := NewVariable(param.Identifier, param.TypeName, param.StorageLocation)
	e.SetVariable(variable)
}
