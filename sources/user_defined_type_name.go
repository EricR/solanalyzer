package sources

import "strings"

// UserDefinedTypeName represents a user-defined type name in Solidity.
type UserDefinedTypeName []string

// NewUserDefinedTypeName returns a new instance of UserDefinedTypeName.
func NewUserDefinedTypeName() *UserDefinedTypeName {
	return &UserDefinedTypeName{}
}

// Add adds an identifier to an instance of UserDefinedTypeName.
func (ut *UserDefinedTypeName) Add(identifier string) {
	*ut = append(*ut, identifier)
}

func (ut *UserDefinedTypeName) String() string {
	return strings.Join(*ut, ".")
}
