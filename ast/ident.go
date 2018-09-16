package ast

type VariableType int

const (
	IntType VariableType = iota + 1
	BoolType
	CharType
	FloatType
	StringType
	// Add more types later
)

// ShadowType is used to:
//	a) As a separate category for types since a variables shadow type should never be dynamic
//	b) To allow variables to act as other types; `upgradeable types`: char -> string, int -> float, struct -> object
// type ShadowType int

// const (
// 	IntType ShadowType = iota + 1
// 	BoolType
// 	CharType
// 	FloatType
// 	StringType
// 	// Add more types later
// )

// [ TYPE ] [ NAME ]
type Ident struct {
	Type       VariableType
	ShadowType VariableType
	Name       string
}

func (i *Ident) expressionNode() {}
