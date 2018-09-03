package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// EventParameter represents an event parameter in Solidity.
type EventParameter struct {
	Tokens
	Event      *Event
	TypeName   *TypeName
	Indexed    bool
	Identifier string
}

// NewEventParameter returns a new instance of EventParameter.
func (s *Source) NewEventParameter(event *Event) *EventParameter {
	param := &EventParameter{Event: event}
	s.AddNode(param)

	return param
}

// Visit is called by a visitor.
func (ep *EventParameter) Visit(s *Source, ctx *parser.EventParameterContext) {
	ep.Start = ctx.GetStart()
	ep.Stop = ctx.GetStop()

	typeName := s.NewTypeName()
	typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

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
