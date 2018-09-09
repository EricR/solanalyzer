package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// StateVariable represents a state variable in Solidity.
type StateVariable struct {
	Tokens
	Contract   *Contract
	TypeName   *TypeName
	Public     bool
	Internal   bool
	Private    bool
	Constant   bool
	Identifier string
	Expression *Expression
}

// NewStateVariable returns a new instance of StateVariable.
func NewStateVariable() *StateVariable {
	return &StateVariable{}
}

// Visit is called by a visitor.
func (sv *StateVariable) Visit(ctx *parser.StateVariableDeclarationContext) {
	sv.Start = ctx.GetStart()
	sv.Stop = ctx.GetStop()

	typeName := NewTypeName()
	typeName.Visit(ctx.TypeName().(*parser.TypeNameContext))

	sv.TypeName = typeName
	sv.Public = ctx.PublicKeyword(0) != nil
	sv.Internal = ctx.InternalKeyword(0) != nil
	sv.Private = ctx.PrivateKeyword(0) != nil
	sv.Constant = ctx.ConstantKeyword(0) != nil

	sv.Identifier = ctx.Identifier().GetText()

	if ctx.Expression() != nil {
		expr := NewExpression()
		expr.Visit(ctx.Expression().(*parser.ExpressionContext))
	}
}

func (sv *StateVariable) String() string {
	str := sv.TypeName.String()

	if sv.Public {
		str += " public"
	}
	if sv.Internal {
		str += " internal"
	}
	if sv.Private {
		str += " private"
	}
	if sv.Constant {
		str += " constant"
	}

	if sv.Expression == nil {
		return fmt.Sprintf("%s %s", str, sv.Identifier)
	}
	return fmt.Sprintf("%s %s = %s", str, sv.Identifier, sv.Expression)
}
