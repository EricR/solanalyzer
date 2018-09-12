package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalPrimary(expr *sources.PrimaryExpression) []*Value {
	defer e.Recover(expr.Tokens)

	var vals []*Value

	switch expr.SubType {
	case sources.ExpressionPrimaryIdentifier:
		variable := e.MustFindVariable(expr.Identifier)
		vals = append(vals, variable.Value)

	case sources.ExpressionPrimaryTuple:
		for _, subExpr := range expr.Tuple.Expressions {
			vals = append(vals, e.Eval(subExpr)...)
		}

	default:
		vals = []*Value{NewValue(expr)}
	}

	return vals
}
