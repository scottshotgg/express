package ast

import "errors"

// AssignmentType encompasses the different types of assignment
type AssignmentType int

const (
	// Equals is the = operator
	Equals AssignmentType = iota + 1

	// Set is the : operator
	Set

	// Init is the := operator
	Init
)

// Assignment statement represents the following form:
// { type } [ ident ] [ assign_op ] [ expression ]
type Assignment struct {
	Declaration bool
	Inferred    bool
	Token       Token
	Type        AssignmentType
	Ident       *Ident
	Value       Expression
}

// When going through the logic for this:
//	- if Declaration is already set to true when Inferred is being set to true -> error
//	- if the variable is already declared and Declared or Inferred is being set to true -> error

func (a *Assignment) statementNode() {}

// TokenLiteral returns the literal value of the token
func (a *Assignment) TokenLiteral() string { return a.Token.Literal }

func (a *Assignment) Kind() NodeType { return AssignmentNode }

// TODO: dont think I wanna do this yet
// func NewAssignmentStatement() Assignment {
// 	return &Assignment{

// 	}
// }

// NewAssignment returns a new assignment statement and determines whether it is inferred
func NewAssignment(t Token, i *Ident, at AssignmentType, e Expression) (*Assignment, error) {
	if e == nil {
		return nil, errors.New("Expression value cannot by nil")
	}

	as := Assignment{
		Token: t,
		Ident: i,
		Type:  at,
		Value: e,
	}

	if at == Init {
		as.Inferred = true
	}

	return &as, nil
}

// SetDeclaration changes the assignment type to a declaration
func (a *Assignment) SetDeclaration(declaration bool) {
	a.Declaration = declaration
}

// SetInferred changes the assignment type to infer the lvalue type
func (a *Assignment) SetInferred(inferred bool) {
	a.Inferred = inferred
}

// func (a *Assignment) ExpressionType() {
// 	if a.Value != nil {
// 		return
// 	}
// }
