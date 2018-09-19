package ast

// Group is an abstract type that is used in the grammar of the form:
// `(` { element }* `)`
type Group struct {
	Elements []Expression
}
