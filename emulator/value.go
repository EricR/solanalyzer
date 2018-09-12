package emulator

import (
	"fmt"
	"github.com/ericr/solanalyzer/sources"
	"github.com/ericr/solanalyzer/utils"
)

// Value either represents a symbolic or solved primary expression.
type Value struct {
	Solved     bool
	Expression *sources.PrimaryExpression
}

// NewValue returns a new instance of Value.
func NewValue(expr *sources.PrimaryExpression) *Value {
	if expr == nil {
		return &Value{}
	}
	return &Value{
		Solved:     true,
		Expression: expr,
	}
}

// InferredType returns the inferred type of a value.
func (val *Value) InferredType() *sources.TypeName {
	if !val.Solved {
		panic("Cannot infer type of an unsolved value")
	}

	switch val.Expression.SubType {
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
		typeName.Elementary = inferIntegerType(val.Expression)

		return typeName

	case sources.ExpressionPrimaryHex:
		elementaryTypeName := sources.NewElementaryTypeName()
		elementaryTypeName.SubType = sources.ElementaryTypeNameString

		typeName := sources.NewTypeName()
		typeName.SubType = sources.TypeNameElementary
		typeName.Elementary = elementaryTypeName

		return typeName

	default:
		panic("Could not infer type")
	}

	return nil
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
