package sessions

import (
	"github.com/ericr/solanalyzer/analyzers"
	"github.com/ericr/solanalyzer/sources"
	"testing"
)

type TestAnalyzer struct{}

func (ta *TestAnalyzer) Name() string {
	return "test"
}

func (ta *TestAnalyzer) ID() string {
	return "test"
}

func (ta *TestAnalyzer) Execute(*sources.Source) ([]*analyzers.Issue, error) {
	return []*analyzers.Issue{}, nil
}

func TestUniqueSources(t *testing.T) {
	session := NewSession()
	session.Parse("file.sol", "")
	session.Parse("file.sol", "")

	if len(session.Sources) > 1 {
		t.Error("Duplicate sources were added to the session")
	}
}

func TestUniqueAnalyzers(t *testing.T) {
	session := NewSession()
	session.AddAnalyzer(&TestAnalyzer{})
	session.AddAnalyzer(&TestAnalyzer{})

	if len(session.Sources) > 1 {
		t.Error("Duplicate analyzers were added to the session")
	}
}
