package scanner

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericr/solanalyzer/parser"
)

// Scanner represents a stateful code scanner.
type Scanner struct {
	*parser.BaseSolidityListener
	source         *Source
	contract       *Contract
	stateVar       *StateVar
	function       *Function
	inReturnParams bool
}

// NewScanner returns a new scanner.
func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Scan(source *Source) {
	s.source = source
	antlr.ParseTreeWalkerDefault.Walk(s, source.tree)
}
