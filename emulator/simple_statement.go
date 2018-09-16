package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalSimpleStatement(stmt *sources.SimpleStatement) {
	defer e.Recover(stmt.Tokens)

	switch stmt.SubType {
	case sources.SimpleStatementVarDec:
		for _, variable := range e.evalVariableDeclarationStatement(stmt.VariableDeclaration) {
			e.SetVariable(variable)
		}

	case sources.SimpleStatementExpr:
		e.Eval(stmt.Expression)
	}
}
