package analyzers

import (
	"fmt"
	"github.com/ericr/solanalyzer/scanner"
)

// FunctionVisibilityAnalyzer is an analyzer that reports issues related to
// function visibility.
type FunctionVisibilityAnalyzer struct{}

// GetName returns the name of the analyzer.
func (fva *FunctionVisibilityAnalyzer) GetName() string {
	return "Function Visibility"
}

// GetID returns the unique ID of the analyer.
func (fva *FunctionVisibilityAnalyzer) GetID() string {
	return "function-visibility"
}

// Execute runs the analyzer on a given source.
func (fva *FunctionVisibilityAnalyzer) Execute(source *scanner.Source) ([]*Issue, error) {
	issues := []*Issue{}

	for _, contract := range source.Contracts {
		for _, function := range contract.Functions {
			if function.Public && function.Identifier != "" {
				msg := fmt.Sprintf("The function %s in the contract %s was found to be public. "+
					"It should be confirmed that this function is intended to be publicly callable.",
					function.GetSignature(), contract.Identifier,
				)

				issues = append(issues, &Issue{
					Severity:   SeverityInfo,
					Title:      "Public Function",
					MsgFormat:  "txt",
					Message:    msg,
					analyzer:   fva,
					sourcePath: source.FilePath,
					tokens:     function.Tokens,
				})
			}
		}
	}
	return issues, nil
}
