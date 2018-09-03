package sources

import (
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// StateMutability represents a state mutability keyword in Solidity.
type StateMutability struct {
	Pure     bool
	Constant bool
	View     bool
	Payable  bool
}

// NewStateMutability returns an instance of StateMutability.
func (s *Source) NewStateMutability() *StateMutability {
	stMut := &StateMutability{}
	s.AddNode(stMut)

	return stMut
}

// NewStateMutabilityFromCtxs returns an instance of StateMutability given a
// parser context.
func (s *Source) NewStateMutabilityFromCtxs(ctxs []parser.IStateMutabilityContext) *StateMutability {
	sm := s.NewStateMutability()

	for _, mutabiltiy := range ctxs {
		smCtx := mutabiltiy.(*parser.StateMutabilityContext)

		if smCtx.PureKeyword() != nil {
			sm.Pure = true
		}
		if smCtx.ConstantKeyword() != nil {
			sm.Constant = true
		}
		if smCtx.ViewKeyword() != nil {
			sm.View = true
		}
		if smCtx.PayableKeyword() != nil {
			sm.Payable = true
		}
	}

	return sm
}

func (sm *StateMutability) String() string {
	mutability := []string{}

	if sm.Pure {
		mutability = append(mutability, "pure")
	}
	if sm.Constant {
		mutability = append(mutability, "constant")
	}
	if sm.View {
		mutability = append(mutability, "view")
	}
	if sm.Payable {
		mutability = append(mutability, "payable")
	}

	return strings.Join(mutability, " ")
}
