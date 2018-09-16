package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalBinaryOperation(expr *sources.Expression) []*Value {
	left := e.Eval(expr.LeftExpression)
	right := e.Eval(expr.RightExpression)

	if len(left) != 1 {
		panic("Wrong number of operands on left side for binary operation")
	}

	if len(right) != 1 {
		panic("Wrong number of operands on right side for binary operation")
	}

	leftExpr := left[0].Expression
	rightExpr := right[0].Expression

	value := NewValue()
	value.Expression = expr
	value.Expression.LeftExpression = leftExpr
	value.Expression.RightExpression = rightExpr

	if leftExpr.SubType == sources.ExpressionPrimary &&
		leftExpr.Primary.SubType == sources.ExpressionPrimaryNumber &&
		rightExpr.SubType == sources.ExpressionPrimary &&
		rightExpr.Primary.SubType == sources.ExpressionPrimaryNumber {

		value = evalMath(expr.Operation, left[0].Expression.Primary, right[0].Expression.Primary)

		switch expr.Operation {
		case "=", "+=", "-=", "*=", "/=", "%=":
			e.MustSetVariable(expr.LeftExpression.String(), value)
		}
	}

	return []*Value{value}
}
