package emulator

import (
	"github.com/ericr/solanalyzer/sources"
	"math/big"
)

func (e *Emulator) evalVariableDeclarationStatement(stmt *sources.VariableDeclarationStatement) []*Variable {
	defer e.Recover(stmt.Tokens)

	variables := []*Variable{}

	switch {
	case len(stmt.Identifiers) > 0:
		if stmt.Expression == nil {
			panic("Assignment necessary for type detection when using var")
		}

		values := e.Eval(stmt.Expression)
		if len(values) != len(stmt.Identifiers) {
			panic("Unbalanced variable declaration")
		}

		for i, identifier := range stmt.Identifiers {
			value := values[i]

			// Inferred types are stored in memory
			variable := NewVariable(identifier, value.InferredTypeName(), "memory")
			variable.Value = value

			variables = append(variables, variable)
		}

	case stmt.VariableDeclaration != nil:
		variables = append(variables, e.evalVariableDeclaration(stmt.VariableDeclaration, stmt.Expression))

	case len(stmt.VariableDeclarationList) > 0:
		for _, varDec := range stmt.VariableDeclarationList {
			variables = append(variables, e.evalVariableDeclaration(varDec, stmt.Expression))
		}
	}

	return variables
}

func (e *Emulator) evalVariableDeclaration(varDec *sources.VariableDeclaration, expr *sources.Expression) *Variable {
	var value *Value

	typeName := varDec.TypeName
	varType := typeName.SubType == sources.TypeNameElementary &&
		typeName.Elementary.SubType == sources.ElementaryTypeNameVar

	if varType && expr == nil {
		panic("Assignment necessary for type detection when using var")
	}

	if expr != nil {
		values := e.Eval(expr)

		if len(values) != 1 {
			panic("Unbalanced variable declaration")
		}

		value = values[0]

		if varType {
			typeName = value.InferredTypeName()
		}
	}

	// Complex types default to storage
	if varDec.StorageLocation == "" {
		if typeName.IsComplex() {
			varDec.StorageLocation = "storage"
		} else {
			varDec.StorageLocation = "memory"
		}
	}

	if value == nil {
		value = defaultValueOfType(typeName)
	}

	variable := NewVariable(varDec.Identifier, typeName, varDec.StorageLocation)
	variable.Value = value

	return variable
}

func defaultValueOfType(typeName *sources.TypeName) *Value {
	expr := sources.NewPrimaryExpression()

	switch typeName.SubType {
	case sources.TypeNameElementary:
		switch typeName.Elementary.SubType {
		case sources.ElementaryTypeNameInt, sources.ElementaryTypeNameUint:
			expr.SubType = sources.ExpressionPrimaryNumber
			expr.Integer = big.NewInt(0)

		case sources.ElementaryTypeNameAddress:
			expr.SubType = sources.ExpressionPrimaryHex
			expr.Hex = "0x0000000000000000000000000000000000000000"

		case sources.ElementaryTypeNameBool:
			expr.SubType = sources.ExpressionPrimaryBoolean
			expr.Boolean = false
		}
	}

	value := NewValue()
	value.Expression.SubType = sources.ExpressionPrimary
	value.Expression.Primary = expr
	value.Solved = true

	return value
}
