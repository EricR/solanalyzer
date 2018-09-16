package emulator

import (
	"fmt"
	"github.com/ericr/solanalyzer/sources"
	"github.com/ericr/solanalyzer/utils"
)

// Value represents a value.
type Value struct {
	Solved     bool
	Expression *sources.Expression
}

// NewValue returns a new instance of Value.
func NewValue() *Value {
	return &Value{
		Solved:     false,
		Expression: sources.NewExpression(),
	}
}

// InferredTypeName returns the inferred type name of a value.
func (val Value) InferredTypeName() *sources.TypeName {
	switch val.Expression.SubType {
	case sources.ExpressionPrimary:
		switch val.Expression.Primary.SubType {
		case sources.ExpressionPrimaryBoolean:
			elementaryTypeName := sources.NewElementaryTypeName()
			elementaryTypeName.SubType = sources.ElementaryTypeNameBool

			typeName := sources.NewTypeName()
			typeName.SubType = sources.TypeNameElementary
			typeName.Elementary = elementaryTypeName

			return typeName

		case sources.ExpressionPrimaryNumber:
			typeName := sources.NewTypeName()
			typeName.SubType = sources.TypeNameElementary
			typeName.Elementary = inferIntegerType(val.Expression.Primary)

			return typeName

		case sources.ExpressionPrimaryHex:
			elementaryTypeName := sources.NewElementaryTypeName()
			elementaryTypeName.SubType = sources.ElementaryTypeNameString

			typeName := sources.NewTypeName()
			typeName.SubType = sources.TypeNameElementary
			typeName.Elementary = elementaryTypeName

			return typeName
		}
	}

	typeName := sources.NewTypeName()
	typeName.SubType = sources.TypeNameUnknown

	return typeName
}

func (val *Value) String() string {
	return val.Expression.String()
}

func inferIntegerType(expr *sources.PrimaryExpression) *sources.ElementaryTypeName {
	typeName := &sources.ElementaryTypeName{}
	integerStr := expr.Integer.String()
	negative := false

	// Hex numbers that are valid addresses are treated as such
	if expr.Hex != "" && utils.FullyValidAddress(expr.Hex) {
		typeName.SubType = sources.ElementaryTypeNameAddress

		return typeName
	}

	if integerStr[0] == '-' {
		negative = true
		integerStr = integerStr[1:len(integerStr)]
		typeName.SubType = sources.ElementaryTypeNameInt
	} else {
		typeName.SubType = sources.ElementaryTypeNameUint
	}

	typeName.Size = utils.SmallestIntSize(expr.Integer, negative)

	if negative {
		typeName.Text = fmt.Sprintf("int%d", typeName.Size)
	} else {
		typeName.Text = fmt.Sprintf("uint%d", typeName.Size)
	}

	return typeName
}
