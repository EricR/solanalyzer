package sources

import "github.com/ericr/solanalyzer/parser"

const (
	// ElementaryTypeNameInt represents an integer.
	ElementaryTypeNameInt = iota
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
	// ElementaryTypeNameByte represents a byte.
	ElementaryTypeNameByte
	// ElementaryTypeNameBytes represents bytes.
	ElementaryTypeNameBytes
	// ElementaryTypeNameFixed represents a fixed type.
	ElementaryTypeNameFixed
	// ElementaryTypeNameUfixed represents an unfixed type.
	ElementaryTypeNameUfixed
)

// ElementaryTypeName represents an elementary type name in Solidity.
type ElementaryTypeName struct {
	Tokens
	SubType int
	Text    string
	Int     string
	Uint    string
	Byte    string
	Fixed   string
	Ufixed  string
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
		etn.Int = ctx.Int().GetText()
		etn.SubType = ElementaryTypeNameInt

	case ctx.Uint() != nil:
		etn.Uint = ctx.Uint().GetText()
		etn.SubType = ElementaryTypeNameUint

	case ctx.GetText() == "address":
		etn.SubType = ElementaryTypeNameAddress

	case ctx.GetText() == "bool":
		etn.SubType = ElementaryTypeNameBool

	case ctx.GetText() == "string":
		etn.SubType = ElementaryTypeNameString

	case ctx.GetText() == "var":
		etn.SubType = ElementaryTypeNameVar

	case ctx.GetText() == "byte":
		etn.SubType = ElementaryTypeNameByte

	case ctx.Byte() != nil:
		etn.Byte = ctx.Byte().GetText()
		etn.SubType = ElementaryTypeNameBytes

	case ctx.Fixed() != nil:
		etn.Fixed = ctx.Fixed().GetText()
		etn.SubType = ElementaryTypeNameFixed

	case ctx.Ufixed() != nil:
		etn.Ufixed = ctx.Ufixed().GetText()
		etn.SubType = ElementaryTypeNameUfixed

	default:
		panic("Unknown elementary type name")
	}
}

func (etn *ElementaryTypeName) String() string {
	return etn.Text
}
