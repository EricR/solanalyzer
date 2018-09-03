package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// Block represents a block in Solidity.
type Block struct {
	Tokens
	Statements []*Statement
}

// NewBlock returns a new instance of Block.
func (s *Source) NewBlock() *Block {
	block := &Block{
		Statements: []*Statement{},
	}
	s.AddNode(block)

	return block
}

// Visit is called by a visitor.
func (b *Block) Visit(s *Source, ctx *parser.BlockContext) {
	b.Start = ctx.GetStart()
	b.Stop = ctx.GetStop()

	for _, stmtCtx := range ctx.AllStatement() {
		statement := s.NewStatement()
		statement.Visit(s, stmtCtx.(*parser.StatementContext))

		b.Statements = append(b.Statements, statement)
	}
}

func (b *Block) String() string {
	statementStrs := []string{}

	for _, statement := range b.Statements {
		statementStrs = append(statementStrs, statement.String())
	}

	return fmt.Sprintf("{%s}", strings.Join(statementStrs, " "))
}
