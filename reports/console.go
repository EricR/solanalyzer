package reports

import (
	"fmt"
	"github.com/ericr/solanalyzer/analyzers"
	"strings"
	"time"
)

// ConsoleReport is a type of report that writes to standard out.
type ConsoleReport struct{}

// NewConsoleReport returns a new instance of ConsoleReport.
func NewConsoleReport() *ConsoleReport {
	return &ConsoleReport{}
}

// Generate generates a text-based report.
func (cr *ConsoleReport) Generate(group *analyzers.Group) {
	generatedAt := time.Now().Format("Mon Jan _2 15:04 2006")
	analyzersRun := strings.Join(group.AnalyzerNames, ", ")
	issues := getSortedIssues(group.Issues)

	cr.println("")
	cr.println("=== Start SolAnalyzer Report ===")
	cr.println("")
	cr.println("Report Date:   %s", generatedAt)
	cr.println("Analyzers Run: %s", analyzersRun)
	cr.println("")

	for i := len(issues) - 1; i >= 0; i-- {
		cr.printHeader(fmt.Sprintf("%s Severity Issues", severityAsWord(i)))

		if len(issues[i]) == 0 {
			cr.println("No issues")
			cr.println("")

			continue
		}

		for _, issue := range issues[i] {
			cr.printIssue(issue)
		}
	}

	cr.println("=== End SolAnalyzer Report ===")
	cr.println("")
}

func (cr *ConsoleReport) printHeader(str string) {
	fmt.Println(str)
	fmt.Println(strings.Repeat("-", len(str)))
}

func (cr *ConsoleReport) println(str string, vars ...interface{}) {
	fmt.Println(fmt.Sprintf(str, vars...))
}

func (cr *ConsoleReport) printIssue(issue *analyzers.Issue) {
	cr.println("Title:       %s", issue.Title)
	cr.println("Description: %s", issue.Message)
	cr.println("Source:      %s", issue.GetSource())
	cr.println("Analyzer ID: %s", issue.GetAnalyzerID())
	cr.println("Instance ID: %s", issue.GetID())
	cr.println("")
}
