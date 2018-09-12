package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalFunctionDefinition(function *sources.Function) {
	defer e.Recover(function.Tokens)

	for _, stmt := range function.Block.Statements {
		if stmt.SubType == sources.StatementSimple {
			e.evalSimpleStatement(stmt.SimpleStatement)
		}
	}
}
