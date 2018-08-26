package scanner

import (
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
)

// StateVar represents a Solidity state variable.
type StateVar struct {
	Identifier string
	TypeName   string
	Expression string
	Public     bool
	Internal   bool
	Private    bool
	Constant   bool
}

// EnterStateVariableDeclaration is called when the parser enters a state variable declaration.
func (s *Scanner) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	identifier := ctx.Identifier().GetText()
	typeName := ctx.TypeName().GetText()
	isPublic := ctx.PublicKeyword(0) != nil
	isInternal := ctx.InternalKeyword(0) != nil
	isPrivate := ctx.PrivateKeyword(0) != nil
	isConstant := ctx.ConstantKeyword(0) != nil

	// Public visibility is default
	if !isPublic && !isInternal && !isPrivate {
		isPublic = true
	}

	var expression string

	if ctx.Expression() != nil {
		expression = ctx.Expression().GetText()
	}

	if _, ok := s.contract.StateVars[identifier]; !ok {
		s.contract.StateVars[identifier] = &StateVar{
			Identifier: identifier,
			TypeName:   typeName,
			Expression: expression,
			Public:     isPublic,
			Internal:   isInternal,
			Private:    isPrivate,
			Constant:   isConstant,
		}
	}
	s.stateVar = s.contract.StateVars[identifier]
}

// EnterStateVariableDeclaration is called when the parser exits a state variable declaration.
func (s *Scanner) ExitStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	logrus.WithFields(logrus.Fields{
		"contract":   s.contract.Identifier,
		"identifier": s.stateVar.Identifier,
		"type":       s.stateVar.TypeName,
		"public":     s.stateVar.Public,
		"internal":   s.stateVar.Internal,
		"private":    s.stateVar.Private,
		"constant":   s.stateVar.Constant,
	}).Debugf("Scanned state variable declaration")

	s.stateVar = nil
}
