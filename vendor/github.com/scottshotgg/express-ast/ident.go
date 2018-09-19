package ast

import "errors"

// Ident represents the following form:
// [ name ]
type Ident struct {
	Token Token
	Type  Type
	Name  string
}

func (i *Ident) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (i *Ident) TokenLiteral() string { return i.Token.Literal }

// Might need to make specific type-functions
// But I don't think identifiers here need to have a type, that's NOT what the AST is for; keep track of that in the parser, etc

func NewIdent(t Token, it Type, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  it,
		Name:  n,
	}, nil
}
