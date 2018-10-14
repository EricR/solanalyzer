package sources

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

// Source represents a unit of source code in Solidity.
type Source struct {
	FilePath      string
	Module        string
	Pragma        *Pragma
	Imports       []*ImportDirective
	Dependencies  []*Source
	DependencyMap map[string]*Source
	Contracts     []*Contract
	tree          *parser.SourceUnitContext
}

// New returns a new instance of Source.
func New(path string, tree *parser.SourceUnitContext) *Source {
	return &Source{
		FilePath:      path,
		Dependencies:  []*Source{},
		DependencyMap: map[string]*Source{},
		tree:          tree,
	}
}

// ParseFile parses a Solidity source file.
func ParseFile(path string) *Source {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("Got error reading file: %s", err)
		return nil
	}

	if filepath.Ext(path) != ".sol" {
		logrus.Errorf("Got error reading file: %s is not a valid source file", path)
		return nil
	}

	return Parse(path, string(bytes))
}

// Parse parses a Solidity source string.
func Parse(path string, source string) *Source {
	logrus.Debugf("Parsing %s", path)

	inputStream := antlr.NewInputStream(source)
	solLexer := parser.NewSolidityLexer(inputStream)
	stream := antlr.NewCommonTokenStream(solLexer, antlr.TokenDefaultChannel)
	solParser := parser.NewSolidityParser(stream)

	solParser.RemoveErrorListeners()
	solParser.AddErrorListener(&ErrorListener{SourceFilePath: path})

	return New(path, solParser.SourceUnit().(*parser.SourceUnitContext))
}

// Visit creates a representation of a Solidity syntax tree by creating and then
// "visiting" each object with a parser context. This is used as an alternative
// to ANTLR's visitor pattern, which is not yet fully implemented in the Go
// library.
func (s *Source) Visit() {
	if s.tree.PragmaDirective(0) != nil {
		pragma := NewPragma(s)
		pragma.Visit(s.tree.PragmaDirective(0).(*parser.PragmaDirectiveContext))

		s.Pragma = pragma
	}

	for _, importCtx := range s.tree.AllImportDirective() {
		importDir := NewImportDirective(s)
		importDir.Visit(importCtx.(*parser.ImportDirectiveContext))

		s.Imports = append(s.Imports, importDir)
	}

	for _, contractCtx := range s.tree.AllContractDefinition() {
		contract := NewContract(s)
		contract.Visit(contractCtx.(*parser.ContractDefinitionContext))

		s.Contracts = append(s.Contracts, contract)
	}
}

func (s *Source) AddDependency(source *Source) {
	if s.DependencyMap[source.FilePath] == nil {
		s.DependencyMap[source.FilePath] = source
		s.Dependencies = append(s.Dependencies, source)
	}
}

func (s *Source) String() string {
	return s.FilePath
}

func getText(tree antlr.Tree) string {
	return tree.(antlr.ParseTree).GetText()
}
