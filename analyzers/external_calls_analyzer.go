package analyzers

import (
	"github.com/ericr/solanalyzer/emulator"
	"github.com/ericr/solanalyzer/sources"
)

// ExternalCallsAnalyzer is an analyzer that reports issues related to external
// function calls.
type ExternalCallsAnalyzer struct{}

// Name returns the name of the analyzer.
func (eca *ExternalCallsAnalyzer) Name() string {
	return "external calls"
}

// ID returns the unique ID of the analyzer.
func (eca *ExternalCallsAnalyzer) ID() string {
	return "external-calls"
}

// Execute runs the analyzer on a given source.
func (eca *ExternalCallsAnalyzer) Execute(source *sources.Source) ([]*Issue, error) {
	issues := []*Issue{}

	emulator := emulator.New(source)
	emulator.Run()

	return issues, nil
}
