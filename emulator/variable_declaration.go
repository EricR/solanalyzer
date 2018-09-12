package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalVariableDeclaration(stmt *sources.VariableDeclarationStatement) {
	defer e.Recover(stmt.Tokens)

	switch {
	case len(stmt.Identifiers) > 0:
		if stmt.Expression == nil {
			panic("Assignment necessary for type detection when using var")
		}

		vals := e.Eval(stmt.Expression)
		if len(vals) != len(stmt.Identifiers) {
			panic("Unbalanced variable declaration")
		}

		for i, identifier := range stmt.Identifiers {
			variable := NewVariable(identifier, vals[i].InferredType())
			variable.Value = vals[i]

			e.localVars[identifier] = variable
		}

	case stmt.VariableDeclaration != nil:
		varDec := stmt.VariableDeclaration
		typeName := varDec.TypeName
		varType := typeName.SubType == sources.TypeNameElementary &&
			typeName.Elementary.SubType == sources.ElementaryTypeNameVar
		variable := NewVariable(varDec.Identifier, typeName)

		if varType && stmt.Expression == nil {
			panic("Assignment necessary for type detection when using var")
		}

		if stmt.Expression != nil {
			vals := e.Eval(stmt.Expression)

			if len(vals) != 1 {
				panic("Unbalanced variable declaration")
			}

			variable.Value = vals[0]

			if varType {
				variable.TypeName = variable.Value.InferredType()
			}
		}

		if variable.StorageLocation != "" {
			variable.StorageLocation = varDec.StorageLocation
		}

		e.localVars[varDec.Identifier] = variable
	}
}
