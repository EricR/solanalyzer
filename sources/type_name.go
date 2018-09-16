package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

const (
	// TypeNameUnknown represents an unknown type.
	TypeNameUnknown int = iota
	// TypeNameElementary represents an elementary type.
	TypeNameElementary
	// TypeNameUserDefined represents a user-defined type.
	TypeNameUserDefined
	// TypeNameMapping represents a mapping type.
	TypeNameMapping
	// TypeNameArray represents an array type.
	TypeNameArray
	// TypeNameFunction represents a function type.
	TypeNameFunction
)

// TypeName represents the name of a type in Solidity.
type TypeName struct {
	Tokens
	SubType     int
	Elementary  *ElementaryTypeName
	UserDefined *UserDefinedTypeName
	Mapping     *Mapping
	TypeName    *TypeName
	Expression  *Expression
	Function    *FunctionTypeName
}

// NewTypeName returns a new instance of TypeName.
func NewTypeName() *TypeName {
	return &TypeName{}
}

// Visit is called by a visitor.
func (tn *TypeName) Visit(ctx *parser.TypeNameContext) {
	tn.Start = ctx.GetStart()
	tn.Stop = ctx.GetStop()

	switch {
	case ctx.ElementaryTypeName() != nil:
		etn := NewElementaryTypeName()
		etn.Visit(ctx.ElementaryTypeName().(*parser.ElementaryTypeNameContext))

		tn.SubType = TypeNameElementary
		tn.Elementary = etn

	case ctx.UserDefinedTypeName() != nil:
		udCtx := ctx.UserDefinedTypeName().(*parser.UserDefinedTypeNameContext)
		userDefined := NewUserDefinedTypeName()

		for _, identifier := range udCtx.AllIdentifier() {
			userDefined.Add(identifier.GetText())
		}

		tn.SubType = TypeNameUserDefined
		tn.UserDefined = userDefined

	case ctx.Mapping() != nil:
		mapping := NewMapping()
		mapping.Visit(ctx.Mapping().(*parser.MappingContext))

		tn.SubType = TypeNameMapping
		tn.Mapping = mapping

	case ctx.TypeName() != nil && ctx.Expression() != nil:
		tn2 := NewTypeName()
		tn2.Visit(ctx.TypeName().(*parser.TypeNameContext))

		expr := NewExpression()
		expr.Visit(ctx.Expression().(*parser.ExpressionContext))

		tn.SubType = TypeNameArray
		tn.Expression = expr
		tn.TypeName = tn2

	case ctx.FunctionTypeName() != nil:
		ftn := NewFunctionTypeName()
		ftn.Visit(ctx.FunctionTypeName().(*parser.FunctionTypeNameContext))

		tn.Function = ftn

	default:
		panic("Unknown type name")
	}
}

func (tn *TypeName) String() string {
	switch tn.SubType {
	case TypeNameElementary:
		return tn.Elementary.String()
	case TypeNameUserDefined:
		return tn.UserDefined.String()
	case TypeNameMapping:
		return tn.Mapping.String()
	case TypeNameArray:
		if tn.Expression != nil {
			return fmt.Sprintf("%s[%s]", tn.TypeName, tn.Expression)
		}
		return tn.TypeName.String()
	case TypeNameFunction:
		tn.Function.String()
	}

	return "<unknown>"
}

// IsComplex returns true if type is of an array, mapping, or struct.
func (tn *TypeName) IsComplex() bool {
	switch tn.SubType {
	case TypeNameUserDefined, TypeNameMapping, TypeNameArray:
		return true
	}
	return false
}

// Equal evaluates the equality of two type names.
func (tn *TypeName) Equal(b *TypeName) bool {
	if tn.SubType != b.SubType {
		return false
	}

	if tn.SubType == TypeNameElementary && b.SubType == TypeNameElementary {
		return tn.Elementary.Equal(b.Elementary)
	}

	if tn.SubType == TypeNameUserDefined && b.SubType == TypeNameUserDefined {
		return tn.UserDefined.String() == b.UserDefined.String()
	}

	return false
}
