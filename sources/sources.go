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
	Nodes     []Node
	tree      *parser.SourceUnitContext
}

// Node represents a node of the Solidity syntax tree.
type Node interface{}

// New returns a new instance of Source.
func New(path string, tree *parser.SourceUnitContext) *Source {
	return &Source{
		FilePath: path,
		Nodes:    []Node{},
		tree:     tree,
	}
}

// Visit creates a representation of a Solidity syntax tree by creating and then
// "visiting" each object with a parser context. This is used as an alternative
// to ANTLR's visitor pattern, which is not yet fully implemented in the Go
// library.
func (s *Source) Visit() {
	if s.tree.PragmaDirective(0) != nil {
		pragma := s.NewPragma()
		pragma.Visit(s, s.tree.PragmaDirective(0).(*parser.PragmaDirectiveContext))

		s.Pragma = pragma
	}

	for _, importCtx := range s.tree.AllImportDirective() {
		importDir := s.NewImportDirective()
		importDir.Visit(s, importCtx.(*parser.ImportDirectiveContext))

		s.Imports = append(s.Imports, importDir)
	}

	for _, contractCtx := range s.tree.AllContractDefinition() {
		contract := s.NewContract()
		contract.Visit(s, contractCtx.(*parser.ContractDefinitionContext))

		s.Contracts = append(s.Contracts, contract)
	}

}

// AddNode adds a source tree node to a flat list for easier traversal.
func (s *Source) AddNode(node Node) {
	s.Nodes = append(s.Nodes, node)
}

func (s *Source) String() string {
	return s.FilePath
}

func getText(tree antlr.Tree) string {
	return tree.(antlr.ParseTree).GetText()
}
