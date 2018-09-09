package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// TupleExpression represents a tuple expression in Solidity.
type TupleExpression struct {
	Tokens
	SquareBrackets bool
	Expressions    []*Expression
}

// NewTupleExpression returns a new instance of TupleExpression.
func NewTupleExpression() *TupleExpression {
	return &TupleExpression{}
}

// Visit is called by a visitor.
func (te *TupleExpression) Visit(ctx *parser.TupleExpressionContext) {
	te.Start = ctx.GetStart()
	te.Stop = ctx.GetStop()

	if getText(ctx.GetChild(0)) == "[" {
		te.SquareBrackets = true
	}

	for _, exprCtx := range ctx.AllExpression() {
		expr := NewExpression()
		expr.Visit(exprCtx.(*parser.ExpressionContext))

		te.Expressions = append(te.Expressions, expr)
	}
}

func (te *TupleExpression) String() string {
	exprs := []string{}

	for _, expr := range te.Expressions {
		exprs = append(exprs, expr.String())
	}

	if te.SquareBrackets {
		return fmt.Sprintf("[%s]", strings.Join(exprs, ", "))
	}
	return fmt.Sprintf("(%s)", strings.Join(exprs, ", "))
}
