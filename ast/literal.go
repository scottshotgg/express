package ast

type IntegerLiteral struct {
	ShadowType VariableType
	Value      int64
}

func (il *IntegerLiteral) expressionNode() {}

type FloatLiteral struct {
	Type VariableType
}
