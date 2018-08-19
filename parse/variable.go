package parse

import "github.com/scottshotgg/express/token"

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
	STRUCT
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

type Variable struct {
	Name       string
	Type       VariableType
	ActingType VariableType
	Value      interface{}
	AccessType AccessType
	Metadata   map[string]interface{}
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

func (v *Variable) SetAccessType(at AccessType) { v.AccessType = at }
