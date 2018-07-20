package parse

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/ExpressRedo/token"
)

// FIXME: move this to its own file
func (p *Parser) LessThanOperands(left, right token.Value) (token.Value, error) {
	// FIXME: this only works for ints right now
	// Need to put a type on this
	return token.Value{
		True:   left.True.(int) < right.True.(int),
		String: strconv.FormatBool(left.True.(int) < right.True.(int)),
	}, nil
}

// TODO: add in * and / and <
func (p *Parser) EvaluateBinaryOperation(left, right, op token.Value) (opToken token.Value, err error) {
	fmt.Println("EvaluateBinaryOperation")

	switch op.Type {
	// case "add":
	// 	opToken, err = p.AddOperands(left, right)
	// 	if err != nil {
	// 		err = errors.New("Error adding operands")
	// 	}

	// case "sub":
	// 	opToken, err = p.SubOperands(left, right)
	// 	if err != nil {
	// 		err = errors.New("Error subtracting operands")
	// 	}

	// case "mult":
	// 	opToken, err = p.MultOperands(left, right)
	// 	if err != nil {
	// 		err = errors.New("Error multiplying operands")
	// 	}

	// case "div":
	// 	opToken, err = p.DivOperands(left, right)
	// 	if err != nil {
	// 		err = errors.New("Error dividing operands")
	// 	}

	case "lthan":
		fmt.Println("lthan")
		opToken, err = p.LessThanOperands(left, right)
		if err != nil {
			err = errors.New("Error evaluating boolean expression")
		}

	default:
		err = errors.Errorf("Undefined operator; left: %+v right: %+v op: %+v", left, right, op)
		fmt.Println(err.Error())
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
		// "string": left.String + op.String + right.String,
	}
	if opToken.Type == token.IntType {
		opToken.String = strconv.Itoa(opToken.True.(int))
	}
	return
}

// // // EvaluateUnaryOperation ...
// // // TODO: implement this stuff
// // func (p *Parser) EvaluateUnaryOperation(left, op token.Value) { // (token.Value, error) {
// // }

func (p *Parser) GetFactor() (token.Value, error) {
	fmt.Println("GetFactor")
	next := p.NextToken
	fmt.Printf("next %+v\n", next)

	var value token.Value
	var err error

	switch p.NextToken.Type {
	case token.Literal:
		p.Shift()
		fmt.Println("after p.GetFactor NextToken", p.NextToken)
		value = p.CurrentToken.Value
		// FIXME: holy fuck haxorz
		if value.Type == token.IntType {
			value.String = strconv.Itoa(value.True.(int))
		}
		fmt.Println("hey its me the value", value)

	case token.Ident:
		p.Shift()
		// var ok bool
		// if value, ok = p.DeclarationMap[p.CurrentToken.Value.String]; !ok {
		// 	if p.LastMeta != nil {
		// 		fmt.Println(p.DeclarationMap)
		// 		if value, ok = (*p.LastMeta).DeclarationMap[p.CurrentToken.Value.String]; !ok {
		// 			// FIXME: holy fuck haxorz
		// 			if value.Type == token.IntType {
		// 				fmt.Printf("fuckthisshit2 %+v\n", p.CurrentToken)
		// 				value.String = next.Value.String
		// 			}
		// 			fmt.Println((*p.LastMeta).DeclarationMap)
		// 			return token.Value{}, errors.New("Undefined variable reference")
		// 		}
		// 	}
		// 	// // FIXME: holy fuck haxorz
		// 	// if value.Type == token.IntType {
		// 	// 	fmt.Printf("fuckthisshit %+v\n", p.CurrentToken)
		// 	// 	value.String = next.Value.String
		// 	// }
		// }

		fmt.Println("holy shit gettin that var")
		variable, ok := p.meta.GetVariable(p.CurrentToken.Value.String)
		fmt.Println(variable, ok)
		if !ok {
			return token.Value{}, errors.New("Undefined variable reference " + p.CurrentToken.Value.String)
		}

		// return p.GetExpression()
		value, err = mapVariableToTokenValue(variable), nil
		if err != nil {
			fmt.Println("Could not map variable to token value")
			os.Exit(9)
		}
		value.Metadata["refs"] = p.CurrentToken.Value.String

	// case token.Group:
	// 	meta := Meta{
	// 		AppendDeclarations: true,
	// 		IgnoreWS:           true,
	// 		Tokens:             p.NextToken.Value.True.([]token.Token),
	// 		Length:             len(p.NextToken.Value.True.([]token.Token)),
	// 		CheckOptmization:   true,
	// 		LastMeta:           m,
	// 		DeclarationMap:     map[string]token.Value{},
	// 	}
	// 	meta.Shift()
	// 	// Might have to change this to GetExpression
	// 	value, err = meta.GetExpression()
	// 	if err != nil {
	// 		return token.Value{}, err
	// 	}
	// 	// FIXME: holy fuck haxorz
	// 	if value.Type == token.IntType {
	// 		value.String = strconv.Itoa(value.True.(int))
	// 	}
	// 	p.Shift()
	// 	// os.Exit(9)

	// // case "":
	// // 	fmt.Println("we at the end?")
	// // 	os.Exit(8)

	case token.Array:
		fmt.Println("ayy rayy")
		arrayContents, ok := next.Value.True.([]token.Token)
		if !ok {
			fmt.Println("wtf no arrray stuffs", next)
			os.Exit(9)
		}
		fmt.Println("current", p.CurrentToken)
		fmt.Println("next", next.Value)

		var arrayContentsExpressions []token.Token
		for _, piece := range arrayContents {
			fmt.Println("expression1", piece)
			expr, err := p.ParseExpression(piece)
			if err != nil {
				return token.Value{}, err
			}
			arrayContentsExpressions = append(arrayContentsExpressions, token.Token{
				ID:    0,
				Type:  token.Literal,
				Value: expr,
			})
		}
		fmt.Println("arrayContents", arrayContents)
		fmt.Println("arrayContentsExpressions", arrayContentsExpressions)

		p.meta.currentVariable.Metadata = map[string]interface{}{
			"length": len(arrayContents),
			"vector": false,
		}
		// p.meta.currentVariable.Metadata["vector"] = false
		// fmt.Println("metadata", p.meta.currentVariable.Metadata["length"])
		// arrayType := VariableTypeString(p.meta.currentVariable.ActingType)
		arrayType := arrayContentsExpressions[0].Value.Type
		if len(arrayContents) > 0 {
			fmt.Println("arrayType", arrayType)
			for _, arrayValue := range arrayContentsExpressions {
				// fmt.Println("arrayType", arrayType, arrayValue.Value.Type)
				fmt.Println("arrayType.Value.Type", arrayValue.Value.Type)
				if arrayValue.Value.Type != arrayType {
					arrayType = ""
					break
				}
			}
		}

		// if p.meta.currentVariable.Type == SET {
		// 	p.meta.currentVariable.Type = ARRAY
		// 	p.meta.currentVariable.ActingType =
		// }

		if arrayType == "" {
			// TODO: we need to figure out what to do here ...
			// detemine if we can coerce this to the desired type ...
			return token.Value{}, errors.New("Error: array elements are not of static type")
		} else if p.meta.currentVariable.Type == SET {
			// If the arrayType is not empty and we have a SET type then
			// set the type to ARRAY and the ActingType to the arrayType
			p.meta.currentVariable.Type = ARRAY
			p.meta.currentVariable.ActingType = variableTypeFromString(arrayType)
		} else if arrayType != VariableTypeString(p.meta.currentVariable.ActingType) {
			fmt.Println("hi its me the variable type", arrayType, VariableTypeString(p.meta.currentVariable.ActingType))
			return token.Value{}, errors.New("Error: array elements are of different type than type declaration")
		}
		// actingType := VariableTypeString(p.meta.currentVariable.ActingType)
		// if arrayValue.Value.Type != arrayType {
		// 	fmt.Println("NOT EQUAL", arrayValue.Value.Type, arrayType)
		// }

		p.meta.currentVariable.Value = arrayContentsExpressions

		fmt.Println("p.meta.currentVariable", p.meta.currentVariable)
		value = mapVariableToTokenValue(p.meta.currentVariable)
		fmt.Printf("next: %+v\n", p.meta.currentVariable)
		fmt.Printf("value: %+v\n", value)
		fmt.Println("value", value)
		p.Shift()

	default:
		fmt.Println("last2", p.LastToken)
		fmt.Println("current2", p.CurrentToken)
		fmt.Println("next2", p.NextToken)
		return token.Value{}, errors.Errorf("default %+v", p.NextToken)
	}
	fmt.Println("value thing again", value)

	// FIXME: TODO: didn't wanna fix right now
	switch p.NextToken.Type {
	// case token.PriOp:
	// 	p.Shift()
	// 	op := p.CurrentToken
	// 	value2, verr := p.GetFactor()
	// 	if verr != nil {
	// 		return token.Value{}, verr
	// 	}
	// 	fmt.Println("value2thing", value2)

	// value, err = p.EvaluateBinaryOperation(value, value2, op.Value)
	// if err != nil {
	// 	return token.Value{}, err
	// }
	// // FIXME: holy fuck haxorz
	// if value.Type == token.IntType {
	// 	value.String = ""
	// }

	case token.Increment:
		value, err = p.AddOperands(value, token.Value{
			Type: token.IntType,
			True: 1,
		})
		if err != nil {
			return token.Value{}, err
		}
		// p.Shift()`
	}

	// FIXME: holy fuck haxorz
	// if value.Type == token.IntType {
	// 	value.String = next.Value.String
	// }
	fmt.Println("returning", value)
	return value, nil
}

// GetTerm ...
func (p *Parser) GetTerm() (token.Value, error) {
	fmt.Println("GetTerm")

	totalTerm, err := p.GetFactor()
	if err != nil {
		return token.Value{}, err
	}
	fmt.Println("totalTERM", totalTerm)

	for {
		switch p.NextToken.Type {
		// FIXME: TODO: didn't wanna fix right now
		// case token.SecOp:
		// 	p.Shift()
		// 	fmt.Println("woah i got a secop")
		// 	op := p.CurrentToken
		// 	factor2, ferr := p.GetFactor()
		// 	if ferr != nil {
		// 		return token.Value{}, ferr
		// 	}
		// 	fmt.Println("factor2", factor2)

		// 	totalTerm, err = p.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
		// 	if err != nil {
		// 		return token.Value{}, err
		// 	}
		// 	// FIXME: holy fuck haxorz
		// 	if totalTerp.Type == token.IntType {
		// 		totalTerp.String = strconv.Itoa(totalTerp.True.(int))
		// 	}

		// // TODO: need to fix this....
		case token.LThan:
			fmt.Println("in the lthan")
			// ident := p.LastToken
			nextTokenOpString := p.NextToken.Value.String
			p.Shift()
			op := p.CurrentToken
			factor2, ferr := p.GetTerm()
			if ferr != nil {
				return token.Value{}, ferr
			}
			fmt.Println("lthan totalTerm", totalTerm)
			fmt.Println("lthan factor2", factor2)
			totalTermEval, err := p.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
			if err != nil {
				return token.Value{}, err
			}
			// FIXME: holy fuck haxorz
			// if totalTerp.Type == token.IntType {
			// TODO: should use totalTerm.String here
			fmt.Printf("factor2before %+v\n", factor2)
			factor2.String = totalTerm.Name + nextTokenOpString + factor2.String
			// }
			fmt.Printf("totalTerm %+v\n", totalTerm)
			fmt.Printf("totalTermEval %+v\n", totalTermEval)
			fmt.Println("factor2", factor2)
			factor2.Metadata = totalTermEval.Metadata
			return factor2, nil

		case token.Separator:
			p.Shift()
			// // FIXME: holy fuck haxorz
			// if totalTerp.Type == token.IntType {
			// 	totalTerp.String = strconv.Itoa(totalTerp.True.(int))
			// }
			return totalTerm, nil

		default:
			// FIXME: holy fuck haxorz
			if totalTerm.Type == token.IntType {
				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
			}
			fmt.Println("i am here", p.NextToken)
			fmt.Println("totalTerm", totalTerm)
			return totalTerm, nil
		}
	}
}

func VariableTypeString(vt VariableType) (st string) {
	switch vt {
	case INT:
		st = "int"
	case FLOAT:
		st = "float"
	case BOOL:
		st = "bool"
	case STRING:
		st = "string"
	case VAR:
		st = "var"
	case SET:
		st = "set"
	case OBJECT:
		st = "object"
	case ARRAY:
		st = "array"
	// case STRINGA:
	// 	st = "string[]"

	default:
		st = ""
	}

	return
}

func AccessTypeString(at AccessType) (st string) {
	switch at {
	case PUBLIC:
		st = "public"
	case PRIVATE:
		st = "private"

	default:
		st = ""
	}

	return
}

func mapVariableToTokenValue(v *Variable) token.Value {
	md := map[string]interface{}{}
	for k, value := range v.Metadata {
		md[k] = value
	}

	fmt.Println("md", md)
	fmt.Println("metadata", v.Metadata)

	return token.Value{
		Name:   v.Name,
		Type:   VariableTypeString(v.Type),
		Acting: VariableTypeString(v.ActingType),
		True:   v.Value,
		// String: ,
		AccessType: AccessTypeString(v.AccessType),
		Metadata:   md,
	}
}

func (p *Parser) ParseExpression(tok token.Token) (token.Value, error) {
	// Kinda hacky, but it works so w/e for now; its 0300 and idc
	// but its always 0300 :^)
	pa := New([]token.Token{
		tok,
	})
	pa.meta.NewScopeFromScope(p.meta.currentScope)
	expression, err := pa.GetExpression()
	if err != nil {
		fmt.Println("Error: could not parse expression inside array")
		fmt.Println(err.Error())
		return token.Value{}, err
	}

	fmt.Println("insidexpression", expression)

	return expression, nil
}

// GetExpression ...
func (p *Parser) GetExpression() (token.Value, error) {
	fmt.Println("GetExpression")
	fmt.Printf("p.NextToken %+v\n", p.NextToken)

	switch p.NextToken.Type {
	// Will have to experiment on where to put this
	// Might need to put this in factor
	case token.Block:
		fmt.Println("found a block")
		block, err := p.CheckBlock()
		if err != nil {
			fmt.Println("waddup blockboi", err)
		}
		return block, nil

	// Assignment Expression
	case token.Assign:
		// FIXME: I think this should go in the token.Ident case of GetStatement
		// p.DeclaredName = p.CurrentToken.Value.String
		// p.DeclaredAccessType = p.CurrentToken.Value.Type
		p.meta.currentVariable.Name = p.CurrentToken.Value.String
		p.meta.currentVariable.AccessType = accessTypeFromString(p.CurrentToken.Value.Type)
		switch p.NextToken.Value.Type {
		case "init":
			if p.meta.currentVariable.Type != UNRECOGNIZED {
				return token.Value{}, errors.New("Type specification with init is not valid")
			}
			p.meta.currentVariable.Type = SET
			fallthrough

		case "assign":
			p.Shift()
			fmt.Println("shifted", p.NextToken)
			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}
			fmt.Printf("expr in assign %+v\n", expr)

			if p.meta.currentVariable.Type == UNRECOGNIZED {
				if variable, ok := p.meta.GetVariable(p.NextToken.Value.String); ok {
					p.meta.currentVariable.Type = variable.Type
				} else {
					fmt.Println("unable to find variable")
					fmt.Println(p.meta.currentVariable)
					fmt.Println("scope", p.meta.currentScope)
					os.Exit(9)
				}
			} else if p.meta.currentVariable.Type == SET {
				if variable, ok := p.meta.GetVariable(p.meta.currentVariable.Name); ok {
					fmt.Println("variable", variable, "ok", ok)
				}

				p.meta.currentVariable.Type = variableTypeFromString(expr.Type)
			} else if p.meta.currentVariable.Type != variableTypeFromString(expr.Type) {
				fmt.Println(VariableTypeString(p.meta.currentVariable.Type), expr.Type)
				// TODO: implicit type casting here
				return token.Value{}, errors.New("No implicit type casting as of now")
			}
			p.meta.currentVariable.Value = expr.True
			if ref, ok := expr.Metadata["refs"]; ok {
				fmt.Println("there was a ref")
				p.meta.currentVariable.Metadata["refs"] = ref
			}
			fmt.Println("p.meta.currentVariable2", p.meta.currentVariable)

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				// TODO:
				fmt.Println("declareVariable error", err.Error())
				os.Exit(9)
			}

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}
			return token.Value{}, errors.New("Could not find variable: " + currentName)

		case "set":
			if p.meta.currentVariable.Type != UNRECOGNIZED {
				return token.Value{}, errors.New("Type specification with set is not valid")
			}
			p.meta.currentVariable.Type = SET

			if variable, ok := p.meta.currentScope[p.meta.currentVariable.Name]; ok {
				fmt.Println("Error: Variable already declared in this scope", variable)
				return token.Value{}, errors.New("Error: Variable already declared in this scope")
			}

			p.Shift()
			fmt.Println("what do", p.NextToken)
			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}
			fmt.Printf("expr in set %+v\n", expr)

			p.meta.currentVariable.Type = variableTypeFromString(expr.Type)
			p.meta.currentVariable.Value = expr.True

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				// TODO:
				fmt.Println("declareVariable error", err.Error())
				os.Exit(9)
			}

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}
			return token.Value{}, errors.New("Could not find variable: " + currentName)
		}

	// case token.LThan:
	// 	fmt.Println("wtf")
	// 	fmt.Println("current", p.CurrentToken)
	// 	fmt.Println("next", p.NextToken)
	// 	p.Shift()
	// 	term, err := p.GetTerm()
	// 	if err != nil {
	// 		return token.Value{}, err
	// 	}
	// 	return term, nil

	case token.Increment:
		fmt.Println("woah increment brah")
		variable, ok := p.meta.GetVariable(p.meta.currentVariable.Name)
		if !ok {
			fmt.Println("COuld not find variable:", p.meta.currentVariable.Name)
			return token.Value{}, errors.New("shit")
		}
		value, err := p.AddOperands(mapVariableToTokenValue(variable), token.Value{
			Type: token.IntType,
			True: 1,
		})
		if err != nil {
			return token.Value{}, err
		}
		fmt.Println("token.Increment", value)

	default:
		return p.GetTerm()
	}

	return token.Value{}, errors.Errorf("default %+v", p.NextToken)
}

func (p *Parser) GetKeyword() (token.Value, error) {
	fmt.Println("LAST666", p.LastToken)
	fmt.Println("CURRENT666", p.CurrentToken)
	fmt.Println("NEXT666", p.NextToken)
	switch p.NextToken.Value.String {
	case "for":
		// Make a new meta
		// Get a statement
		// Get an expression
		// Get another expression
		// Sub operands to find step
		p.Shift()
		fmt.Println("GETTING LOOP")
		p.meta.NewInheritedScope()
		stmt, err := p.GetStatement()
		if err != nil {
			fmt.Println("Error: Could not get statement")
			os.Exit(9)
			return token.Value{}, err
		}

		expr, err := p.GetExpression()
		if err != nil {
			fmt.Println("Error: Could not get expression")
			os.Exit(9)
			return token.Value{}, err
		}

		expr2, err := p.GetExpression()
		if err != nil {
			fmt.Println("Error: Could not get expression")
			os.Exit(9)
			return token.Value{}, err
		}
		fmt.Println("stmt", stmt)
		fmt.Println("expr1", expr)
		fmt.Println("expr2", expr2)
		// os.Exit(9)

		step, err := p.SubOperands(expr2, stmt)
		if err != nil {
			fmt.Println("Could not sub operands")
			return token.Value{}, err
		}

		p.Shift()
		body, err := p.CheckBlock()
		if err != nil {
			fmt.Println("Could not check block")
			return token.Value{}, err
		}

		// TODO: don't know if we need to do this
		md := map[string]interface{}{
			"start": stmt,
			"end":   expr,
			"step":  step,
		}
		for k, v := range expr.Metadata {
			md[k] = v
		}
		expr.Metadata = map[string]interface{}{}

		return token.Value{
			Type:     token.For,
			True:     body.True.([]token.Value),
			String:   expr.String,
			Metadata: md,
		}, nil

		os.Exit(9)

	case "if":
		p.Shift()
		expr, err := p.GetExpression()
		if err != nil {
			fmt.Println("Error: Could not get expresssion")
			return token.Value{}, err
		}

		block, err := p.CheckBlock()
		if err != nil {
			fmt.Println("Error: Could not get block")
			return token.Value{}, err
		}

		fmt.Printf("expr %+v\n", expr)
		fmt.Printf("block %+v\n", block)
		fmt.Printf("next555 %+v\n", p.NextToken)

		expr.Metadata["check"] = expr.String

		return token.Value{
			Type:     token.If,
			True:     block.True.([]token.Value),
			String:   expr.String,
			Metadata: expr.Metadata,
		}, nil

	default:
		fmt.Println("woah idk", p.CurrentToken)
		os.Exit(9)
	}

	return token.Value{}, nil
}

func variableTypeFromString(vtString string) (vt VariableType) {
	switch vtString {
	case "int":
		vt = INT
	case "bool":
		vt = BOOL
	case "string":
		vt = STRING
	case "float":
		vt = FLOAT
	case "BLOCK":
		vt = OBJECT
	case "set":
		vt = SET
	case "array":
		vt = ARRAY
	}

	return
}

func accessTypeFromString(atString string) (at AccessType) {
	switch atString {
	case "public":
		at = PUBLIC
	case "private":
		at = PRIVATE
	}

	return
}

// GetStatement ...
func (p *Parser) GetStatement() (token.Value, error) {
	fmt.Println("GetStatement")
	fmt.Println("p.NextToken", p.NextToken)
	// p.Shift()
	switch p.NextToken.Type {
	case token.Type:
		p.meta.currentVariable.Type = variableTypeFromString(p.NextToken.Value.Type)
		p.meta.currentVariable.ActingType = variableTypeFromString(p.NextToken.Value.Acting)
		p.Shift()
		// TODO: could either recurse here, or fallthrough
		if p.NextToken.Type != token.Ident {
			// FIXME: what to do here
			break
		}
		fallthrough

	// // TODO: will have to consider declarations too
	case token.Ident:
		fmt.Println("ident", p.NextToken)
		fmt.Println("declaredMap", p.meta.currentScope)
		if p.meta.currentVariable.Type == UNRECOGNIZED {
			// TODO: maybe we should just load the entire variable at this point
			if variable, ok := p.meta.GetVariable(p.NextToken.Value.String); ok {
				variable.Metadata["assign"] = true
				fmt.Println("FOUND THE VAR", p.NextToken.Value.String)
				p.meta.currentVariable.Type = variable.Type
				p.meta.currentVariable.Metadata = variable.Metadata
			} else {
				// fmt.Println("ASSIGNMENT DECLARED VALUE", m.DeclaredValue)
				p.Shift()
				expr, err := p.GetExpression()
				fmt.Printf("THIS IS THE EXPRESSION %+v %s\n", expr, err)
				return expr, err
			}
		}
		fmt.Println("ASSIGNMENT DECLARED VALUE", p.meta.currentVariable.Type)
		p.Shift()
		fmt.Println(p.NextToken)
		tv, err := p.GetExpression()
		fmt.Println("nofind THIS IS THE EXPRESSION", tv, err)
		if err != nil {
			fmt.Println("getExpressionErr", err)
			os.Exit(9)
		}
		fmt.Println("another", p.NextToken)
		return tv, nil

	// FIXME: TODO: didn't wanna fix right now
	case token.Keyword:
		keyword, err := p.GetKeyword()
		if err != nil {
			return token.Value{}, err
		}
		// p.Shift()
		return keyword, nil

	case token.Separator:
		// fmt.Println("should we have gotten this here?")
		// os.Exit(9)
		// TODO: not sure if we should return something else here
		p.Shift()
		// return token.Value{}, nil

	case token.SecOp:
		switch p.CurrentToken.Value.Type {
		case "sub":
			// TODO: need to do something here for negative expression

		default:
			return token.Value{}, errors.New("Unrecognized position for operator")
		}

	case token.Block:
		fmt.Println("blockboi")
		p.Shift()
		block, err := p.CheckBlock()
		if err != nil {
			fmt.Println("waddup blockboi", err)
			os.Exit(9)
		}
		p.Shift()
		return block, err

	default:
		// TODO: this causes infinite loops when you cant parse
		fmt.Println("hey its me the default", p.NextToken)
		os.Exit(9)
	}

	return token.Value{}, nil
}

// CheckBlock ...
func (p *Parser) CheckBlock() (token.Value, error) {
	fmt.Println("CheckBlock")

	p.Shift()
	// Open up the block here
	tokensFromBlock, ok := p.CurrentToken.Value.True.([]token.Token)
	if !ok {
		fmt.Println("Error: Current token does not contain an array", p.CurrentToken)
		os.Exit(9)
	}

	// FIXME: TODO: we need to fix this hacky shit
	// works for now, but hacky as shit
	pNew := New(tokensFromBlock)
	pNew.meta.NewScopeFromScope(p.meta.currentScope)

	blockTokens := []token.Value{}

	for {
		fmt.Println("pNew.NextToken", pNew.NextToken)
		stmt, err := pNew.GetStatement()
		if err != nil {
			fmt.Println()
			fmt.Println("err", err)
			os.Exit(9)
		}

		// This is by-passing the blank "{}" token that is
		// produced from the comma somtimes; need to solve
		// it more elegantly
		if reflect.DeepEqual(stmt, token.Value{}) {
			return token.Value{
				Type: token.Block,
				True: blockTokens,
			}, nil
		}

		blockTokens = append(blockTokens, stmt)

		// p.meta.NewVariable()
		fmt.Println("CheckBlock currentScope: ", pNew.meta.currentScope)
		fmt.Println()

		if reflect.DeepEqual(pNew.NextToken, token.Token{}) {
			// fmt.Println(p.meta.GetVariable("a"))
			// p.meta.NewScopeFromScope(pNew.meta.currentScope)
			return token.Value{
				Type: token.Block,
				True: blockTokens,
			}, nil
		}
	}
}

// Semantic ...
func (p *Parser) Semantic() (token.Value, error) {
	fmt.Println("Semantic")

	block, err := p.CheckBlock()
	if err != nil {
		// TODO:
		return token.Value{}, err
	}
	fmt.Println("block", block)
	fmt.Println()

	fmt.Println("End currentScope: ", p.meta.currentScope)
	fmt.Println()

	return block, nil
}

// TODO: start here
// TODO: use next token
// TODO: start very simply with the definition in documentation/notes_about_shit
// TODO: VERY SIMPLE requirements parsing vars with the return architecture of semantic2
