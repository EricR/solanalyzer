package sessions

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericr/solanalyzer/analyzers"
	"github.com/ericr/solanalyzer/parser"
	"github.com/ericr/solanalyzer/reports"
	"github.com/ericr/solanalyzer/sources"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Session represents a session of running the tool, referencing source units,
// analyzers, and issues found.
type Session struct {
	Sources   []*sources.Source
	Analyzers []analyzers.Analyzer
	Issues    []*analyzers.Issue
}

// NewSession returns a new instance of Session.
func NewSession() *Session {
	return &Session{
		Sources: []*sources.Source{},
	}
}

// ParsePath walks a directory containing Solidity source files and parses each
// one.
func (s *Session) ParsePath(paths []string) {
	files := []string{}

	for _, path := range paths {
		if isDir(path) {
			filepath.Walk(path, s.pathWalkFunc)
		} else {
			files = append(files, path)
		}
	}

	for _, file := range files {
		s.ParseFile(file)
	}
}

// ParseFile parses a Solidity source file.
func (s *Session) ParseFile(path string) error {
	logrus.Debugf("Parsing %s", path)

	fs, err := antlr.NewFileStream(path)
	if err != nil {
		return err
	}

	solLexer := parser.NewSolidityLexer(fs)
	stream := antlr.NewCommonTokenStream(solLexer, antlr.TokenDefaultChannel)
	solParser := parser.NewSolidityParser(stream)

	solParser.RemoveErrorListeners()
	solParser.AddErrorListener(&sources.ErrorListener{SourceFilePath: path})

	s.addSource(path, solParser.SourceUnit().(*parser.SourceUnitContext))

	return nil
}

// VisitSources "visits" all sources trees.
func (s *Session) VisitSources() {
	for _, source := range s.Sources {
		logrus.Debugf("Scanning %s", source.FilePath)
		source.Visit()
	}
}

// AddAnalyzer adds a new analyzer to be run during the session.
func (s *Session) AddAnalyzer(analyzer analyzers.Analyzer) {
	s.Analyzers = append(s.Analyzers, analyzer)
}

// Analyze runs all analyzers on all sources.
func (s *Session) Analyze() {
	for _, analyzer := range s.Analyzers {
		for _, source := range s.Sources {
			logrus.Debugf("Analyzing %s in %s", analyzer.Name(), source)
			newIssues, err := analyzer.Execute(source)
			if err != nil {
				logrus.Errorf("Got error from analyzer: %s", err)
			}

			s.Issues = append(s.Issues, newIssues...)
		}
	}
}

// GenerateReport generates a report with a given generator.
func (s *Session) GenerateReport(generator reports.Generator) {
	report := reports.New(generator)
	report.AddSources(s.Sources)
	report.AddAnalyzers(s.Analyzers)
	report.AddIssues(s.Issues)
	report.Generate()
}

func (s *Session) addSource(path string, tree *parser.SourceUnitContext) {
	s.Sources = append(s.Sources, sources.New(path, tree))
}

func (s *Session) pathWalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		logrus.Errorf("Got error reading %s: %s", path, err)
		return err
	}

	if !info.IsDir() && filepath.Ext(path) == ".sol" {
		s.ParseFile(path)
	}

	return nil
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		logrus.Errorf("Got error reading %s: %s", path, err)
		return false
	}

	return fileInfo.IsDir()
}
