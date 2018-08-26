package scanner

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
	"strings"
)

// Function represents a Solidity function.
type Function struct {
	Tokens     *Tokens
	Identifier string
	Params     []string
	Modifiers  []string
	Returns    []string
	External   bool
	Public     bool
	Internal   bool
	Private    bool
}

// NewFunction returns a new instance of Function.
func NewFunction(identifier string) *Function {
	return &Function{
		Tokens:     &Tokens{},
		Identifier: identifier,
		Modifiers:  []string{},
	}
}

// GetSignature returns a function's signature composed of an identifier
// and parameter list.
func (f *Function) GetSignature() string {
	return fmt.Sprintf("%s(%s)", f.Identifier, strings.Join(f.Params, ","))
}

// GetFullSignature returns a function's signature composed of an identifier,
// parameter list, and return list.
func (f *Function) GetFullSignature() string {
	sig := f.GetSignature()

	if len(f.Returns) > 0 {
		sig = fmt.Sprintf("%s (%s)", sig, strings.Join(f.Returns, ","))
	}

	return sig
}

// EnterFunctionDefinition is called when the parser enters a function definition.
func (s *Scanner) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	var identifier string

	if ctx.Identifier() != nil {
		identifier = ctx.Identifier().GetText()
	}

	function := NewFunction(identifier)
	function.Tokens.Start = ctx.GetStart()
	function.Tokens.Stop = ctx.GetStop()

	s.contract.Functions = append(s.contract.Functions, function)
	s.function = function
}

// ExitFunctionDefinition is called when the parser exits a function definition.
func (s *Scanner) ExitFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	logrus.WithFields(logrus.Fields{
		"contract":   s.contract.Identifier,
		"identifier": s.function.Identifier,
		"params":     s.function.Params,
		"modifiers":  s.function.Modifiers,
		"signature":  s.function.GetFullSignature(),
		"returns":    s.function.Returns,
		"external":   s.function.External,
		"public":     s.function.Public,
		"internal":   s.function.Internal,
		"private":    s.function.Private,
	}).Debugf("Scanned function definition")

	s.function = nil
}
