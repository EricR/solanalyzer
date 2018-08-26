package scanner

import (
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
)

// EnterConstructorDefinition is called when the parser enters a constructor definition.
func (s *Scanner) EnterConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	constructor := NewFunction("")

	s.contract.Constructor = constructor
	s.function = constructor
}

// ExitConstructorDefinition is called when the parser exits a constructor definition.
func (s *Scanner) ExitConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	logrus.WithFields(logrus.Fields{
		"contract":  s.contract.Identifier,
		"params":    s.function.Params,
		"modifiers": s.function.Modifiers,
		"external":  s.function.External,
		"public":    s.function.Public,
		"internal":  s.function.Internal,
		"private":   s.function.Private,
	}).Debug("Scanned constructor definition")

	s.function = nil
}
