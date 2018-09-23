package ast

// Array represents array type data structures
type Array struct {
	Token Token
	// How will this act with `var` elements?
	TypeOf LiteralType
	Length int
}

func (a *Array) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (a *Array) TokenLiteral() string { return a.Token.Literal }

// Type implements Literal
func (a *Array) Type() LiteralType { return a.TypeOf }

func (a *Array) Kind() NodeType { return ArrayNode }
