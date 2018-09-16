package emulator

import "github.com/ericr/solanalyzer/sources"

func (e *Emulator) evalFunctionCall(expr *sources.Expression, args *sources.FunctionCallArguments) []*Value {
	e.Emit("function_call", &Event{
		Source:         e.source,
		Contract:       e.contract,
		CallerFunction: e.function,
		Function:       e.mustResolveFunction(expr, args),
	})

	return []*Value{}
}

func (e *Emulator) mustResolveFunction(expr *sources.Expression, args *sources.FunctionCallArguments) *sources.Function {
	identifier := expr.Primary.Identifier

	if args.SubType == sources.FunctionCallArgsWithNameValues {
		for _, fn := range e.functions {
			if fn.Identifier == identifier && e.nvArgsEqual(args.NameValues, fn.Parameters) {
				return fn
			}
		}
	}

	for _, fn := range e.functions {
		if fn.Identifier == identifier && e.exprArgsEqual(args.Expressions, fn.Parameters) {
			return fn
		}
	}

	panic("Failed to resolve function")

	return nil
}

func (e *Emulator) nvArgsEqual(nameValues []*sources.NameValue, params []*sources.Parameter) bool {
	for i, nv := range nameValues {
		if params[i].Identifier != nv.Identifier {
			return false
		}
	}

	return true
}

func (e *Emulator) exprArgsEqual(exprs []*sources.Expression, params []*sources.Parameter) bool {
	if len(exprs) != len(params) {
		return false
	}
	
	for i, expr := range exprs {
		values := e.Eval(expr)

		if len(values) != 1 {
			panic("Parameter expression did not evaluate to one value")
		}

		inferredType := values[0].InferredTypeName()

		if !inferredType.Equal(params[i].TypeName) {
			return false
		}
	}

	return true
}
