package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalBinaryOperation(op string, expr1 *sources.Expression, expr2 *sources.Expression) []*Value {
	left := e.Eval(expr1)
	right := e.Eval(expr2)
	answer := &Value{}

	if len(left) != 1 {
		panic("Wrong number of operands on left side for binary operation")
	}

	if len(right) != 1 {
		panic("Wrong number of operands on right side for binary operation")
	}

	if left[0].Expression.SubType != sources.ExpressionPrimaryNumber {
		panic("Left side of binary operation is NaN")
	}

	if right[0].Expression.SubType != sources.ExpressionPrimaryNumber {
		panic("Right side of binary operation is NaN")
	}

	answer = evalMath(op, left[0].Expression, right[0].Expression)

	switch op {
	case "=", "+=", "-=", "*=", "/=", "%=":
		e.MustAssignVariable(expr1.String(), answer)
	}

	return []*Value{answer}
}
