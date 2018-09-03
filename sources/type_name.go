package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

const (
	// TypeNameElementary represents an elementary type.
	TypeNameElementary int = iota
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
func (s *Source) NewTypeName() *TypeName {
	typeName := &TypeName{}
	s.AddNode(typeName)

	return typeName
}

// Visit is called by a visitor.
func (tn *TypeName) Visit(s *Source, ctx *parser.TypeNameContext) {
	tn.Start = ctx.GetStart()
	tn.Stop = ctx.GetStop()

	switch {
	case ctx.ElementaryTypeName() != nil:
		etn := s.NewElementaryTypeName()
		etn.Visit(s, ctx.ElementaryTypeName().(*parser.ElementaryTypeNameContext))

		tn.SubType = TypeNameElementary
		tn.Elementary = etn

	case ctx.UserDefinedTypeName() != nil:
		udCtx := ctx.UserDefinedTypeName().(*parser.UserDefinedTypeNameContext)
		userDefined := s.NewUserDefinedTypeName()

		for _, identifier := range udCtx.AllIdentifier() {
			userDefined.Add(identifier.GetText())
		}

		tn.SubType = TypeNameUserDefined
		tn.UserDefined = userDefined

	case ctx.Mapping() != nil:
		mapping := s.NewMapping()
		mapping.Visit(s, ctx.Mapping().(*parser.MappingContext))

		tn.SubType = TypeNameMapping
		tn.Mapping = mapping

	case ctx.TypeName() != nil && ctx.Expression() != nil:
		tn2 := s.NewTypeName()
		tn2.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

		expr := s.NewExpression()
		expr.Visit(s, ctx.Expression().(*parser.ExpressionContext))

		tn.SubType = TypeNameArray
		tn.Expression = expr
		tn.TypeName = tn2

	case ctx.FunctionTypeName() != nil:
		ftn := s.NewFunctionTypeName()
		ftn.Visit(s, ctx.FunctionTypeName().(*parser.FunctionTypeNameContext))

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

	return ""
}
