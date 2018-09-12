package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalSimpleStatement(stmt *sources.SimpleStatement) {
	defer e.Recover(stmt.Tokens)

	switch stmt.SubType {
	case sources.SimpleStatementVarDec:
		e.evalVariableDeclaration(stmt.VariableDeclaration)

	case sources.SimpleStatementExpr:
		e.Eval(stmt.Expression)
	}
}
