package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// Event represents an event in Solidity.
type Event struct {
	Tokens
	Contract   *Contract
	Identifier string
	Anonymous  bool
	Parameters []*EventParameter
}

// NewEvent returns a new instance of Event.
func NewEvent() *Event {
	return &Event{}
}

// Visit is called by a visitor.
func (e *Event) Visit(ctx *parser.EventDefinitionContext) {
	e.Start = ctx.GetStart()
	e.Stop = ctx.GetStop()

	e.Identifier = ctx.Identifier().GetText()
	e.Anonymous = ctx.AnonymousKeyword() != nil

	pramList := ctx.EventParameterList().(*parser.EventParameterListContext)

	for _, paramCtx := range pramList.AllEventParameter() {
		param := NewEventParameter(e)
		param.Visit(paramCtx.(*parser.EventParameterContext))

		e.Parameters = append(e.Parameters, param)
	}
}

func (e *Event) String() string {
	paramStrs := []string{}

	for _, param := range e.Parameters {
		paramStrs = append(paramStrs, param.String())
	}

	if e.Anonymous {
		return fmt.Sprintf("event %s (%s) anonymous",
			e.Identifier, strings.Join(paramStrs, ", "))
	}
	return fmt.Sprintf("event %s (%s)",
		e.Identifier, strings.Join(paramStrs, ", "))
}
