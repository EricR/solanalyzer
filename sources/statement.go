package sources

import "github.com/ericr/solanalyzer/parser"

const (
	// StatementIf represents an if statement.
	StatementIf int = iota
	// StatementWhile represents a while statement.
	StatementWhile
	// StatementFor represents a for statement.
	StatementFor
	// StatementBlock represents a block.
	StatementBlock
	// StatementInlineAssembly represents inline assembly.
	StatementInlineAssembly
	// StatementDoWhile represents a do while statement.
	StatementDoWhile
	// StatementContinue represents a continue.
	StatementContinue
	// StatementBreak represents a break.
	StatementBreak
	// StatementReturn represents a return.
	StatementReturn
	// StatementThrow represents a throw.
	StatementThrow
	// StatementEmit represents emit.
	StatementEmit
	// StatementSimple represents a simple statement.
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
	SimpleStatement *SimpleStatement
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
		simpleStatement := NewSimpleStatement()
		simpleStatement.Visit(ctx.SimpleStatement().(*parser.SimpleStatementContext))

		s.SimpleStatement = simpleStatement
		s.SubType = StatementSimple

	default:
		panic("Unknown type of statement")
	}
}

func (s *Statement) String() string {
	switch s.SubType {
	case StatementIf:
		return s.If.String()

	case StatementWhile:
		return s.While.String()

	case StatementFor:
		return s.For.String()

	case StatementBlock:
		return s.Block.String()

	case StatementInlineAssembly:
		// TODO
		return "TODO"

	case StatementDoWhile:
		return s.DoWhile.String()

	case StatementContinue:
		return "continue"

	case StatementBreak:
		return "break"

	case StatementReturn:
		return s.ReturnStatement.String()

	case StatementThrow:
		return "throw"

	case StatementEmit:
		return "emit"

	case StatementSimple:
		return s.SimpleStatement.String()

	default:
		panic("Unknown statement type")
	}

	return "unknown statement type"
}
