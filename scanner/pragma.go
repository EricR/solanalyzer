package scanner

import (
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
)

// Pragma represents a Solidity pragma.
type Pragma struct {
	Tokens  *Tokens
	Version string
}

// EnterPragmaValue is called when the parser enters a pragma.
func (s *Scanner) EnterPragmaValue(ctx *parser.PragmaValueContext) {
	s.source.Pragma = &Pragma{
		Tokens: &Tokens{
			Start: ctx.GetStart(),
			Stop:  ctx.GetStop(),
		},
		Version: ctx.GetText(),
	}
}

// ExitPragmaValue is called when the parser exits a pragma.
func (s *Scanner) ExitPragmaValue(ctx *parser.PragmaValueContext) {
	logrus.WithFields(logrus.Fields{
		"version": s.source.Pragma.Version,
	}).Debugf("Scanned pragma")
}
