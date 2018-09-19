package ast

// Return represents the following form:
// `return` [ expression ]
type Return struct {
	Token Token
	Value []Expression
}

func (r *Return) statementNode() {}

// TokenLiteral returns the literal value of the token
func (r *Return) TokenLiteral() string { return r.Token.Literal }
