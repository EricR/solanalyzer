package scanner

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Source represents a unit of Solidity source code.
type Source struct {
	FilePath  string
	Pragma    *Pragma
	Imports   []*Import
	Contracts map[string]*Contract
	tree      antlr.Tree
}

// Tokens store the start and stop of a parse tree node.
// This is useful for later reporting line rows and columns.
type Tokens struct {
	Start antlr.Token
	Stop  antlr.Token
}

// NewSource returns a new instance of Source.
func NewSource(path string, tree antlr.Tree) *Source {
	return &Source{
		FilePath:  path,
		Contracts: map[string]*Contract{},
		tree:      tree,
	}
}
