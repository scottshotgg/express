package ast

// Statement is an abstract type that represents a complete sentence in Express
type Statement interface {
	Node

	// This is just something to force the interface
	statementNode()

	// TODO: this should have a Type function
}
