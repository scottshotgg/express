package ast

// IfElse represents the following form:
// if [ condition ] [ block ] { [ else ] [ statement ] }
type IfElse struct {
	Token         Token
	IfCondition   *Condition
	If            *Block
	ElseCondition *Condition
	// TODO: Hmmm this is supposed to only be a block or another if statement
	// but should we try to bound it?
	Else *ElseStatement
}

func (ie *IfElse) statementNode()     {}
func (ie *IfElse) elseStatementNode() {}

// TokenLiteral returns the literal value of the token
func (ie *IfElse) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfElse) Kind() NodeType { return IfElseNode }
