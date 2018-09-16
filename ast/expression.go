package ast

type ExpressionNodeType int

const (
	BinaryOpExpressionType ExpressionNodeType = iota
	UnaryOpExpressionType
)

type Expression interface {
	// Node
	expressionNode()

	Left() *Expression
	Right() *Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// type ExpressionStatement struct {
// 	// Kind         NodeType
// 	// Children     Nodes
// 	// NodeLength   int
// 	// NodeLocation *Location
// 	// NodeMetadata *map[string]Meta

// 	LeftExpr  *Expression
// 	RightExpr *Expression
// }

// // func (e *Expression) ExpressionType() ExpressionNodeType {
// // 	return e.ExprType
// // }

// func (n *Expression) expressionNode() {}

// func (e *Expression) Left() *Expression {
// 	return LeftExpr
// }

// func (e *Expression) Right() *Expression {
// 	return RightExpr
// }

// // func (n *Expression) Type() NodeType {
// // 	return n.Kind
// // }

// // func (n *Expression) Nodes() Nodes {
// // 	return n.Children
// // }

// // func (n *Expression) Length() int {
// // 	return n.NodeLength
// // }

// // func (n *Expression) Location() *Location {
// // 	return n.NodeLocation
// // }

// // func (n *Expression) Metadata() *map[string]Meta {
// // 	return n.NodeMetadata
// // }
