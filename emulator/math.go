package emulator

import (
	"github.com/ericr/solanalyzer/sources"
	"github.com/ericr/solanalyzer/utils"
	"math/big"
)

func evalMath(op string, expr1 *sources.PrimaryExpression, expr2 *sources.PrimaryExpression) *Value {
	sExpr := sources.NewPrimaryExpression()
	a := utils.MustParseBigInt(expr1.String())
	b := utils.MustParseBigInt(expr2.String())
	c := big.NewInt(0)

	switch op {
	case "=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = b

	case "+", "+=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Add(a, b)

	case "-", "-=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Sub(a, b)

	case "*", "*=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Mul(a, b)

	case "/", "/=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Div(a, b)

	case "%", "%=":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Mod(a, b)

	case "**":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Exp(a, b, c)

	case ">>":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Rsh(a, utils.MustParseUint(expr2.String()))

	case "<<":
		sExpr.SubType = sources.ExpressionPrimaryNumber
		sExpr.Integer = c.Lsh(a, utils.MustParseUint(expr2.String()))

	case "==":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) == 0

	case "!=":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) != 0

	case "<":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) == -1

	case "<=":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) <= 0

	case ">":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) == 1

	case ">=":
		sExpr.SubType = sources.ExpressionPrimaryBoolean
		sExpr.Boolean = a.Cmp(b) >= 0
	}

	value := NewValue()
	value.Expression.SubType = sources.ExpressionPrimary
	value.Expression.Primary = sExpr
	value.Solved = true

	return value
}
