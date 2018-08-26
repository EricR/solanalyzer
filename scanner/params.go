package scanner

import "github.com/ericr/solanalyzer/parser"

// EnterReturnParameters is called when the parser enters return parameters.
func (s *Scanner) EnterReturnParameters(ctx *parser.ReturnParametersContext) {
	s.inReturnParams = true
}

// ExitReturnParameters is called when the parser exits return parameters.
func (s *Scanner) ExitReturnParameters(ctx *parser.ReturnParametersContext) {
	s.inReturnParams = false
}

// ExitParameter is called when the parser exits a parameter.
func (s *Scanner) ExitParameter(ctx *parser.ParameterContext) {
	typeName := ctx.TypeName().GetText()

	if s.inReturnParams {
		s.function.Returns = append(s.function.Returns, typeName)
	} else {
		s.function.Params = append(s.function.Params, typeName)
	}
}
