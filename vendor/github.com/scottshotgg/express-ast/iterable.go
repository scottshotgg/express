package ast

// Iterable is an abstract type in Express that represents the
// ability for an expression (specifically a literal) to be iterated over
type Iterable interface {
	Node

	// This is just something to force the interface
	expressionNode()

	Next() *Literal
	Prev() *Literal

	Fields() []*Literal

	// TODO: this should have a Type function
	// TODO: this should have a Length function
}
