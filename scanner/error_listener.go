package scanner

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirupsen/logrus"
	"os"
)

// ErrorListener is a custom error listener for the parser.
type ErrorListener struct {
	*antlr.DefaultErrorListener
	SourceFilePath string
}

// SyntaxError is called when the parser's error listener reports a syntax error.
func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {

	logrus.Errorf("Error parsing %s:%d:%d: %s", el.SourceFilePath, line, column, msg)
	os.Exit(1)
}
