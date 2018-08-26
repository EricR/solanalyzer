package reports

import "github.com/ericr/solanalyzer/analyzers"

// Report is a type that can generate a report.
type Report interface {
	Generate([]*analyzers.Issue)
}

func getSortedIssues(issues []*analyzers.Issue) map[int][]*analyzers.Issue {
	sortedIssues := map[int][]*analyzers.Issue{
		0: []*analyzers.Issue{}, // informational
		1: []*analyzers.Issue{}, // low
		2: []*analyzers.Issue{}, // medium
		3: []*analyzers.Issue{}, // high
	}

	for _, issue := range issues {
		sortedIssues[issue.Severity] = append(sortedIssues[issue.Severity], issue)
	}

	return sortedIssues
}

func severityAsWord(i int) string {
	return map[int]string{
		0: "Informational",
		1: "Low",
		2: "Medium",
		3: "High",
	}[i]
}
