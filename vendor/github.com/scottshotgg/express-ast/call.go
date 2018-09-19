package ast

// FIXME: need to think about this more

// Call represents the following form:
// [ ident ] [ group ]
type Call struct {
	Token     Token
	Ident     *Ident
	Arguments []Expression
	Returns   []Expression
}

func (c *Call) expressionNode() {}
func (c *Call) statementNode()  {}

// TokenLiteral returns the literal value of the token
func (c *Call) TokenLiteral() string { return c.Token.Literal }
