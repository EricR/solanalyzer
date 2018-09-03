package sources

import (
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// ModifierList represents a list of modifiers in Solidity.
type ModifierList struct {
	Tokens
	Invocations     []*ModifierInvocation
	External        bool
	Public          bool
	Internal        bool
	Private         bool
	StateMutability *StateMutability
}

// NewModifierList returns a new instance of ModifierList.
func (s *Source) NewModifierList() *ModifierList {
	modList := &ModifierList{}
	s.AddNode(modList)

	return modList
}

// Visit is called by a visitor.
func (ml *ModifierList) Visit(s *Source, ctx *parser.ModifierListContext) {
	ml.Start = ctx.GetStart()
	ml.Stop = ctx.GetStop()
	ml.External = ctx.ExternalKeyword(0) != nil
	ml.Public = ctx.PublicKeyword(0) != nil
	ml.Internal = ctx.InternalKeyword(0) != nil
	ml.Private = ctx.PrivateKeyword(0) != nil
	ml.StateMutability = s.NewStateMutabilityFromCtxs(ctx.AllStateMutability())

	for _, miCtx := range ctx.AllModifierInvocation() {
		invocation := s.NewModifierInvocation()
		invocation.Visit(s, miCtx.(*parser.ModifierInvocationContext))

		ml.Invocations = append(ml.Invocations, invocation)
	}
}

func (ml *ModifierList) String() string {
	modifiers := []string{}

	if ml.External {
		modifiers = append(modifiers, "external")
	}
	if ml.Public {
		modifiers = append(modifiers, "public")
	}
	if ml.Internal {
		modifiers = append(modifiers, "internal")
	}
	if ml.Private {
		modifiers = append(modifiers, "private")
	}

	for _, invocation := range ml.Invocations {
		modifiers = append(modifiers, invocation.String())
	}

	return strings.Join(modifiers, " ")
}
