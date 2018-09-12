package emulator

import "github.com/ericr/solanalyzer/sources"

// Variable represents an emulated variable.
type Variable struct {
	Identifier      string
	TypeName        *sources.TypeName
	StorageLocation string
	Value           *Value
}

// NewVariable returns a new instance of Variable.
func NewVariable(identifier string, typeName *sources.TypeName) *Variable {
	return &Variable{
		Identifier: identifier,
		TypeName:   typeName,
	}
}
