package astmap

// type NodeType int
// type MetaType int

// const (
// 	ProgramType NodeType = iota + 1
// 	BlockType
// 	StatementType

// 	AssignType
// 	InitAssignType
// 	SetAssignType

// 	AddType
// 	SubType
// 	MultType
// 	DivType

// 	IdentType
// 	TypeType

// 	IntMeta MetaType = iota + 1
// 	StringMeta
// 	BoolMeta
// 	FloatMeta
// 	TokenMeta
// )

// type Meta struct {
// 	Type  MetaType
// 	Value interface{}
// }

// // type Pos struct {
// // 	Line   int
// // 	Column int
// // }

// // type Location struct {
// // 	Start *Pos
// // 	End   *Pos
// // }

// type Node struct {
// 	Info map[string]*Meta
// }

// func (n *Node) Type() *Meta {
// 	return n.Info["type"]
// }

// func (n *Node) Nodes() Nodes {
// 	return n.Children
// }

// func (n *Node) Length() int {
// 	return n.NodeLength
// }

// func (n *Node) Location() *Location {
// 	return n.NodeLocation
// }

// func (n *Node) Metadata() *map[string]Meta {
// 	return n.NodeMetadata
// }

// // func NewProgramNode() Node {
// // 	return &ProgramNode{
// // 		Kind: ProgramType,
// // 	}
// // }
