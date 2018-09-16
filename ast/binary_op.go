package ast

// type BinaryOpNodeType int

// const (
// 	AdditionOpType BinaryOpNodeType = iota
// 	SubtractionOpType
// )

// type BinaryOpNode interface {
// 	ExpressionNode

// 	// For right now they return abstract Node types
// 	Right() *ExpressionNode
// 	Left() *ExpressionNode
// 	OperationType() BinaryOpNodeType
// }

// type BinaryOp struct {
// 	*Expression

// 	LeftNode  ExpressionNode
// 	RightNode ExpressionNode
// 	OpType    BinaryOpNodeType
// }

// func (b *BinaryOp) Left() ExpressionNode {
// 	return b.LeftNode
// }

// func (b *BinaryOp) Right() ExpressionNode {
// 	return b.RightNode
// }

// func (b *BinaryOp) OperationType() BinaryOpNodeType {
// 	return b.OpType
// }
