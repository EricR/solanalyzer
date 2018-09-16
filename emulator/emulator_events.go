package emulator

import "github.com/ericr/solanalyzer/sources"

// Event is an event emitted by the emulator.
type Event struct {
	Source                *sources.Source
	Contract              *sources.Contract
	Function              *sources.Function
	CallerFunction        *sources.Function
	Expression            *sources.Expression
	Expressions           []*sources.Expression
	FunctionCallArguments *sources.FunctionCallArguments
}

// OnEvent adds a function that will be executed on an event.
func (e Emulator) OnEvent(subject string, handler func(*Event)) {
	e.eventHandlers[subject] = append(e.eventHandlers[subject], handler)
}

// Emit emits an event.
func (e *Emulator) Emit(subject string, event *Event) {
	for _, handler := range e.eventHandlers[subject] {
		handler(event)
	}
}
