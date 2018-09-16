package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

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
	// logger.Debug("Declaring variable: " + m.currentVariable.Name)
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

func (m *Meta) VariableFromScope(scope Scope, variable *Variable) *Variable {
	var values []token.Value
	for _, v := range m.currentScope {
		values = append(values, mapVariableToTokenValue(v))
	}
	variable.Value = values
	fmt.Println("variable", variable)
	// m.DeclareVariableFromTokenValue(mapVariableToTokenValue(variable))
	// return nil
	return variable
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

func (m *Meta) Height() int { return m.scopes.Length() + 1 }

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

func (m *Meta) NewScopeFromVariable(variableToInherit *Variable) {
	m.scopes.Push(m.currentScope)

	// Copy all from variableToInherit to new scope
	m.currentScope = Scope{}
	// FIXME: I guess this function has to return an error
	vTokens, ok := variableToInherit.Value.([]token.Value)
	if !ok {
		// TODO: ...
	}

	for _, v := range vTokens {
		// m.currentScope[k] = v
		m.DeclareVariableFromTokenValue(v)
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
