package sources

import "github.com/ericr/solanalyzer/parser"

// Contract represents a contract in Solidity.
type Contract struct {
	Tokens
	DefType     string
	Name        string
	Inheritance []*Inheritance
	StateVars   []*StateVar
	Functions   []*Function
}

// NewContract returns a new instance of Contract.
func NewContract() *Contract {
	return &Contract{}
}

// Visit is called by a visitor. See source.go for additional information on
// this pattern.
func (c *Contract) Visit(ctx *parser.ContractDefinitionContext) {
	c.Start = ctx.GetStart()
	c.Stop = ctx.GetStop()
	c.DefType = getText(ctx.GetChild(0))
	c.Name = ctx.Identifier().GetText()

	for _, inheritanceSpec := range ctx.AllInheritanceSpecifier() {
		inheritance := NewInheritance()
		inheritance.Visit(inheritanceSpec.(*parser.InheritanceSpecifierContext))

		c.Inheritance = append(c.Inheritance, inheritance)
	}

	for _, contractPart := range ctx.AllContractPart() {
		part := contractPart.(*parser.ContractPartContext)

		if part.StateVariableDeclaration() != nil {
			stateVar := NewStateVar()
			stateVar.Visit(part.StateVariableDeclaration().(*parser.StateVariableDeclarationContext))

			c.StateVars = append(c.StateVars, stateVar)
		}

		if part.FunctionDefinition() != nil {
			function := NewFunction()
			function.Visit(part.FunctionDefinition().(*parser.FunctionDefinitionContext))

			c.Functions = append(c.Functions, function)
		}
	}
}

func (c *Contract) String() string {
	return c.Name
}
