package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

// Inheritance represents inheritance in Solidity.
type Inheritance struct {
	Tokens
	TypeName    *UserDefinedTypeName
	Expressions []*Expression
}

// NewInheritance returns a new instance of Inheritance.
func NewInheritance() *Inheritance {
	return &Inheritance{
		Expressions: []*Expression{},
	}
}

// Visit is called by a visitor.
func (i *Inheritance) Visit(ctx *parser.InheritanceSpecifierContext) {
	i.Start = ctx.GetStart()
	i.Stop = ctx.GetStop()

	// udtn := NewUserDefinedTypeName()
	// udtn.Visit(ctx.UserDefinedTypeName().(*parser.UserDefinedTypeNameContext))

	for _, exprCtx := range ctx.AllExpression() {
		expr := NewExpression()
		expr.Visit(exprCtx.(*parser.ExpressionContext))

		i.Expressions = append(i.Expressions, expr)
	}
}

func (i *Inheritance) String() string {
	exprs := []string{}

	for _, expr := range i.Expressions {
		exprs = append(exprs, expr.String())
	}

	return fmt.Sprintf("%s(%s)", i.TypeName, strings.Join(exprs, ","))
}
