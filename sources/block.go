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
func NewBlock() *Block {
	return &Block{
		Statements: []*Statement{},
	}
}

// Visit is called by a visitor.
func (b *Block) Visit(ctx *parser.BlockContext) {
	b.Start = ctx.GetStart()
	b.Stop = ctx.GetStop()

	for _, stmtCtx := range ctx.AllStatement() {
		statement := NewStatement()
		statement.Visit(stmtCtx.(*parser.StatementContext))

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
