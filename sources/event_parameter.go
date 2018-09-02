package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// EvenParameter represents an event parameter in Solidity.
type EventParameter struct {
	Tokens
	Event      *Event
	TypeName   *TypeName
	Indexed    bool
	Identifier string
}

// NewEventParameter returns a new instance of EventParameter.
func NewEventParameter(event *Event) *EventParameter {
	return &EventParameter{Event: event}
}

// Visit is called by a visitor.
func (ep *EventParameter) Visit(ctx *parser.EventParameterContext) {
	ep.Start = ctx.GetStart()
	ep.Stop = ctx.GetStop()

	typeName := NewTypeName()
	typeName.Visit(ctx.TypeName().(*parser.TypeNameContext))

	ep.TypeName = typeName
	ep.Identifier = ctx.Identifier().GetText()

	if ctx.IndexedKeyword() != nil {
		ep.Indexed = true
	}
}

func (ep *EventParameter) String() string {
	str := fmt.Sprintf("%s ", ep.TypeName)

	if ep.Indexed {
		str += " indexed"
	}

	if ep.Identifier != "" {
		str += fmt.Sprintf(" %s", ep.Identifier)
	}

	return str
}
