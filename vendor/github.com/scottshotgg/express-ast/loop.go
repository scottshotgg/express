package ast

// LoopType encompasses all types of loops
type LoopType int

const (
	// StdFor is a standard for loop containing a condition and expression
	// operating around a declared variable:
	// `for` [ statement ] [ condition ] [ expression ] [ block ]
	StdFor LoopType = iota + 1

	// ForEver is the result of a blank StdFor loop:
	// `for` [ block ]
	ForEver

	// ForIn is a for loop that auto iterates over the keys of an iterable:
	// `for` [ literal ] `in` [ iterable ] [ block ]
	ForIn

	// ForOf is a for loop that auto iterates over the values of an iterable:
	// `for` [ literal ] `in` [ iterable ] [ block ]
	ForOf

	// ForOver is a for loop that auto iterates over the keys and values of an iterable:
	// `for` [ literal ] `,` [ literal ] `in` [ iterable ] [ block ]
	// `for` [ -object- ] `in` [ iterable ] [ block ]
	ForOver

	// While is a loop that operates only on a single required condition:
	// `while` [ condition ] [ block ]
	While

	// Until is a reverse-logic while loop:
	// `until` [ condition ] [ block ]
	Until
)

// Loop represents the following form:
// [ loop_type ] { iterable } [ block ]
type Loop struct {
	Token Token
	Type  LoopType
	Start int
	End   int
	Step  int
	Body  *Block
	Iter  *Iterable
	Temps map[string]*Ident
}

func (l *Loop) Kind() NodeType { return LoopNode }
