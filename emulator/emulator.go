package emulator

import (
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

// Emulator is a Solidity emulator.
type Emulator struct {
	ErrorCount    uint
	source        *sources.Source
	eventHandlers map[string][]func(*Event)
	stack         *Stack
	imports       []*sources.ImportDirective
	importsMap    map[string]*sources.ImportDirective
	structs       []*sources.Struct
	structsMap    map[string]*sources.Struct
	functions     []*sources.Function
	functionsMap  map[string]*sources.Function
	storage       []*Variable
	storageMap    map[string]*Variable
}

// New returns a new instance of Emulator.
func New(source *sources.Source) *Emulator {
	return &Emulator{
		source:        source,
		eventHandlers: map[string][]func(*Event){},
		stack:         &Stack{},
		imports:       []*sources.ImportDirective{},
		structs:       []*sources.Struct{},
		structsMap:    map[string]*sources.Struct{},
		functions:     []*sources.Function{},
		functionsMap:  map[string]*sources.Function{},
		storage:       []*Variable{},
		storageMap:    map[string]*Variable{},
	}
}

// Run runs the emulator.
func (e *Emulator) Run() {
	for _, contract := range e.source.Contracts {
		e.evalContractDefinition(contract)
	}

	if e.ErrorCount > 0 {
		logrus.Warnf("Finished emulation of %s with %d error(s)",
			e.source, e.ErrorCount)
	}
}

// Reset resets the emulator's state.
func (e *Emulator) Reset() {
	e.ResetStructs()
	e.ResetFunctions()
	e.ResetStack()
	e.ResetStorage()
}

// Recover acts as a recovery strategy when the emulator panics.
func (e *Emulator) Recover(tokens sources.Tokens) {
	if r := recover(); r != nil {
		e.ErrorCount++

		logrus.Errorf("Error evaluating %s:%d:%d: %s",
			e.source.FilePath, tokens.Start.GetLine(), tokens.Start.GetColumn(), r)
		logrus.Debugf("Stack trace for debugging:\n\n%s", debug.Stack())
	}
}

// ResetStructs resets struct declarations.
func (e *Emulator) ResetStructs() {
	e.structs = []*sources.Struct{}
	e.structsMap = map[string]*sources.Struct{}
}

// ResetFunctions resets function definitions.
func (e *Emulator) ResetFunctions() {
	e.functions = []*sources.Function{}
	e.functionsMap = map[string]*sources.Function{}
}

// ResetStorage resets the storage.
func (e *Emulator) ResetStorage() {
	e.storage = []*Variable{}
	e.storageMap = map[string]*Variable{}
}

// ResetMemory resets the stack.
func (e *Emulator) ResetStack() {
	e.stack = &Stack{}
}

// AddStructDeclaration adds a struct declaration.
func (e *Emulator) AddStructDeclaration(_struct *sources.Struct) {
	e.structs = append(e.structs, _struct)
	e.structsMap[_struct.Identifier] = _struct
}

// AddFunctionDefinition adds a function definition.
func (e *Emulator) AddFunctionDefinition(function *sources.Function) {
	e.functions = append(e.functions, function)
	e.functionsMap[function.Identifier] = function
}

// ResolveIdentifier attempts to resolve an identifier's value.
func (e *Emulator) ResolveIdentifier(name string) *Value {
	variable := e.FindVariable(name)

	if variable != nil {
		return variable.Value
	}

	return nil
}

// MustResolveIdentifier must return a value, otherwise the emulator panics.
func (e *Emulator) MustResolveIdentifier(name string) *Value {
	value := e.ResolveIdentifier(name)

	if value == nil {
		panic("Could not resolve identifier")
	}

	return value
}
