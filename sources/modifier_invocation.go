package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// ModifierInvocation represents a modifier invocation in Solidity.
type ModifierInvocation struct {
	Tokens
	Identifier  string
	Expressions []*Expression
}

// NewModifierInvocation returns a new instance of ModifierInvocation.
func NewModifierInvocation() *ModifierInvocation {
	return &ModifierInvocation{
		Expressions: []*Expression{},
	}
}

// Visit is called by a visitor.
func (mi *ModifierInvocation) Visit(ctx *parser.ModifierInvocationContext) {
	mi.Start = ctx.GetStart()
	mi.Stop = ctx.GetStop()
	mi.Identifier = ctx.Identifier().GetText()

	if ctx.ExpressionList() != nil {
		expList := ctx.ExpressionList().(*parser.ExpressionListContext)

		for _, exprCtx := range expList.AllExpression() {
			expression := NewExpression()
			expression.Visit(exprCtx.(*parser.ExpressionContext))

			mi.Expressions = append(mi.Expressions, expression)
		}
	}
}

func (mi *ModifierInvocation) String() string {
	exprs := []string{}

	for _, expr := range mi.Expressions {
		exprs = append(exprs, expr.String())
	}

	return fmt.Sprintf("%s(%s)", mi.Identifier, strings.Join(exprs, ", "))
}
