package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Pragma represents a pragma in Solidity.
type Pragma struct {
	Tokens
	Name  string
	Value string
}

// NewPragma returns a new instance of Pragma.
func NewPragma() *Pragma {
	return &Pragma{}
}

// Visit is called by a visitor.
func (p *Pragma) Visit(ctx *parser.PragmaDirectiveContext) {
	p.Start = ctx.GetStart()
	p.Stop = ctx.GetStop()
	p.Name = ctx.PragmaName().GetText()
	p.Value = ctx.PragmaValue().GetText()
}

func (p *Pragma) String() string {
	return fmt.Sprintf("pragma %s %s", p.Name, p.Value)
}
