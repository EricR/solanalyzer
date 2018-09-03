package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// Contract represents a contract in Solidity.
type Contract struct {
	Tokens
	DefType     string
	Identifier  string
	Inheritance []*Inheritance
	StateVars   []*StateVariable
	UsingFor    []*UsingFor
	Structs     []*Struct
	Modifiers   []*Modifier
	Constructor *Constructor
	Functions   []*Function
	Events      []*Event
	Enums       []*Enum
}

// NewContract returns a new instance of Contract.
func (s *Source) NewContract() *Contract {
	contract := &Contract{
		Inheritance: []*Inheritance{},
		StateVars:   []*StateVariable{},
		UsingFor:    []*UsingFor{},
		Structs:     []*Struct{},
		Modifiers:   []*Modifier{},
		Functions:   []*Function{},
	}
	s.AddNode(contract)

	return contract
}

// Visit is called by a visitor.
func (c *Contract) Visit(s *Source, ctx *parser.ContractDefinitionContext) {
	c.Start = ctx.GetStart()
	c.Stop = ctx.GetStop()
	c.DefType = getText(ctx.GetChild(0))
	c.Identifier = ctx.Identifier().GetText()

	for _, inheritanceSpec := range ctx.AllInheritanceSpecifier() {
		inheritance := s.NewInheritance()
		inheritance.Visit(s, inheritanceSpec.(*parser.InheritanceSpecifierContext))

		c.Inheritance = append(c.Inheritance, inheritance)
	}

	for _, contractPart := range ctx.AllContractPart() {
		part := contractPart.(*parser.ContractPartContext)

		switch {
		case part.StateVariableDeclaration() != nil:
			stateVar := s.NewStateVariable()
			stateVar.Visit(s, part.StateVariableDeclaration().(*parser.StateVariableDeclarationContext))

			c.StateVars = append(c.StateVars, stateVar)

		case part.UsingForDeclaration() != nil:
			usingFor := s.NewUsingFor()
			usingFor.Visit(s, part.UsingForDeclaration().(*parser.UsingForDeclarationContext))

			c.UsingFor = append(c.UsingFor, usingFor)

		case part.StructDefinition() != nil:
			structDef := s.NewStruct()
			structDef.Visit(s, part.StructDefinition().(*parser.StructDefinitionContext))

			c.Structs = append(c.Structs, structDef)

		case part.ConstructorDefinition() != nil:
			constructor := s.NewConstructor()
			constructor.Visit(s, part.ConstructorDefinition().(*parser.ConstructorDefinitionContext))

			c.Constructor = constructor

		case part.FunctionDefinition() != nil:
			function := s.NewFunction()
			function.Visit(s, part.FunctionDefinition().(*parser.FunctionDefinitionContext))

			c.Functions = append(c.Functions, function)

		case part.EventDefinition() != nil:
			event := s.NewEvent()
			event.Visit(s, part.EventDefinition().(*parser.EventDefinitionContext))

			c.Events = append(c.Events, event)

		case part.EnumDefinition() != nil:
			enum := s.NewEnum()
			enum.Visit(s, part.EnumDefinition().(*parser.EnumDefinitionContext))

			c.Enums = append(c.Enums, enum)

		default:
			panic("Unknown type of contract part")
		}
	}
}

func (c *Contract) String() string {
	return fmt.Sprintf("%s %s", c.DefType, c.Identifier)
}
