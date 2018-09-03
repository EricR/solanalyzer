package analyzers

import (
	"reflect"
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
)

// CallGraphAnalyzer is an analyzer that generates a call graph.
type CallGraphAnalyzer struct{}

// Name returns the name of the analyzer.
func (cga *CallGraphAnalyzer) Name() string {
	return "Call Graph"
}

// ID cga the unique ID of the analyzer.
func (fva *CallGraphAnalyzer) ID() string {
	return "call-graph"
}

// Execute runs the analyzer on a given source.
func (cga *CallGraphAnalyzer) Execute(source *sources.Source) ([]*Issue, error) {
	for _, node := range source.Nodes {
		logrus.Infof("%s: %s", reflect.TypeOf(node), node)
	}

	return []*Issue{}, nil
}
