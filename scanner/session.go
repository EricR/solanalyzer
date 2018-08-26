package scanner

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ericr/solanalyzer/parser"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Session struct {
	Scanner *Scanner
	Sources []*Source
}

// NewSession returns a new scanner session.
func NewSession() *Session {
	return &Session{
		Scanner: NewScanner(),
		Sources: []*Source{},
	}
}

// ParsePath walks a directory containing solidity source files and parses each.
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

// ParseFile parses a solidity source file.
func (s *Session) ParseFile(path string) error {
	fs, err := antlr.NewFileStream(path)
	if err != nil {
		return err
	}

	solLexer := parser.NewSolidityLexer(fs)
	stream := antlr.NewCommonTokenStream(solLexer, antlr.TokenDefaultChannel)
	solParser := parser.NewSolidityParser(stream)

	solParser.RemoveErrorListeners()
	solParser.AddErrorListener(&ErrorListener{SourceFilePath: path})

	s.addSource(path, solParser.SourceUnit())

	return nil
}

// Scan scans all sources.
func (s *Session) Scan() {
	for _, source := range s.Sources {
		s.Scanner.Scan(source)
	}
}

func (s *Session) addSource(path string, tree antlr.Tree) {
	s.Sources = append(s.Sources, NewSource(path, tree))
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		logrus.Errorf("Got error reading %s: %s", path, err)
		return false
	}

	return fileInfo.IsDir()
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
