package reports

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

// ConsoleReport is a type of text-based report that writes to stdout.
type ConsoleReport struct{}

// Generate generates a text-based report.
func (cr *ConsoleReport) Generate(report *Report) {
	issues := sortedIssues(report.Issues)

	width, _, err := terminal.GetSize(0)
	if err != nil {
		width = 70
	}

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

		cr.println("")

		for _, issue := range issues[i] {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetColWidth(width - 18)
			table.SetBorder(false)
			table.SetColumnSeparator("")
			table.Append([]string{"Title", issue.Title})

			switch issue.Format {
			case "txt":
				table.Append([]string{"Description", issue.Message})

			case "dot":
				table.Append([]string{"Description",
					processDot(issue.AnalyzerID(), []byte(issue.Message))})
			}

			table.Append([]string{"Source", issue.Source()})
			table.Append([]string{"Analyzer ID", issue.AnalyzerID()})
			table.Append([]string{"Instance ID", issue.ID()})
			table.Render()

			cr.println("")
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
