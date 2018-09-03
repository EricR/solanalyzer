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
func (s *Source) NewInheritance() *Inheritance {
	in := &Inheritance{
		Expressions: []*Expression{},
	}
	s.AddNode(in)

	return in
}

// Visit is called by a visitor.
func (i *Inheritance) Visit(s *Source, ctx *parser.InheritanceSpecifierContext) {
	i.Start = ctx.GetStart()
	i.Stop = ctx.GetStop()

	// udtn := s.NewUserDefinedTypeName()
	// udtn.Visit(s, ctx.UserDefinedTypeName().(*parser.UserDefinedTypeNameContext))

	for _, exprCtx := range ctx.AllExpression() {
		expr := s.NewExpression()
		expr.Visit(s, exprCtx.(*parser.ExpressionContext))

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
