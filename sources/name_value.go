package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// NameValue represents a name-value in Solidity.
type NameValue struct {
	Tokens
	Identifier string
	Expression *Expression
}

// NewNameValue returns a new instance of NameValue.
func NewNameValue() *NameValue {
	return &NameValue{}
}

// Visit is called by a visitor.
func (nv *NameValue) Visit(ctx *parser.NameValueContext) {
	nv.Start = ctx.GetStart()
	nv.Stop = ctx.GetStop()

	expr := NewExpression()
	expr.Visit(ctx.Expression().(*parser.ExpressionContext))

	nv.Identifier = ctx.Identifier().GetText()
	nv.Expression = expr
}

func (nv *NameValue) String() string {
	return fmt.Sprintf("%s:%s", nv.Identifier, nv.Expression)
}
