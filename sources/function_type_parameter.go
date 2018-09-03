package sources

import (
	"fmt"
	"github.com/ericr/solanalyzer/parser"
)

// FunctionTypeParameter represents a function type parameter in Solidity.
type FunctionTypeParameter struct {
	Tokens
	TypeName        *TypeName
	StorageLocation string
}

// NewFunctionTypeParameter returns a new instance of FunctionTypeParameter.
func (s *Source) NewFunctionTypeParameter() *FunctionTypeParameter {
	fnTypeParam := &FunctionTypeParameter{}
	s.AddNode(fnTypeParam)

	return fnTypeParam
}

// Visit is called by a visitor.
func (ftp *FunctionTypeParameter) Visit(s *Source, ctx *parser.FunctionTypeParameterContext) {
	ftp.Start = ctx.GetStart()
	ftp.Stop = ctx.GetStop()

	typeName := s.NewTypeName()
	typeName.Visit(s, ctx.TypeName().(*parser.TypeNameContext))

	ftp.TypeName = typeName

	if ctx.StorageLocation() != nil {
		ftp.StorageLocation = ctx.StorageLocation().GetText()
	}
}

func (ftp *FunctionTypeParameter) String() string {
	if ftp.StorageLocation == "" {
		return ftp.TypeName.String()
	}
	return fmt.Sprintf("%s %s", ftp.TypeName, ftp.StorageLocation)
}
