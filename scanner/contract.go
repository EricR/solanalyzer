package scanner

import (
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
)

// Contract represents a Solidity contract.
type Contract struct {
	Identifier  string
	Inheritence []string
	StateVars   map[string]*StateVar
	Functions   []*Function
	Constructor *Function
}

// NewContract returns a new instance of Contract.
func NewContract(identifier string) *Contract {
	return &Contract{
		Identifier:  identifier,
		Inheritence: []string{},
		StateVars:   map[string]*StateVar{},
	}
}

// EnterContractDefinition is called when the parser enters a contract definition.
func (s *Scanner) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	identifier := ctx.Identifier().GetText()

	if _, ok := s.source.Contracts[identifier]; !ok {
		s.source.Contracts[identifier] = NewContract(identifier)
	}
	s.contract = s.source.Contracts[identifier]

	for _, identifier := range ctx.AllInheritanceSpecifier() {
		s.contract.Inheritence = append(s.contract.Inheritence, identifier.GetText())
	}

	logrus.WithFields(logrus.Fields{
		"identifier": identifier,
	}).Debugf("Scanning contract definition")
}

// ExitContractDefinition is called when the parser exits a contract definition.
func (s *Scanner) ExitContractDefinition(ctx *parser.ContractDefinitionContext) {
	s.contract = nil
}
