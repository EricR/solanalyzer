package reports

import (
	"github.com/ericr/solanalyzer/analyzers"
	"github.com/ericr/solanalyzer/sources"
	"strings"
	"time"
)

// Report represents a report of issues found.
type Report struct {
	GeneratedAt time.Time
	Analyzers   []analyzers.Analyzer
	Sources     []*sources.Source
	Issues      []*analyzers.Issue
	generator   Generator
}

// Generator is a type that can generate reports.
type Generator interface {
	Generate(*Report)
}

// New returns a new instance of Report.
func New(generator Generator) *Report {
	return &Report{
		generator: generator,
	}
}

// AddSources adds a list of sources to the report.
func (r *Report) AddSources(sources []*sources.Source) {
	r.Sources = append(r.Sources, sources...)
}

// AddAnalyzers adds a list of analyzers to the report.
func (r *Report) AddAnalyzers(analyzers []analyzers.Analyzer) {
	r.Analyzers = append(r.Analyzers, analyzers...)
}

// AddIssues adds a list of issues to a report.
func (r *Report) AddIssues(issues []*analyzers.Issue) {
	r.Issues = append(r.Issues, issues...)
}

// Generate generates a report.
func (r *Report) Generate() {
	r.GeneratedAt = time.Now()
	r.generator.Generate(r)
}

// AnalyzersRun returns a list of IDs associated with the analyzers that were
// run.
func (r *Report) AnalyzersRun() string {
	ids := []string{}

	for _, analyzer := range r.Analyzers {
		ids = append(ids, analyzer.ID())
	}

	return strings.Join(ids, ", ")
}

func sortedIssues(issues []*analyzers.Issue) map[int][]*analyzers.Issue {
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
