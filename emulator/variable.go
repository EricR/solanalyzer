package emulator

import (
	"fmt"
	"github.com/ericr/solanalyzer/sources"
)

// Variable represents an emulated variable.
type Variable struct {
	Identifier      string
	TypeName        *sources.TypeName
	StorageLocation string
	Value           *Value
}

// NewVariable returns a new instance of Variable.
func NewVariable(identifier string, typeName *sources.TypeName, storage string) *Variable {
	return &Variable{
		Identifier:      identifier,
		TypeName:        typeName,
		StorageLocation: storage,
		Value:           NewValue(),
	}
}

func (v *Variable) String() string {
	if v.Value.Expression == nil {
		return fmt.Sprintf("%s %s %s", v.TypeName, v.StorageLocation, v.Identifier)
	}
	return fmt.Sprintf("%s %s %s = %s",
		v.TypeName, v.StorageLocation, v.Identifier, v.Value.Expression)
}

// SetVariable sets a new variable.
func (e *Emulator) SetVariable(variable *Variable) {
	switch variable.StorageLocation {
	case "storage":
		e.storage = append(e.storage, variable)
		e.storageMap[variable.Identifier] = variable

	case "memory":
		frame := e.stack.CurrentFrame()

		frame.LocalVariables = append(frame.LocalVariables, variable)
		frame.LocalVariablesMap[variable.Identifier] = variable

	default:
		panic("Unknown variable storage location")
	}
}

// MustSetVariable must find and set a variable, otherwise the emulator panics.
func (e *Emulator) MustSetVariable(name string, value *Value) {
	e.MustFindVariable(name).Value = value
}

// FindVariable attempts to find and return a variable.
func (e *Emulator) FindVariable(name string) *Variable {
	frame := e.stack.CurrentFrame()

	memoryVar, memoryFound := frame.LocalVariablesMap[name]
	storageVar, _ := e.storageMap[name]

	if memoryFound {
		return memoryVar
	}
	return storageVar
}

// MustFindVariable must return a variable, otherwise the emulator panics.
func (e *Emulator) MustFindVariable(name string) *Variable {
	variable := e.FindVariable(name)

	if variable == nil {
		panic("Undeclared variable")
	}

	return variable
}
