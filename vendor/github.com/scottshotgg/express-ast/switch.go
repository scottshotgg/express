package ast

// Switch statements represents the following form:
// `switch` { expression } [ case_block ]
type Switch struct {
	Token      Token
	Expression Expression
	Cases      *CaseBlock
	Default    Statement
}

// CaseBlock represents the following form:
// `{` [ case ]* `}`
type CaseBlock struct {
	Cases []Case
}

// Case represents the following form:
// `case` [ expression ] `:` [ block ]
type Case struct {
	Token      Token
	Expression Expression
	Body       Statement
}

// Implement Node and Statement

func (s *Switch) statementNode() {}

// TokenLiteral returns the literal value of the token
func (s *Switch) TokenLiteral() string { return s.Token.Literal }

func (s *Switch) Kind() NodeType { return SwitchNode }
