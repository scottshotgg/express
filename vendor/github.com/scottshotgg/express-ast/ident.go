package ast

import (
	"errors"
)

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

func (i *Ident) Kind() NodeType { return IdentNode }

// Might need to make specific type-functions
// But I don't think identifiers here need to have a type, that's NOT what the AST is for; keep track of that in the parser, etc

// NewIdent returns a new identifier
// func NewIdent(t Token, it Type, n string) (*Ident, error) {
func NewIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		// Type:  it,
		Name: n,
	}, nil
}

// NewIntIdent returns a new identifier for an int type
func NewIntIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewIntType(),
		Name:  n,
	}, nil
}

// NewBoolIdent returns a new identifier for an bool type
func NewBoolIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewBoolType(),
		Name:  n,
	}, nil
}

// NewFloatIdent returns a new identifier for an float type
func NewFloatIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewFloatType(),
		Name:  n,
	}, nil
}

// NewCharIdent returns a new identifier for an char type
func NewCharIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewCharType(),
		Name:  n,
	}, nil
}

// NewStringIdent returns a new identifier for an string type
func NewStringIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewStringType(),
		Name:  n,
	}, nil
}

// func NewStructIdent(t Token, n string) (*Ident, error) {
// 	if n == "" {
// 		return nil, errors.New("Cannot use empty string as identifier name")
// 	}

// 	return &Ident{
// 		Token: t,
// 		Type:  NewStructType(),
// 		Name:  n,
// 	}, nil
// }

// NewObjectIdent returns a new identifier for an object type
func NewObjectIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewObjectType(),
		Name:  n,
	}, nil
}

// NewFunctionIdent returns a new identifier for an function type
func NewFunctionIdent(t Token, n string) (*Ident, error) {
	if n == "" {
		return nil, errors.New("Cannot use empty string as identifier name")
	}

	return &Ident{
		Token: t,
		Type:  NewFunctionType(),
		Name:  n,
	}, nil
}

// func NewVarIdent(t Token, n string) (*Ident, error) {
// 	if n == "" {
// 		return nil, errors.New("Cannot use empty string as identifier name")
// 	}

// 	return &Ident{
// 		Token: t,
// 		Type:  NewVarType(),
// 		Name:  n,
// 	}, nil
// }
