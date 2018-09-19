package ast

// Position is used to specify where in the code the token for the node was found
type Position struct {
	Line   int
	Column int
}

// Location holds the start and end positions of the node
type Location struct {
	Start *Position
	End   *Position
}

// TODO: we should make more utilization functions around this

// Node is an abstract type that is used to recursively define the AST
type Node interface {
	// TODO: this will just be a string for now until I rework the lexer
	TokenLiteral() string

	// Location() map[string]*Location
}
