package sources

// FunctionTypeName represents a function type name in Solidity.
type FunctionTypeName struct {
	Tokens
	Parameters      []*FunctionTypeParameter
	Internal        bool
	External        bool
	StateMutability *StateMutability
	Returns         []*FunctionTypeParameter
}
