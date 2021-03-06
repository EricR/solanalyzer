package analyzers

import (
	"crypto/sha256"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Issue contains output generated by an analyzer.
type Issue struct {
	Severity    int
	Title       string
	Format      string
	Message     string
	sourcePath  string
	sourceStart antlr.Token
	sourceStop  antlr.Token
	analyzer    Analyzer
}

const (
	// SeverityInfo represents informational severity
	SeverityInfo int = iota
	// SeverityLow represents low severity
	SeverityLow
	// SeverityMed represents medium severity
	SeverityMed
	// SeverityHigh represents high severity
	SeverityHigh
)

// ID returns the ID associated with an issue.
func (i *Issue) ID() string {
	id := fmt.Sprintf("%s:%s:%s", i.AnalyzerID(), i.Source(), i.Message)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(id)))
}

// AnalyzerID returns the ID of an analyzer associated with an issue.
func (i *Issue) AnalyzerID() string {
	return i.analyzer.ID()
}

// Source returns the source location of an issue.
func (i *Issue) Source() string {
	if i.sourceStart == nil {
		return fmt.Sprintf("%s:1:0", i.sourcePath)
	}
	return fmt.Sprintf("%s:%d:%d",
		i.sourcePath, i.sourceStart.GetLine(), i.sourceStart.GetColumn())
}
