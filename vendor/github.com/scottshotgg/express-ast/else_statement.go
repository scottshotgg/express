package ast

// ElseStatement is an abstract type that represents a complete sentence in Express
type ElseStatement interface {
	Node

	// This is just something to force the interface
	elseStatementNode()
}
