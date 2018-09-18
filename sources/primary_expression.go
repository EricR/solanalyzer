package sources

import (
	"github.com/ericr/solanalyzer/parser"
	"github.com/ericr/solanalyzer/utils"
	"math/big"
	"regexp"
)

var supportedNumberLiteral = regexp.MustCompile(`^-?[0-9]+$`)

const (
	// ExpressionPrimaryUnknown represents an unknown expression.
	ExpressionPrimaryUnknown = iota
	// ExpressionPrimaryBoolean represents a boolean expression.
	ExpressionPrimaryBoolean
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
	Boolean            bool
	Integer            *big.Int
	Hex                string
	Text               string
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
		pe.Boolean = ctx.BooleanLiteral().GetText() == "true"

	case ctx.NumberLiteral() != nil:
		numLiteral := ctx.NumberLiteral().(*parser.NumberLiteralContext)

		switch {
		case numLiteral.DecimalNumber() != nil:
			decimal := getText(numLiteral.DecimalNumber())
			checkNumberLiteral(decimal)
			pe.Integer = utils.MustParseBigInt(decimal)
			pe.SubType = ExpressionPrimaryNumber

		case numLiteral.HexNumber() != nil:
			pe.Hex = getText(numLiteral.HexNumber())
			pe.SubType = ExpressionPrimaryHex
		}

		if numLiteral.NumberUnit() != nil {
			units := getText(numLiteral.NumberUnit())
			pe.Integer = utils.IntOfUnit(pe.Integer, units)
			pe.SubType = ExpressionPrimaryNumber
		}

	case ctx.HexLiteral() != nil:
		pe.SubType = ExpressionPrimaryHex
		pe.Text = ctx.HexLiteral().GetText()

	case ctx.StringLiteral() != nil:
		pe.SubType = ExpressionPrimaryString
		pe.Text = ctx.StringLiteral().GetText()

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
	switch pe.SubType {
	case ExpressionPrimaryBoolean:
		if pe.Boolean {
			return "true"
		}
		return "false"

	case ExpressionPrimaryNumber:
		return pe.Integer.String()

	case ExpressionPrimaryHex, ExpressionPrimaryString:
		return pe.Text

	case ExpressionPrimaryIdentifier:
		return pe.Identifier

	case ExpressionPrimaryTuple:
		return pe.Tuple.String()

	case ExpressionPrimaryElementaryTypeName:
		return pe.ElementaryTypeName.String()

	default:
		return "<UNKNOWN>"
	}
}

func checkNumberLiteral(text string) string {
	if !supportedNumberLiteral.MatchString(text) {
		panic("Non-integers in integer literals are not supported")
	}

	return text
}
