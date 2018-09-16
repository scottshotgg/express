package ast

// FIXME: change this to an interface like ExpressionNode

// type StatementNode interface {
// 	Kind         NodeType
// 	Children     Nodes
// 	NodeLength   int
// 	NodeLocation *Location
// 	NodeMetadata *map[string]Meta
// }

// func (n *StatementNode) Type() NodeType {
// 	return n.Kind
// }

// func (n *StatementNode) Nodes() Nodes {
// 	return n.Children
// }

// func (n *StatementNode) Length() int {
// 	return n.NodeLength
// }

// func (n *StatementNode) Location() *Location {
// 	return n.NodeLocation
// }

// func (n *StatementNode) Metadata() *map[string]Meta {
// 	return n.NodeMetadata
// }

// func NewStatementNode(location *Location) Node {
// 	return &StatementNode{
// 		Kind:         StatementType,
// 		NodeLocation: location,
// 	}
// }

type Statement interface {
	Node
	statementNode()
}
