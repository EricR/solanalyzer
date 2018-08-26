package scanner

import "github.com/ericr/solanalyzer/parser"

// EnterModifierList is called when the parser enters a modifier list.
func (s *Scanner) EnterModifierList(ctx *parser.ModifierListContext) {
	isExternal := ctx.ExternalKeyword(0) != nil
	isPublic := ctx.PublicKeyword(0) != nil
	isInternal := ctx.InternalKeyword(0) != nil
	isPrivate := ctx.PrivateKeyword(0) != nil

	// Public visibility is default
	if !isExternal && !isInternal && !isPrivate {
		isPublic = true
	}

	s.function.External = isExternal
	s.function.Public = isPublic
	s.function.Internal = isInternal
	s.function.Private = isPrivate
}

// EnterModifierList is called when the parser enters a modifier invocation.
func (s *Scanner) EnterModifierInvocation(ctx *parser.ModifierInvocationContext) {
	identifier := ctx.Identifier().GetText()

	s.function.Modifiers = append(s.function.Modifiers, identifier)
}
