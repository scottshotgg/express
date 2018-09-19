package ast

// BinaryOpType encompasses the types of binary operations
type BinaryOpType int

const (
	// AdditionBinaryOp is the + operator
	AdditionBinaryOp BinaryOpType = iota + 1

	// SubtractionBinaryOp is the - operator
	SubtractionBinaryOp

	// MultiplicationBinaryOp is the * operator
	MultiplicationBinaryOp

	// DivisionBinaryOp is the / operator
	DivisionBinaryOp
)

// type BinaryOp interface {
// 	Expression

// 	Type() BinaryOpType
// 	Right() *Expression
// 	Left() *Expression
// 	Evaluate() *Literal
// }

// BinaryOperation represents the following form:
// [ expression ] [ binary_op ] [ expression ]
type BinaryOperation struct {
	Token     Token
	Kind      BinaryOpType
	LeftNode  Expression
	RightNode Expression
	Value     Literal
}

// func (b *BinaryOperation) Type() *Expression     { return b.Kind }
// func (b *BinaryOperation) Right() *Expression    { return b.RightExpr }
// func (b *BinaryOperation) Left() *Expression     { return b.LeftExpr }
// func (b *BinaryOperation) Evaluate() *Expression { return b.Value }

func (b *BinaryOperation) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (b *BinaryOperation) TokenLiteral() string { return b.Token.Literal }
