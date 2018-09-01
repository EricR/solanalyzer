package sources

// ElementaryTypeName represents an elementary type name in Solidity.
type ElementaryTypeName struct {
	Tokens
	TypeName string
}

// NewElementaryTypeName returns a new instance of ElementaryTypeName.
func NewElementaryTypeName(typeName string) *ElementaryTypeName {
	return &ElementaryTypeName{TypeName: typeName}
}

func (etn *ElementaryTypeName) String() string {
	return etn.TypeName
}
