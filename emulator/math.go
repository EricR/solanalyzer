package emulator

import (
	"github.com/ericr/solanalyzer/sources"
	"github.com/ericr/solanalyzer/utils"
	"math/big"
)

func evalMath(op string, expr1 *sources.PrimaryExpression, expr2 *sources.PrimaryExpression) *Value {
	expr := sources.NewPrimaryExpression()
	a := utils.MustParseBigInt(expr1.String())
	b := utils.MustParseBigInt(expr2.String())
	c := big.NewInt(0)

	switch op {
	case "+", "+=":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Add(a, b)

	case "-", "-=":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Sub(a, b)

	case "*", "*=":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Mul(a, b)

	case "/", "/=":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Div(a, b)

	case "%", "%=":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Mod(a, b)

	case "**":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Exp(a, b, c)

	case ">>":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Rsh(a, utils.MustParseUint(expr2.String()))

	case "<<":
		expr.SubType = sources.ExpressionPrimaryNumber
		expr.Integer = c.Lsh(a, utils.MustParseUint(expr2.String()))

	case "==":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) == 0

	case "!=":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) != 0

	case "<":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) == -1

	case "<=":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) <= 0

	case ">":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) == 1

	case ">=":
		expr.SubType = sources.ExpressionPrimaryBoolean
		expr.Boolean = a.Cmp(b) >= 0
	}

	return NewValue(expr)
}
