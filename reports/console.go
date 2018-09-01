package reports

import (
	"fmt"
	"github.com/ericr/solanalyzer/analyzers"
	"strings"
)

// ConsoleReport is a type of text-based report that writes to stdout.
type ConsoleReport struct{}

// Generate generates a text-based report.
func (cr *ConsoleReport) Generate(report *Report) {
	issues := sortedIssues(report.Issues)

	cr.println("")
	cr.println("=== Start SolAnalyzer Report ===")
	cr.println("")
	cr.println("Report Date:   %s", report.GeneratedAt.Format("Mon Jan _2 3:04 PM 2006"))
	cr.println("Analyzers Run: %s", report.AnalyzersRun())
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
	cr.println("Source:      %s", issue.Source())
	cr.println("Analyzer ID: %s", issue.AnalyzerID())
	cr.println("Instance ID: %s", issue.ID())
	cr.println("")
}
