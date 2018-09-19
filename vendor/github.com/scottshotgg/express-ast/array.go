package ast

type Array struct {
	Type LiteralType
}

// Implement Literal
func (a *Array) TypeOf() LiteralType {
	return a.Type
}
