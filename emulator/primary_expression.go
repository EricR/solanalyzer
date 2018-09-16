package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalPrimary(pExpr *sources.PrimaryExpression) []*Value {
	defer e.Recover(pExpr.Tokens)

	var values []*Value

	switch pExpr.SubType {
	case sources.ExpressionPrimaryIdentifier:
		values = append(values, e.MustResolveIdentifier(pExpr.Identifier))

	case sources.ExpressionPrimaryTuple:
		for _, sExpr := range pExpr.Tuple.Expressions {
			values = append(values, e.Eval(sExpr)...)
		}

	default:
		value := NewValue()
		value.Expression.SubType = sources.ExpressionPrimary
		value.Expression.Primary = pExpr

		values = []*Value{value}
	}

	return values
}
