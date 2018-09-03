package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

const (
	// ExpressionPrimary represents a primary expression.
	ExpressionPrimary = iota
	// ExpressionNew represents a 'new' expression.
	ExpressionNew
	// ExpressionUnaryOperation represents a unary operation expression.
	ExpressionUnaryOperation
	// ExpressionParentheses represents a sub-expression in parentheses.
	ExpressionParentheses
	// ExpressionMemberAccess represents a member access expression.
	ExpressionMemberAccess
	// ExpressionBinaryOperation represents a binary operation expression.
	ExpressionBinaryOperation
	// ExpressionFunctionCall represents a function call expression.
	ExpressionFunctionCall
	// ExpressionIndexAccess represents an index access expression.
	ExpressionIndexAccess
	// ExpressionTernary represents a ternary expression.
	ExpressionTernary
)

// Expression represents an expression in Solidity.
type Expression struct {
	Tokens
	SubType          int
	Operation        string
	MemberName       string
	TypeName         *TypeName
	Primary          *PrimaryExpression
	FunctionCallArgs *FunctionCallArguments
	SubExpression    *Expression
	LeftExpression   *Expression
	RightExpression  *Expression
	IndexExpression  *Expression
	TernaryIf        *Expression
	TernaryThen      *Expression
	TernaryElse      *Expression
}

// NewExpression returns a new instance of Expression.
func (s *Source) NewExpression() *Expression {
	expr := &Expression{}
	s.AddNode(expr)

	return expr
}

// Visit is called by a visitor.
func (e *Expression) Visit(s *Source, ctx *parser.ExpressionContext) {
	e.Start = ctx.GetStart()
	e.Stop = ctx.GetStop()

	switch ctx.GetChildCount() {
	case 1:
		primaryExpr := s.NewPrimaryExpression()
		primaryExpr.Visit(s, ctx.PrimaryExpression().(*parser.PrimaryExpressionContext))

		e.SubType = ExpressionPrimary
		e.Primary = primaryExpr

	case 2:
		switch getText(ctx.GetChild(0)) {
		case "new":
			typeName := s.NewTypeName()
			typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

			e.SubType = ExpressionNew
			e.TypeName = typeName

		case "++", "--", "after", "delete", "!", "~":
			expr := s.NewExpression()
			expr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

			e.SubType = ExpressionUnaryOperation
			e.Operation = getText(ctx.GetChild(0))
			e.SubExpression = expr

		default:
			switch getText(ctx.GetChild(1)) {
			case "++", "--":
				expr := s.NewExpression()
				expr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

				e.SubType = ExpressionUnaryOperation
				e.Operation = getText(ctx.GetChild(1))
				e.SubExpression = expr

			default:
				panic("Unknown expression(2)")
			}
		}

	case 3:
		if getText(ctx.GetChild(0)) == "(" && getText(ctx.GetChild(2)) == ")" {
			expr := s.NewExpression()
			expr.Visit(s, ctx.Expression(1).(*parser.ExpressionContext))

			e.SubType = ExpressionParentheses
			e.SubExpression = expr

			return
		}

		switch getText(ctx.GetChild(1)) {
		case ".":
			expr := s.NewExpression()
			expr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

			e.SubType = ExpressionMemberAccess
			e.MemberName = getText(ctx.GetChild(2))
			e.SubExpression = expr

		case "**", "*", "/", "%", "+", "-", "<<", ">>", "&", "^", "|", "<", ">",
			"<=", ">=", "==", "!=", "&&", "||", "=", "|=", "^=", "&=", "<<=", ">>=",
			"+=", "-=", "*=", "/=", "%=":
			leftExpr := s.NewExpression()
			leftExpr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

			rightExpr := s.NewExpression()
			rightExpr.Visit(s, ctx.Expression(1).(*parser.ExpressionContext))

			e.SubType = ExpressionBinaryOperation
			e.Operation = getText(ctx.GetChild(1))
			e.LeftExpression = leftExpr
			e.RightExpression = rightExpr

		default:
			panic("Unknown expression(3)")
		}

	case 4:
		switch {
		case getText(ctx.GetChild(1)) == "(" && getText(ctx.GetChild(3)) == ")":
			expr := s.NewExpression()
			expr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

			funCallArgs := s.NewFunctionCallArguments()
			funCallArgs.Visit(s, ctx.FunctionCallArguments().(*parser.FunctionCallArgumentsContext))

			e.SubType = ExpressionFunctionCall
			e.SubExpression = expr
			e.FunctionCallArgs = funCallArgs

		case getText(ctx.GetChild(1)) == "[" && getText(ctx.GetChild(3)) == "]":
			subExpr := s.NewExpression()
			subExpr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

			if ctx.Expression(1) != nil {
				indexExpr := s.NewExpression()
				indexExpr.Visit(s, ctx.Expression(1).(*parser.ExpressionContext))

				e.IndexExpression = indexExpr
			}

			e.SubType = ExpressionIndexAccess
			e.SubExpression = subExpr

		default:
			panic("Unknown expression(4)")
		}
	case 5:
		ifExpr := s.NewExpression()
		ifExpr.Visit(s, ctx.Expression(0).(*parser.ExpressionContext))

		thenExpr := s.NewExpression()
		thenExpr.Visit(s, ctx.Expression(1).(*parser.ExpressionContext))

		elseExpr := s.NewExpression()
		elseExpr.Visit(s, ctx.Expression(2).(*parser.ExpressionContext))

		e.SubType = ExpressionTernary
		e.TernaryIf = ifExpr
		e.TernaryThen = thenExpr
		e.TernaryElse = elseExpr
	}
}

func (e *Expression) String() string {
	switch e.SubType {
	case ExpressionPrimary:
		return e.Primary.String()

	case ExpressionNew:
		return fmt.Sprintf("new %s", e.TypeName)

	case ExpressionUnaryOperation:
		return fmt.Sprintf("%s %s", e.Operation, e.SubExpression)

	case ExpressionParentheses:
		return fmt.Sprintf("(%s)", e.SubExpression)

	case ExpressionMemberAccess:
		return fmt.Sprintf("%s.%s", e.SubExpression, e.MemberName)

	case ExpressionBinaryOperation:
		return fmt.Sprintf("%s %s %s", e.LeftExpression, e.Operation, e.RightExpression)

	case ExpressionFunctionCall:
		return fmt.Sprintf("%s(%s)", e.SubExpression, e.FunctionCallArgs)

	case ExpressionIndexAccess:
		return fmt.Sprintf("%s[%s]", e.SubExpression, e.IndexExpression)

	case ExpressionTernary:
		return fmt.Sprintf("%s ? %s : %s", e.TernaryIf, e.TernaryThen, e.TernaryElse)

	default:
		panic("unknown expression sub-type")
	}

	return "(unknown expression)"
}
