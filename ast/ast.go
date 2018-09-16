package ast

type NodeType int
type MetaType int

const (
	ProgramType NodeType = iota + 1
	BlockType
	StatementType
	ExpressionType

	AssignType
	InitType
	SetType

	AddType
	SubType
	MultType
	DivType

	IdentType
	TypeType

	IntMeta MetaType = iota + 1
	StringMeta
	BoolMeta
	FloatMeta
	TokenMeta
)

type Meta struct {
	Type  MetaType
	Value interface{}
}

type Pos struct {
	Line   int
	Column int
}

type Location struct {
	Start *Pos
	End   *Pos
}

type Nodes *[]*Node

type Node interface {
	TokenLiteral() string
	// String() string

	// Type() NodeType
	// Nodes() Nodes
	// Length() int
	// Location() *Location
	// Metadata() *map[string]Meta
}
