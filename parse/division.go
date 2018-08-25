package parse

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

// DivOperands returns the addition of two operands based on their type
func (p *Parser) DivOperands(left, right token.Value) (token.Value, error) {
	// fmt.Println("DivOperands")
	var valueToken token.Value
	leftType := left.Type
	rightType := right.Type
	// fmt.Println("firsttime", left, right, leftType, rightType)

	if leftType == rightType {
		valueToken.Type = leftType

		switch leftType {
		case token.IntType:
			// TODO: could make it round by doing this
			// Might need to do something like this depending
			// on the type that is is going into
			// valueToken.True = int(math.Round(float64(left.True.(int)) / float64(right.True.(int))))
			valueToken.True = left.True.(int) / right.True.(int)
			valueToken.String = strconv.Itoa(valueToken.True.(int))
			fmt.Println("valueToken", valueToken)

		// case token.StringType:
		// 	valueToken.True = left.True.(string) / right.True.(string)
		// 	valueToken.String = valueToken.True.(string)

		case token.FloatType:
			valueToken.True = left.True.(float64) / right.True.(float64)
			// TODO: need to count the decimal place if we start using this
			valueToken.String = strconv.FormatFloat(valueToken.True.(float64), 'f', 5, 64)

		case token.BoolType:
			valueToken.True = !(left.True.(bool) && right.True.(bool))
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
			valueToken, err = p.DivOperands(left, right)
			if err != nil {
				fmt.Println("ERROR", err)
			}

		case token.ObjectType:
			result := right.True.(map[string]token.Value)
			// if ok := left.True.(map[string]token.Value)

			for key, value1 := range left.True.(map[string]token.Value) {
				if value2, ok := result[key]; ok {
					resultValue, err := p.DivOperands(value1, value2)
					// resultValue.AccessType = value1.AccessType
					// TODO: for some reason we couldnt access the `.True` of the map result
					// TODO: this means we could not add the operands, do something here later on: ideally we shouldnt get this
					if err != nil {
						return token.Value{}, errors.Wrap(err, "p.DivOperands")
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
			fmt.Println("Type not declared for DivOperands", left, right, leftType, rightType)

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
