package ast

// Expression is an abstract type in Express that represents a combination
// of variables, functions, and other values that produces a value
type Expression interface {
	Node

	// This is just something to force the interface
	expressionNode()

	// TODO: this should have a Type function
	// TODO: this should have a Length function
}
