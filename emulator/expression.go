package emulator

import "github.com/ericr/solanalyzer/sources"

// Eval evaluates an expression, returning a new expression.
func (e *Emulator) Eval(expr *sources.Expression) []*Value {
	defer e.Recover(expr.Tokens)

	switch expr.SubType {
	case sources.ExpressionPrimary:
		return e.evalPrimary(expr.Primary)

	case sources.ExpressionParentheses:
		return e.Eval(expr.SubExpression)

	case sources.ExpressionBinaryOperation:
		return e.evalBinaryOperation(expr)

	case sources.ExpressionFunctionCall:
		return e.evalFunctionCall(expr.SubExpression, expr.FunctionCallArgs)
	}

	value := NewValue()
	value.Expression.SubType = sources.ExpressionUnknown

	return []*Value{value}
}
