package ast

// Literal is an abstract type that represents a literal value, in constrast with a value-producer, such as an expression
type Literal interface {
	Expression
	Type() Type
}

// Literals should have acting types and acting values that get set when the value is set

// IntLiteral represents any non floating-point number
type IntLiteral struct {
	Token  Token
	TypeOf Type
	Value  int
}

func (il *IntLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }

// Type implements literal
func (il *IntLiteral) Type() Type { return il.TypeOf }

func (il *IntLiteral) Kind() NodeType { return LiteralNode }

// BoolLiteral represents a variable that is restricted to either a true or false value
type BoolLiteral struct {
	Token  Token
	TypeOf Type
	Value  bool
}

func (bl *BoolLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }

// Type implements literal
func (bl *BoolLiteral) Type() Type { return bl.TypeOf }

func (bl *BoolLiteral) Kind() NodeType { return LiteralNode }

// FloatLiteral represents any floating point number
type FloatLiteral struct {
	Token  Token
	TypeOf Type
	Value  float64
}

func (fl *FloatLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }

// Type implements literal
func (fl *FloatLiteral) Type() Type { return fl.TypeOf }

func (fl *FloatLiteral) Kind() NodeType { return LiteralNode }

// CharLiteral represents a single-character capped string:
// `'` [ _single_character_ ] `'`
type CharLiteral struct {
	Token  Token
	TypeOf Type
	Value  rune
}

func (cl *CharLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (cl *CharLiteral) TokenLiteral() string { return cl.Token.Literal }

// Type implements literal
func (cl *CharLiteral) Type() Type { return cl.TypeOf }

func (cl *CharLiteral) Kind() NodeType { return LiteralNode }

// StringLiteral represents a double quoted body of text:
// TODO: how to do a backtick quoted body of text
// `"` [ _text_ ] `"`
type StringLiteral struct {
	Token  Token
	TypeOf Type
	Value  string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// Type implements literal
func (sl *StringLiteral) Type() Type { return sl.TypeOf }

func (sl *StringLiteral) Kind() NodeType { return LiteralNode }

// VarLiteral represents a dynamically typed variable; it can hold anything
type VarLiteral struct {
	Token  Token
	TypeOf Type
	// TODO: could either do it this way or this can reference another literal type
	Value interface{}
}

func (vl *VarLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (vl *VarLiteral) TokenLiteral() string { return vl.Token.Literal }

// Type implements literal
func (vl *VarLiteral) Type() Type { return vl.TypeOf }

func (vl *VarLiteral) Kind() NodeType { return LiteralNode }

// ObjectLiteral represents a named block : this produces a variable
type ObjectLiteral struct {
	Token  Token
	TypeOf Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal
	Value Block

	// Only allow assignment operations inside objects for now
	// Value map[string]Literal
}

func (ol *ObjectLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (ol *ObjectLiteral) TokenLiteral() string { return ol.Token.Literal }

// Type implements literal
func (ol *ObjectLiteral) Type() Type { return ol.TypeOf }

func (ol *ObjectLiteral) Kind() NodeType { return LiteralNode }

// StructLiteral represents a named object : this produces a type
// TODO: this might need to be moved to the type.go file
// FIXME: this might need to be fixed or something
type StructLiteral struct {
	Token  Token
	TypeOf Type
	Value  map[string]Expression
}

func (sl *StructLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (sl *StructLiteral) TokenLiteral() string { return sl.Token.Literal }

// Type implements literal
func (sl *StructLiteral) Type() Type { return sl.TypeOf }

func (sl *StructLiteral) Kind() NodeType { return LiteralNode }

// FunctionLiteral represents a named object : this produces a type
type FunctionLiteral struct {
	Token  Token
	TypeOf Type
	// TODO: could either do it this way or make block implement literal and then it can be directly used as a literal

	// On the backend, a function would essentially just be a block (i.e, object) that is able to be called
	Value Block
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// Type implements literal
func (fl *FunctionLiteral) Type() Type { return fl.TypeOf }

func (fl *FunctionLiteral) Kind() NodeType { return LiteralNode }

// func NewLiteral(t Token, v interface{}) *Literal {

// }

// NewInt returns a new int literal
func NewInt(token Token, value int) *IntLiteral {
	return &IntLiteral{
		Token:  token,
		TypeOf: NewIntType(),
		Value:  value,
	}
}

// Make this take a type and initialize the default from the default map
// func NewDefault(token Token) *IntLiteral {
// 	return NewIntFromValue(token, 0)
// }

// NewBool returns a new bool literal
func NewBool(token Token, value bool) *BoolLiteral {
	return &BoolLiteral{
		Token:  token,
		TypeOf: NewBoolType(),
		Value:  value,
	}
}

// NewFloat returns a new float literal
func NewFloat(token Token, value float64) *FloatLiteral {
	return &FloatLiteral{
		Token:  token,
		TypeOf: NewFloatType(),
		Value:  value,
	}
}

// NewChar returns a new char literal
func NewChar(token Token, value rune) *CharLiteral {
	return &CharLiteral{
		Token:  token,
		TypeOf: NewCharType(),
		Value:  value,
	}
}

// NewString returns a new string literal
func NewString(token Token, value string) *StringLiteral {
	return &StringLiteral{
		Token:  token,
		TypeOf: NewStringType(),
		Value:  value,
	}
}

// NewStruct returns a new struct literal
func NewStruct(token Token, structType LiteralType, value map[string]Expression) *StructLiteral {
	return &StructLiteral{
		Token:  token,
		TypeOf: NewStructType(structType),

		// This is for the properties of the struct, but somehow we probably need to have a
		// UserDefinedValueMap like we have for the UserDefinedTypeMap
		Value: value,
	}
}

// NewObject returns a new object literal
func NewObject(token Token, value Block) *ObjectLiteral {
	return &ObjectLiteral{
		Token:  token,
		TypeOf: NewObjectType(),
		Value:  value,
	}
}

// NewVarFromInt returns a new int shadow-typed var
func NewVarFromInt(token Token, value int) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(IntType),
		Value:  value,
	}
}

// NewVarFromBool returns a new bool shadow-typed var
func NewVarFromBool(token Token, value bool) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(BoolType),
		Value:  value,
	}
}

// NewVarFromFloat returns a new float shadow-typed var
func NewVarFromFloat(token Token, value float64) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(FloatType),
		Value:  value,
	}
}

// NewVarFromChar returns a new char shadow-typed var
func NewVarFromChar(token Token, value rune) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(CharType),
		Value:  value,
	}
}

// NewVarFromString returns a new string shadow-typed var
func NewVarFromString(token Token, value string) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(StringType),
		Value:  value,
	}
}

// NewVarFromObject returns a new object shadow-typed var
func NewVarFromObject(token Token, value Block) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(ObjectType),
		Value:  value,
	}
}

// // NewVarFromStruct returns a new struct shadow-typed var
// func NewVarFromStruct(token Token, structType LiteralType, value map[string]Expression) *VarLiteral {
// 	return &VarLiteral{
// 		Token:  token,
// 		TypeOf: NewVarType(StructType),
// 		Value:  value,
// 	}
// }

// NewVarFromFunction returns a new function shadow-typed var
func NewVarFromFunction(token Token, value Block) *VarLiteral {
	return &VarLiteral{
		Token:  token,
		TypeOf: NewVarType(FunctionType),
		Value:  value,
	}
}
