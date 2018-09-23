package ast

// Function represents the following form:
// [ `func` | `fn` ] [ ident ] [ group ] { group } [ block ]
type Function struct {
	Lambda    bool
	Async     bool
	Token     Token
	Name      string
	Arguments *Group
	Returns   *Group
	Body      Block
}

// Implement statement
func (f *Function) statementNode() {}

// Implement expression
func (f *Function) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (f *Function) TokenLiteral() string { return f.Token.Literal }

// Type implements literal so that functions can be assigned to idents
func (f *Function) Type() LiteralType { return FunctionType }

func (f *Function) Kind() NodeType { return FunctionNode }
