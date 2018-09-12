package sources

import "github.com/ericr/solanalyzer/parser"

// Contract represents a contract in Solidity.
type Contract struct {
	Tokens
	Source      *Source
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
func NewContract(source *Source) *Contract {
	return &Contract{
		Source:      source,
		Inheritance: []*Inheritance{},
		StateVars:   []*StateVariable{},
		UsingFor:    []*UsingFor{},
		Structs:     []*Struct{},
		Modifiers:   []*Modifier{},
		Functions:   []*Function{},
	}
}

// Visit is called by a visitor.
func (c *Contract) Visit(ctx *parser.ContractDefinitionContext) {
	c.Start = ctx.GetStart()
	c.Stop = ctx.GetStop()
	c.DefType = getText(ctx.GetChild(0))
	c.Identifier = ctx.Identifier().GetText()

	for _, inheritanceSpec := range ctx.AllInheritanceSpecifier() {
		inheritance := NewInheritance()
		inheritance.Visit(inheritanceSpec.(*parser.InheritanceSpecifierContext))

		c.Inheritance = append(c.Inheritance, inheritance)
	}

	for _, contractPart := range ctx.AllContractPart() {
		part := contractPart.(*parser.ContractPartContext)

		switch {
		case part.StateVariableDeclaration() != nil:
			stateVar := NewStateVariable()
			stateVar.Visit(part.StateVariableDeclaration().(*parser.StateVariableDeclarationContext))

			c.StateVars = append(c.StateVars, stateVar)

		case part.UsingForDeclaration() != nil:
			usingFor := NewUsingFor()
			usingFor.Visit(part.UsingForDeclaration().(*parser.UsingForDeclarationContext))

			c.UsingFor = append(c.UsingFor, usingFor)

		case part.StructDefinition() != nil:
			structDef := NewStruct()
			structDef.Visit(part.StructDefinition().(*parser.StructDefinitionContext))

			c.Structs = append(c.Structs, structDef)

		case part.ConstructorDefinition() != nil:
			constructor := NewConstructor()
			constructor.Visit(part.ConstructorDefinition().(*parser.ConstructorDefinitionContext))

			c.Constructor = constructor

		case part.FunctionDefinition() != nil:
			function := NewFunction()
			function.Visit(part.FunctionDefinition().(*parser.FunctionDefinitionContext))

			c.Functions = append(c.Functions, function)

		case part.EventDefinition() != nil:
			event := NewEvent()
			event.Visit(part.EventDefinition().(*parser.EventDefinitionContext))

			c.Events = append(c.Events, event)

		case part.EnumDefinition() != nil:
			enum := NewEnum()
			enum.Visit(part.EnumDefinition().(*parser.EnumDefinitionContext))

			c.Enums = append(c.Enums, enum)

		default:
			panic("Unknown type of contract part")
		}
	}
}

func (c *Contract) String() string {
	return c.Identifier
}
