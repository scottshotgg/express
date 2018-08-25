package parse

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

// AddOperands returns the addition of two operands based on their type
func (p *Parser) AddOperands(left, right token.Value) (token.Value, error) {
	// fmt.Println("AddOperands")
	var valueToken token.Value
	leftType := left.Type
	rightType := right.Type

	if leftType == rightType {
		valueToken.Type = leftType

		switch leftType {
		case token.IntType:
			valueToken.True = left.True.(int) + right.True.(int)
			valueToken.String = strconv.Itoa(valueToken.True.(int))

		case token.StringType:
			valueToken.True = left.True.(string) + right.True.(string)
			valueToken.String = valueToken.True.(string)

		case token.FloatType:
			valueToken.True = left.True.(float64) + right.True.(float64)
			// TODO: need to count the decimal place if we start using this
			valueToken.String = strconv.FormatFloat(valueToken.True.(float64), 'f', 5, 64)

		case token.BoolType:
			valueToken.True = left.True.(bool) || right.True.(bool)
			valueToken.String = strconv.FormatBool(valueToken.True.(bool))

		case token.CharType:
			// TODO: we will need to take into account the character encoding here and overflowing
			valueToken.True = string(rune(left.True.(string)[0]) + rune(right.True.(string)[0]))
			valueToken.String = valueToken.True.(string)

		// TODO: this will need some more thinking
		// case token.Byte:

		case token.VarType:
			left.Type = left.Acting
			right.Type = right.Acting

			var err error
			valueToken, err = p.AddOperands(left, right)
			if err != nil {
				return token.Value{}, errors.Wrap(err, "p.AddOperands")
			}

		case token.ObjectType:
			result := right.True.(map[string]token.Value)
			// if ok := left.True.(map[string]token.Value)

			for key, value1 := range left.True.(map[string]token.Value) {
				if value2, ok := result[key]; ok {
					resultValue, err := p.AddOperands(value1, value2)
					// resultValue.AccessType = value1.AccessType
					// TODO: for some reason we couldnt access the `.True` of the map result
					if err != nil {
						// TODO: this means we could not add the operands, do something here later on: ideally we shouldnt get this
						return token.Value{}, errors.Wrap(err, "p.AddOperands")
					}

					value2.True = resultValue.True
					result[key] = value2

				} else {
					result[key] = value1
				}
			}
			valueToken.True = result

		case token.ArrayType:
			valueToken.True = append(left.True.([]token.Value), right.True.([]token.Value)...)

		default:
			return token.Value{}, errors.Errorf("Type not declared for AddOperands %+v %+v %s %s", left, right, leftType, rightType)
		}

		return valueToken, nil
	}
	//  else {
	// 	if left side or right side is a string
	//		-> string
	//		-> object will be stringitized
	//	if left or right side is a float and the other is an int
	//		-> promote to float
	//	if left is int and right is bool
	//		-> int
	//	if right is bool and left is int
	//		-> bool
	// 	if left or right is float and other is bool
	//		-> float
	//	if left or right is object
	//		if other is string
	//			if there is no key with that name
	//				-> key and value named as string
	//			else
	//				-> undefined for now // FIXME: TODO:
	//		if other is IDENT
	//			-> ident name as key, ident value as value
	//		else
	//			-> undefined for now // FIXME: TODO:
	//	if left or right is array
	//		if other is SAME
	//			-> append to array
	//		if other is DIFFERENT
	//			-> dump into object? // FIXME: TODO:
	// }

	return token.Value{}, errors.New("Could not perform AddOperand on operands")
}
