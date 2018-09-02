package sources

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericr/solanalyzer/parser"
)

// Source represents a unit of source code in Solidity.
type Source struct {
	FilePath  string
	Pragma    *Pragma
	Imports   []*ImportDirective
	Contracts []*Contract
	tree      *parser.SourceUnitContext
}

// New returns a new instance of Source.
func New(path string, tree *parser.SourceUnitContext) *Source {
	return &Source{
		FilePath: path,
		tree:     tree,
	}
}

// Visit creates a representation of a Solidity syntax tree by creating and then
// "visiting" each object with a parser context. This is used as an alternative
// to ANTLR's visitor pattern, which is not yet fully implemented in the Go
// library.
func (s *Source) Visit() {
	if s.tree.PragmaDirective(0) != nil {
		pragma := NewPragma()
		pragma.Visit(s.tree.PragmaDirective(0).(*parser.PragmaDirectiveContext))

		s.Pragma = pragma
	}

	for _, importCtx := range s.tree.AllImportDirective() {
		importDir := NewImportDirective()
		importDir.Visit(importCtx.(*parser.ImportDirectiveContext))

		s.Imports = append(s.Imports, importDir)
	}

	for _, contractCtx := range s.tree.AllContractDefinition() {
		contract := NewContract()
		contract.Visit(contractCtx.(*parser.ContractDefinitionContext))

		s.Contracts = append(s.Contracts, contract)
	}

}

func (s *Source) String() string {
	return s.FilePath
}

func getText(tree antlr.Tree) string {
	return tree.(antlr.ParseTree).GetText()
}
