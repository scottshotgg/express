package ast

// ConditionType encompasses all types of conditions
type ConditionType int

const (
	// EqualTo is the == operator
	EqualTo ConditionType = iota + 1

	// StrictlyEqualTo is the === operator that compares both type and value for dynamic variables
	StrictlyEqualTo

	// LessThan is the < operator
	LessThan

	// GreaterThan is the > operator
	GreaterThan

	// LessThanOrEqual is the <= operator
	LessThanOrEqual

	// GreaterThanOrEqual is the >= operator
	GreaterThanOrEqual

	// Not is the ! operator
	Not

	// Or is the || operator
	Or

	// And is the && operator
	And

	// Xor is the ^^ operator
	Xor

	// Nand is the !& operator
	Nand

	// Nor is the !| operator
	Nor

	// Xnor is the !^ operator
	Xnor
)

// Condition represents an expression that always evaluates to a boolean value:
// [ expression ] [ condition_op ] [ expression ]
type Condition struct {
	Token Token
	Type  ConditionType
	Left  *Expression
	Right *Expression
	Value bool
}

func (c *Condition) expressionNode() {}

// TokenLiteral returns the literal value of the token
func (c *Condition) TokenLiteral() string { return c.Token.Literal }
