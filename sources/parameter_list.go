package sources

import "strings"

// ParameterList represents a list of parameters in Solidity.
type ParameterList []*Parameter

// NewParameterList returns a new instance of ParameterList.
func NewParameterList() *ParameterList {
	return &ParameterList{}
}

// Add adds a new parameter to an instance of ParameterList.
func (pl *ParameterList) Add(param *Parameter) {
	*pl = append(*pl, param)
}

func (pl *ParameterList) String() string {
	paramStrs := []string{}

	for _, param := range *pl {
		paramStrs = append(paramStrs, param.String())
	}

	return strings.Join(paramStrs, ", ")
}
