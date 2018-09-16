package ast

// type ProgramNode struct {
// 	Kind         NodeType
// 	Children     Nodes
// 	NodeLength   int
// 	NodeLocation *Location
// 	NodeMetadata *map[string]Meta
// }

// func (n *ProgramNode) Type() NodeType {
// 	return n.Kind
// }

// func (n *ProgramNode) Nodes() Nodes {
// 	return n.Children
// }

// func (n *ProgramNode) Length() int {
// 	return n.NodeLength
// }

// func (n *ProgramNode) Location() *Location {
// 	return n.NodeLocation
// }

// func (n *ProgramNode) Metadata() *map[string]Meta {
// 	return n.NodeMetadata
// }

// func NewProgramNode() *ProgramNode {
// 	return &ProgramNode{
// 		Kind: ProgramType,
// 	}
// }

// func (n *ProgramNode) NewStatement() {

// }

type Program struct {
	Statements []Statement
}
