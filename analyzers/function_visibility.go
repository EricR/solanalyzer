package analyzers

import (
	"fmt"
	"github.com/ericr/solanalyzer/sources"
)

// FunctionVisibilityAnalyzer is an analyzer that reports issues related to
// function visibility.
type FunctionVisibilityAnalyzer struct{}

// Name returns the name of the analyzer.
func (fva *FunctionVisibilityAnalyzer) Name() string {
	return "Function Visibility"
}

// ID returns the unique ID of the analyzer.
func (fva *FunctionVisibilityAnalyzer) ID() string {
	return "function-visibility"
}

// Execute runs the analyzer on a given source.
func (fva *FunctionVisibilityAnalyzer) Execute(source *sources.Source) ([]*Issue, error) {
	issues := []*Issue{}

	for _, contract := range source.Contracts {
		for _, function := range contract.Functions {
			// Skip constructor functions
			if function.Identifier == "" {
				continue
			}

			modifiers := function.Modifiers

			if !modifiers.Public && !modifiers.Private && !modifiers.Internal &&
				!modifiers.External {
				msg := fmt.Sprintf("No visibility is specified for function %s in "+
					"contract %s. The default is public. It should be confirmed that "+
					"this is desired, and the visibility of the function should be "+
					"explicitly set.", function.ShortSignature(), contract)

				issues = append(issues, &Issue{
					Severity:    SeverityInfo,
					Title:       "Default Function Visibility",
					MsgFormat:   "txt",
					Message:     msg,
					analyzer:    fva,
					sourcePath:  source.FilePath,
					sourceStart: function.Start,
					sourceStop:  function.Stop,
				})
			}
		}
	}
	return issues, nil
}
