package parse

import (
	"os"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/token"
	"go.uber.org/zap"
)

const (
	// FIXME: see if we can change this in the binary as a compile time flag
	expressDebug = "EXPR_DEBUG"
)

var (
	ErrCouldNotParseType = errors.New("Could not parse type")
	logger               *zap.Logger
	sugar                *zap.SugaredLogger
	err                  error
)

type VariableType int

const (
	UNRECOGNIZED VariableType = iota
	INT
	FLOAT
	BOOL
	STRING
	VAR // I think we should prevent var from being a multiple things
	ARRAY
	GROUP
	OBJECT // Use this for a hashmap
	POINTER
	FUNCTION
	SET
)

type AccessType int

const (
	NOTSET AccessType = iota
	PRIVATE
	PUBLIC
)

type Token struct {
	TokenClass string
	Type       string
	Lexemes    []lex.Lexeme
	Value      map[string]interface{}
}

type Parser struct {
	source []token.Token
	length int
	Index  int
	// CurrentLexeme lex.Lexeme
	// CurrentToken     Token
	meta             *Meta
	Output           []Token
	IgnoreWhiteSpace bool

	// LastLexeme    lex.Lexeme
	// CurrentLexeme lex.Lexeme
	// NextLexeme    lex.Lexeme

	ProcessedTokens []token.Token
	LastToken       token.Token
	CurrentToken    token.Token
	NextToken       token.Token

	States []Parser
}

// type Variable struct {
// 	name         string
// 	value        interface{}
// 	variableType VariableType
// }

func init() {
	// FIXME: for now just check 'true' for now
	// If debug is not turned on, instantiate a logger that ignore those
	if os.Getenv(expressDebug) == "true" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	// Use a sugared logger; slower but has print/f/ln which makes it more versatile and readable
	sugar = logger.Sugar()
}

func NewVariable(name string, value interface{}, variableType VariableType) *Variable {
	return &Variable{
		Name:       name,
		Type:       variableType,
		ActingType: variableType,
		Value:      value,
		AccessType: PRIVATE,
		Metadata:   map[string]interface{}{},
	}
}

func NewVariableFromTokenValue(tv token.Value) *Variable {
	trueValue := tv.True
	return &Variable{
		Name:       tv.Name,
		Type:       variableTypeFromString(tv.Type),
		ActingType: variableTypeFromString(tv.Type),
		Value:      trueValue,
		AccessType: accessTypeFromString(tv.AccessType),
		Metadata:   tv.Metadata,
	}
}

func (v *Variable) SetAccessType(accessType AccessType) {
	v.AccessType = accessType
}

// type TokenMetadata struct {
// 	Type string
// 	Data interface{}
// }

type Variable struct {
	Name       string
	Type       VariableType
	ActingType VariableType
	Value      interface{}
	AccessType AccessType
	Metadata   map[string]interface{}
}

// func (v *Variable) String() {

// }

// func (v *Variable) Type() {

// }

type Scope map[string]*Variable

type Meta struct {
	// global       Scope
	currentScope Scope
	scopes       *Stack

	currentVariable *Variable
	variableStack   *Stack
}

func NewMeta() *Meta {
	globalScope := Scope{
		"test": NewVariable("test", "test", STRING),
	}

	return &Meta{
		currentScope: globalScope,
		scopes:       NewStack(),
		currentVariable: &Variable{
			Metadata: map[string]interface{}{},
		},
		variableStack: NewStack(),
	}
}

func (m *Meta) DeclareVariable() error {
	if m.currentVariable.Type == UNRECOGNIZED ||
		// Commented this; don't really care if it's set for normal
		// variables and it was messing up arrays
		// m.currentVariable.ActingType != UNRECOGNIZED ||
		m.currentVariable.AccessType == NOTSET ||
		m.currentVariable.Name == "" ||
		m.currentVariable.Value == nil {
		return errors.Errorf("Variable does not contain all required fields: %+v", m.currentVariable)
	}

	if m.currentVariable.Type == SET {
		return errors.Errorf("Variable cannot be declared with type SET")
	}

	if m.currentVariable.Type == VAR && m.currentVariable.ActingType == UNRECOGNIZED {
		return errors.Errorf("Variable of type VAR cannot be declared with no acting type: %+v", m.currentVariable)
	}

	// TODO: check all the required fields and matching types, etc
	m.currentScope[m.currentVariable.Name] = m.currentVariable
	////fmt.Println("metadata3", m.currentVariable)
	m.currentVariable = &Variable{
		Metadata: map[string]interface{}{},
	}

	return nil
}

func (m *Meta) DeclareVariableFromTokenValue(tv token.Value) error {
	variable := NewVariableFromTokenValue(tv)

	//fmt.Println("variable::", variable)

	if variable.Type == UNRECOGNIZED {
		return errors.Errorf("Variable type still unrecognized: %+v", tv)
	}
	// Commented this; don't really care if it's set for normal
	// variables and it was messing up arrays
	// tv.ActingType != UNRECOGNIZED ||
	if variable.AccessType == NOTSET {
		return errors.Errorf("Variable access type not set: %+v", tv)
	}
	if tv.Name == "" {
		return errors.Errorf("Variable does not have a name: %+v", tv)
	}
	if tv.True == nil {
		return errors.Errorf("Variable does not have any value: %+v", tv)
	}

	if variable.Type == SET {
		return errors.Errorf("Variable cannot be declared with type SET")
	}

	if variable.Type == VAR && variable.ActingType == UNRECOGNIZED {
		return errors.Errorf("Variable of type VAR cannot be declared with no acting type: %+v", tv)
	}

	// TODO: check all the required fields and matching types, etc
	m.currentScope[tv.Name] = variable
	//fmt.Println("metadata3", variable)

	return nil
}

func (m *Meta) NewVariable() *Variable {
	m.variableStack.Push(m.currentVariable)
	m.currentVariable = &Variable{
		Metadata: map[string]interface{}{},
	}
	return m.currentVariable
}

func (m *Meta) Height() int {
	return m.scopes.Length() + 1
}

func (m *Meta) GetVariable(variableName string) (*Variable, bool) {
	// Might have problems with the pointer here
	if variable, ok := m.currentScope[variableName]; ok {
		return variable, true
	}

	currentScope := m.currentScope
	defer func(m *Meta, current Scope) {
		m.currentScope = current
	}(m, currentScope)

	pop, err := m.ExitScope()
	defer m.scopes.Push(pop)
	if err != nil {
		return nil, false
	}

	return m.GetVariable(variableName)
}

func (m *Meta) NewScope() {
	m.scopes.Push(m.currentScope)
	m.currentScope = Scope{}
}

func (m *Meta) NewScopeFromScope(scopeToInherit Scope) {
	m.scopes.Push(m.currentScope)

	// Copy all from scopeToInherit to new scope
	m.currentScope = Scope{}
	for k, v := range scopeToInherit {
		m.currentScope[k] = v
	}
}

func (m *Meta) NewInheritedScope() {
	// Just push the current scope there, leave all the vars in the scope
	// since Variable is a pointer, it should set it if we set it in a
	// lower scope

	// Copy all from current scope to new scope
	newScope := Scope{}
	for k, v := range m.currentScope {
		newScope[k] = v
	}

	// Push new scope
	m.scopes.Push(newScope)
}

func (m *Meta) ExitScope() (Scope, error) {
	scope, err := m.scopes.Pop()
	if err != nil {
		// TODO:
		return Scope{}, err
	}

	m.currentScope = scope.(Scope)
	return m.currentScope, nil
}

// func (p *Parser) PeekNext() lex.Lexeme {
// 	//fmt.Println("peeking")
// 	if p.Index < p.length {
// 		current := p.source[p.Index+1]
// 		for p.IgnoreWhiteSpace && current.Type == "space" {
// 			//fmt.Println("in the loop")
// 			return p.PeekNext()
// 		}
// 		return current
// 	}
// 	return lex.Lexeme{}
// }

// func (p *Parser) PeekLast() lex.Lexeme {
// 	if p.Index > 0 {
// 		return p.source[p.Index-1]
// 	}
// 	return lex.Lexeme{}
// }

func (p *Parser) Length() int {
	return p.length
}

func New(tokens []token.Token) *Parser {
	p := &Parser{
		// source:           append(append([]token.Token{token.TokenMap["{"]}, tokens...), token.TokenMap["}"]),
		source:           tokens,
		length:           len(tokens),
		meta:             NewMeta(),
		IgnoreWhiteSpace: true,
	}
	p.Shift()
	return p
}

// // ParseDeclaration ...
// func (p *Parser) ParseDeclaration() {
// 	current := p.Shift()
// 	next := p.PeekNext()

// 	switch next.Type {
// 		case ""
// 	}
// }

func (p *Parser) Parse() (token.Value, error) {
	syntacticTokens, err := p.Syntactic()
	if err != nil {
		//fmt.Println("error in syntactic parsing", err)
		os.Exit(9)
	}

	pNew := New(syntacticTokens)
	semanticToken, err := pNew.Semantic()
	if err != nil {
		//fmt.Println("error in semantic parsing", err)
		os.Exit(9)
	}

	return semanticToken, nil
}
