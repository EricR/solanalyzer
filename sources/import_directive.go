package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
	"strings"
)

const (
	// ImportPath imports all modules from path
	ImportPath int = iota
	// ImprotModule imports a module from path
	ImprotModule
	// ImportModules imports modules from path
	ImportModules
)

// ImportDirective represents an import directive in Solidity.
type ImportDirective struct {
	Tokens
	SubType      int
	Module       string
	From         string
	As           string
	Declarations []*ImportDeclaration
}

// NewImportDirective returns a new instance of ImportDirective.
func (s *Source) NewImportDirective() *ImportDirective {
	dir := &ImportDirective{}
	s.AddNode(dir)

	return dir
}

// Visit is called by a visitor.
func (id *ImportDirective) Visit(s *Source, ctx *parser.ImportDirectiveContext) {
	id.Start = ctx.GetStart()
	id.Stop = ctx.GetStop()

	switch getText(ctx.GetChild(1))[0] {
	case '"':
		from := ctx.StringLiteral().GetText()

		id.SubType = ImportPath
		id.From = from[1 : len(from)-1]
		id.As = ctx.Identifier(0).GetText()

	case '{':
		for _, decCtx := range ctx.AllImportDeclaration() {
			from := ctx.StringLiteral().GetText()

			dec := s.NewImportDeclaration()
			dec.Visit(s, decCtx.(*parser.ImportDeclarationContext))

			id.SubType = ImprotModule
			id.Declarations = append(id.Declarations, dec)
			id.From = from[1 : len(from)-1]
		}

	default:
		// For cases that match 'import' ('*' | identifier) ...
		from := ctx.StringLiteral().GetText()

		id.SubType = ImportModules
		id.From = from[1 : len(from)-1]
		id.Module = getText(ctx.GetChild(0))

		if ctx.Identifier(1) != nil {
			id.As = ctx.Identifier(1).GetText()
		}
	}
}

func (id *ImportDirective) String() string {
	switch id.SubType {
	case ImportPath:
		if id.As == "" {
			return fmt.Sprintf("import \"%s\"", id.From)
		}
		return fmt.Sprintf("import \"%s\" as %s", id.From, id.As)
	case ImprotModule:
		if id.As == "" {
			return fmt.Sprintf("import %s from \"%s\"", id.Module, id.From)
		}
		return fmt.Sprintf("import %s as %s from \"%s\"", id.Module, id.As, id.From)
	case ImportModules:
		decs := []string{}

		for _, dec := range id.Declarations {
			decs = append(decs, dec.String())
		}

		return fmt.Sprintf("import { %s } from \"%s\"",
			strings.Join(decs, ","), id.From)
	default:
		panic("Unknown import directive sub-type")
	}
}
