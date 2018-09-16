package ast

type AssignmentOpType int

const (
	StandardAssignType AssignmentOpType = iota + 1
	InitAssignType
	SetAssignType
)

// [ IDENT ] = [ EXPR ]
type AssignmentStatement struct {
	Ident          *Identifier
	AssignmentType AssignmentOpType
	Value          Expression
}

func (a *AssignmentStatement) statementNode() {}
