package scanner

// Import represents a Solidity import.
type Import struct {
	Identifier string
	As         string
	From       string
}
