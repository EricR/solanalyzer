package sources

import (
	"github.com/ericr/solanalyzer/parser"
	"strconv"
)

const (
	// ElementaryTypeNameUnknown represents an unknown type.
	ElementaryTypeNameUnknown = iota
	// ElementaryTypeNameInt represents an integer.
	ElementaryTypeNameInt
	// ElementaryTypeNameUint represents an unsigned integer.
	ElementaryTypeNameUint
	// ElementaryTypeNameAddress represents an address.
	ElementaryTypeNameAddress
	// ElementaryTypeNameBool represents a boolean.
	ElementaryTypeNameBool
	// ElementaryTypeNameString represents a string.
	ElementaryTypeNameString
	// ElementaryTypeNameVar represents a variable.
	ElementaryTypeNameVar
	// ElementaryTypeNameBytes represents bytes.
	ElementaryTypeNameBytes
	// ElementaryTypeNameDynamicBytes represents a dynamic byte array.
	ElementaryTypeNameDynamicBytes
	// ElementaryTypeNameFixed represents a fixed type.
	ElementaryTypeNameFixed
	// ElementaryTypeNameUfixed represents an unfixed type.
	ElementaryTypeNameUfixed
)

// ElementaryTypeName represents an elementary type name in Solidity.
type ElementaryTypeName struct {
	Tokens
	SubType int
	Size    int
	Text    string
}

// NewElementaryTypeName returns a new instance of ElementaryTypeName.
func NewElementaryTypeName() *ElementaryTypeName {
	return &ElementaryTypeName{}
}

// Visit is called by a visitor.
func (etn *ElementaryTypeName) Visit(ctx *parser.ElementaryTypeNameContext) {
	etn.Start = ctx.GetStart()
	etn.Stop = ctx.GetStop()
	etn.Text = ctx.GetText()

	switch {
	case ctx.Int() != nil:
		etn.SubType = ElementaryTypeNameInt
		etn.Size = mustParseSize(etn.Text, 3)

	case ctx.Uint() != nil:
		etn.SubType = ElementaryTypeNameUint
		etn.Size = mustParseSize(etn.Text, 4)

	case ctx.GetText() == "address":
		etn.SubType = ElementaryTypeNameAddress

	case ctx.GetText() == "bool":
		etn.SubType = ElementaryTypeNameBool

	case ctx.GetText() == "string":
		etn.SubType = ElementaryTypeNameString

	case ctx.GetText() == "var":
		etn.SubType = ElementaryTypeNameVar

	case ctx.GetText() == "byte":
		etn.SubType = ElementaryTypeNameBytes
		etn.Size = 1

	case ctx.GetText() == "bytes":
		etn.SubType = ElementaryTypeNameDynamicBytes

	case ctx.Byte() != nil:
		etn.SubType = ElementaryTypeNameBytes
		etn.Size = mustParseSize(etn.Text, 5)

	case ctx.Fixed() != nil:
		etn.SubType = ElementaryTypeNameFixed

	case ctx.Ufixed() != nil:
		etn.SubType = ElementaryTypeNameUfixed

	default:
		panic("Unknown elementary type name")
	}
}

func (etn *ElementaryTypeName) String() string {
	return etn.Text
}

// Equal evaluates the equality of two elementary type names.
func (etn *ElementaryTypeName) Equal(b *ElementaryTypeName) bool {
	// uints can be a subset of ints
	if etn.SubType == ElementaryTypeNameUint && b.SubType == ElementaryTypeNameInt {
		return true
	}

	return etn.SubType == b.SubType
}

func mustParseSize(str string, offset int) int {
	switch str {
	case "int":
		str = "int256"
	case "uint":
		str = "uint256"
	}

	i, err := strconv.Atoi(str[offset:])
	if err != nil {
		panic("Failed to parse type size")
	}

	return i
}
