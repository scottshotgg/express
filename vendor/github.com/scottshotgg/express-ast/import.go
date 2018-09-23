package ast

// Import is an import statement in the form of:
// `import` [ string_lit ]
type Import struct {
	Token Token
	Name  *Ident
	Path  string
}

// Implement Node and Statement
func (i *Import) statementNode() {}

// TokenLiteral returns the literal value of the token
func (i *Import) TokenLiteral() string { return i.Token.Literal }

func (i *Import) Kind() NodeType { return ImportNode }
