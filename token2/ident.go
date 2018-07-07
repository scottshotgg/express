package token2

type IdentifierToken struct {
	*LiteralToken
	name           string
	accessModifier AccessModifierType
}

func (i *IdentifierToken) GetName() string {
	return i.name
}

func (i *IdentifierToken) GetAccessModifier() AccessModifierType {
	return i.accessModifier
}

func NewIdentFromInt() *IdentifierToken {
	return &IdentifierToken{
		LiteralToken: NewInt(),
	}
}

// func NewIdentFromBool() {
// 	i := NewBool()
// }

// func NewIdentFromFloat() {
// 	i := NewFloat()
// }

// func NewIdentFromString() {
// 	i := NewString()
// }
