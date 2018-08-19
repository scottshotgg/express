package parse

import (
	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

// TODO: add in * and / and <
func (p *Parser) EvaluateBinaryOperation(left, right, op token.Value) (opToken token.Value, err error) {
	//fmt.Println("EvaluateBinaryOperation")

	switch op.Type {
	case "add":
		opToken, err = p.AddOperands(left, right)
		if err != nil {
			err = errors.New("Error adding operands")
		}

	case "sub":
		opToken, err = p.SubOperands(left, right)
		if err != nil {
			err = errors.New("Error subtracting operands")
		}

	case "mult":
		opToken, err = p.MultOperands(left, right)
		if err != nil {
			err = errors.New("Error multiplying operands")
		}

	case "div":
		opToken, err = p.DivOperands(left, right)
		if err != nil {
			err = errors.New("Error dividing operands")
		}

	case "lthan":
		//fmt.Println("lthan")
		opToken, err = p.LessThanOperands(left, right)
		if err != nil {
			err = errors.New("Error evaluating boolean expression")
			return
		}

	default:
		err = errors.Errorf("Undefined operator; left: %+v right: %+v op: %+v", left.True, right.True, op.String)
		return
		//fmt.Println(err.Error())
	}

	// opToken.Name = op.Type + "Op"
	// opToken.Type = "OP"
	// opToken.OpMap = opMap
	// opToken.True = opMap["eval"].(token.Value)
	// opToken.String = left.String + op.String + right.String

	opToken.Metadata = map[string]interface{}{
		"eval":  opToken.True,
		"type":  token.BoolType,
		"left":  left,
		"op":    op,
		"right": right,
		// "string": left.String + op.String + right.Stri fng,
	}
	// if opToken.Type == token.IntType {
	// 	opToken.String = strconv.Itoa(opToken.True.(int))
	// }

	leftString := left.Name
	if leftString == "" {
		leftString = left.String
	}

	rightString := right.Name
	if rightString == "" {
		rightString = right.String
	}

	opToken.String = leftString + op.String + rightString

	return
}

// EvaluateUnaryOperation ...
// TODO: implement this stuff
func (p *Parser) EvaluateUnaryOperation(left, op token.Value) { // (token.Value, error) {
}
