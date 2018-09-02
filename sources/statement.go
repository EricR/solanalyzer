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
	SubType         int
	If              *IfStatement
	While           *WhileStatement
	For             *ForStatement
	Block           *Block
	DoWhile         *DoWhileStatement
	ReturnStatement *ReturnStatement
}

// NewStatement returns a new instance of Statement.
func NewStatement() *Statement {
	return &Statement{}
}

// Visit is called by a visitor.
func (s *Statement) Visit(ctx *parser.StatementContext) {
	s.Start = ctx.GetStart()
	s.Stop = ctx.GetStop()

	switch {
	case ctx.IfStatement() != nil:
		ifStatement := NewIfStatement()
		ifStatement.Visit(ctx.IfStatement().(*parser.IfStatementContext))

		s.SubType = StatementIf
		s.If = ifStatement

	case ctx.WhileStatement() != nil:
		whileStatement := NewWhileStatement()
		whileStatement.Visit(ctx.WhileStatement().(*parser.WhileStatementContext))

		s.SubType = StatementWhile
		s.While = whileStatement

	case ctx.ForStatement() != nil:
		forStatement := NewForStatement()
		forStatement.Visit(ctx.ForStatement().(*parser.ForStatementContext))

		s.SubType = StatementFor
		s.For = forStatement

	case ctx.Block() != nil:
		block := NewBlock()
		block.Visit(ctx.Block().(*parser.BlockContext))

		s.SubType = StatementBlock
		s.Block = block

	case ctx.InlineAssemblyStatement() != nil:
		s.SubType = StatementInlineAssembly
		// TODO

	case ctx.DoWhileStatement() != nil:
		doWhile := NewDoWhileStatement()
		doWhile.Visit(ctx.DoWhileStatement().(*parser.DoWhileStatementContext))

		s.SubType = StatementDoWhile
		s.DoWhile = doWhile

	case ctx.ContinueStatement() != nil:
		s.SubType = StatementContinue

	case ctx.BreakStatement() != nil:
		s.SubType = StatementBreak

	case ctx.ReturnStatement() != nil:
		returnStatement := NewReturnStatement()
		returnStatement.Visit(ctx.ReturnStatement().(*parser.ReturnStatementContext))

		s.SubType = StatementReturn
		s.ReturnStatement = returnStatement

	case ctx.ThrowStatement() != nil:
		s.SubType = StatementThrow

	case ctx.EmitStatement() != nil:
		s.SubType = StatementEmit

	case ctx.SimpleStatement() != nil:
		s.SubType = StatementSimple
	}
}

func (s *Statement) String() string {
	return "TODO"
}
