package sources

import "github.com/ericr/solanalyzer/parser"

const (
	// ExpressionPrimaryBoolean represents a boolean expression.
	ExpressionPrimaryBoolean = iota
	// ExpressionPrimaryNumber represents a number expression.
	ExpressionPrimaryNumber
	// ExpressionPrimaryHex represents a hexadecimal expression.
	ExpressionPrimaryHex
	// ExpressionPrimaryString represents a string expression.
	ExpressionPrimaryString
	// ExpressionPrimaryIdentifier represents an identifier expression.
	ExpressionPrimaryIdentifier
	// ExpressionPrimaryTuple represents an tuple expression.
	ExpressionPrimaryTuple
	// ExpressionPrimaryElementaryTypeName represents an elementary type name
	// expression.
	ExpressionPrimaryElementaryTypeName
)

// PrimaryExpression represents a primary expression in Solidity.
type PrimaryExpression struct {
	Tokens
	SubType            int
	Boolean            string
	Number             string
	Hex                string
	StringLit          string
	Identifier         string
	Tuple              *TupleExpression
	ElementaryTypeName *ElementaryTypeName
}

// NewPrimaryExpression returns a new instance of PrimaryExpression.
func NewPrimaryExpression() *PrimaryExpression {
	return &PrimaryExpression{}
}

// Visit is called by a visitor.
func (pe *PrimaryExpression) Visit(ctx *parser.PrimaryExpressionContext) {
	pe.Start = ctx.GetStart()
	pe.Stop = ctx.GetStop()

	switch {
	case ctx.BooleanLiteral() != nil:
		pe.SubType = ExpressionPrimaryBoolean
		pe.Boolean = ctx.BooleanLiteral().GetText()

	case ctx.NumberLiteral() != nil:
		pe.SubType = ExpressionPrimaryNumber
		pe.Number = ctx.NumberLiteral().GetText()

	case ctx.HexLiteral() != nil:
		pe.SubType = ExpressionPrimaryHex
		pe.Number = ctx.HexLiteral().GetText()

	case ctx.StringLiteral() != nil:
		pe.SubType = ExpressionPrimaryString
		pe.StringLit = ctx.StringLiteral().GetText()

	case ctx.Identifier() != nil:
		pe.SubType = ExpressionPrimaryIdentifier
		pe.Identifier = ctx.Identifier().GetText()

	case ctx.TupleExpression() != nil:
		tuple := NewTupleExpression()
		tuple.Visit(ctx.TupleExpression().(*parser.TupleExpressionContext))

		pe.SubType = ExpressionPrimaryTuple
		pe.Tuple = tuple

	case ctx.ElementaryTypeNameExpression() != nil:
		etneCtx := ctx.ElementaryTypeNameExpression().(*parser.ElementaryTypeNameExpressionContext)
		etnCtx := etneCtx.ElementaryTypeName().(*parser.ElementaryTypeNameContext)

		elementaryTypeName := NewElementaryTypeName()
		elementaryTypeName.Visit(etnCtx)

		pe.SubType = ExpressionPrimaryElementaryTypeName
		pe.ElementaryTypeName = elementaryTypeName

	default:
		panic("Unknown type of primary expression")
	}
}

func (pe *PrimaryExpression) String() string {
	return "TODO"
}
