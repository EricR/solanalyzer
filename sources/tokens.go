package sources

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Tokens store the start and stop of a parse tree node.
type Tokens struct {
	Start antlr.Token
	Stop  antlr.Token
}
