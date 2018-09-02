package sources

import "github.com/ericr/solanalyzer/parser"

const (
	// If
	StatementIf int = iota
	// While
	StatementWhile
	// For
	StatementFor
	// Block
	StatementBlock
	// Inline assembly
	StatementInlineAssembly
	// Do while
	StatementDoWhile
	// Continue
	StatementContinue
	// Break
	StatementBreak
	// Return
	StatementReturn
	// Throw
	StatementThrow
	// Emit
	StatementEmit
	// Simple
	StatementSimple
)

// Statement represents a statement in Solidity.
type Statement struct {
	Tokens
	SubType     int
	IfStatement *IfStatement
}

func NewStatement() *Statement {
	return &Statement{}
}

func (s *Statement) Visit(ctx *parser.StatementContext) {
	s.Start = ctx.GetStart()
	s.Stop = ctx.GetStop()

	switch {
	case ctx.IfStatement() != nil:
		ifStatement := NewIfStatement()
		ifStatement.Visit(ctx.IfStatement().(*parser.IfStatementContext))

		s.SubType = StatementIf
		s.IfStatement = ifStatement
	}
}

func (s *Statement) String() string {
	return "TODO"
}
