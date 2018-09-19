package ast

// Literal is an abstract types that represents a literal value, in constrast with a value-producer, such as an expression
type Literal interface {
	Type() LiteralType
}

// Literals should have acting types and acting values that get set when the value is set

// IntLiteral represents any non floating-point number
type IntLiteral struct {
	Token Token
	Type  Type
	Value int
}

func (il *IntLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }

// FloatLiteral represents any floating point number
type FloatLiteral struct {
	Token Token
	Type  Type
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

// CharLiteral represents a single-character capped string:
// `'` [ _single_character_ ] `'`
type CharLiteral struct {
	Token Token
	Type  Type
	Value rune
}

func (cl *CharLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

// StringLiteral represents a double quoted body of text:
// TODO: how to do a backtick quoted body of text
// `"` [ _text_ ] `"`
type StringLiteral struct {
	Token Token
	Type  Type
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// BoolLiteral represents a variable that is restricted to either a true or false value
type BoolLiteral struct {
	Token Token
	Type  Type
	Value bool
}

func (bl *BoolLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }

// VarLiteral represents a dynamically typed variable; it can hold anything
type VarLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or this can reference another literal type
	Value interface{}
}

func (vl *VarLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (vl *VarLiteral) TokenLiteral() string { return vl.Token.Literal }

// ObjectLiteral represents a named block : this produces a variable
type ObjectLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal
	Value Block

	// Only allow assignment operations inside objects for now
	// Value map[string]Literal
}

func (ol *ObjectLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (ol *ObjectLiteral) TokenLiteral() string { return ol.Token.Literal }

// StructLiteral represents a named object : this produces a type
// TODO: this might need to be moved to the type.go file
// FIXME: this might need to be fixed or something
type StructLiteral struct {
	Token Token
	Type  Type
	Value map[string]Expression
}

func (sl *StructLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StructLiteral) TokenLiteral() string { return sl.Token.Literal }

// FunctionLiteral represents a named object : this produces a type
type FunctionLiteral struct {
	Token Token
	Type  Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal

	// On the backend, a function would essentially just be a block (i.e, object) that is able to be called
	Value Block
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// NewInt is used for making a new int literal
func NewInt(token Token, value int) *IntLiteral {
	return &IntLiteral{
		Token: token,
		Type:  NewIntType(),
		Value: value,
	}
}

// Make this take a type and initialize the default from the default map
// func NewDefault(token Token) *IntLiteral {
// 	return NewIntFromValue(token, 0)
// }

func NewBool(token Token, value bool) *BoolLiteral {
	return &BoolLiteral{
		Token: token,
		Type:  NewBoolType(),
		Value: value,
	}
}

func NewFloat(token Token, value float64) *FloatLiteral {
	return &FloatLiteral{
		Token: token,
		Type:  NewFloatType(),
		Value: value,
	}
}

func NewChar(token Token, value rune) *CharLiteral {
	return &CharLiteral{
		Token: token,
		Type:  NewCharType(),
		Value: value,
	}
}

func NewString(token Token, value string) *StringLiteral {
	return &StringLiteral{
		Token: token,
		Type:  NewStringType(),
		Value: value,
	}
}

func NewStruct(token Token, structType LiteralType, value map[string]Expression) *StructLiteral {
	return &StructLiteral{
		Token: token,
		Type:  NewStructType(structType),

		// This is for the properties of the struct, but somehow we probably need to have a
		// UserDefinedValueMap like we have for the UserDefinedTypeMap
		Value: value,
	}
}

func NewObject(token Token, value Block) *ObjectLiteral {
	return &ObjectLiteral{
		Token: token,
		Type:  NewObjectType(),
		Value: value,
	}
}

func NewVarFromInt(token Token, value int) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(IntType),
		Value: value,
	}
}

func NewVarFromBool(token Token, value bool) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(BoolType),
		Value: value,
	}
}

func NewVarFromFloat(token Token, value float64) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(FloatType),
		Value: value,
	}
}

func NewVarFromChar(token Token, value rune) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(CharType),
		Value: value,
	}
}

func NewVarFromString(token Token, value string) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(StringType),
		Value: value,
	}
}

func NewVarFromObject(token Token, value Block) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(ObjectType),
		Value: value,
	}
}

// TODO: don't know how to represent these internally
func NewVarFromStruct(token Token, structType LiteralType, value map[string]Expression) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(StructType),
		Value: value,
	}
}

// This is essentially the same thing as the block right now but it is callable
func NewVarFromFunction(token Token, value Block) *VarLiteral {
	return &VarLiteral{
		Token: token,
		Type:  NewVarType(FunctionType),
		Value: value,
	}
}
