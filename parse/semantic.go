package parse

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

var (
	functionOpType = ""
	inStruct       = false
	inObject       = false
)

func (p *Parser) GetFactor() (token.Value, error) {
	//fmt.Println("GetFactor")
	next := p.NextToken
	//fmt.Printf("next %+v\n", next)

	value := token.Value{
		Metadata: map[string]interface{}{},
	}

	switch p.NextToken.Type {
	case token.Literal:
		p.Shift()
		//fmt.Println("after p.GetFactor NextToken", p.NextToken)
		value = p.CurrentToken.Value

		// FIXME: holy fuck haxorz
		if value.Type == token.IntType {
			value.String = strconv.Itoa(value.True.(int))
		}
		//fmt.Println("hey its me the value", value)

	case token.Ident:
		p.Shift()
		// var ok bool
		// if value, ok = p.DeclarationMap[p.CurrentToken.Value.String]; !ok {
		// 	if p.LastMeta != nil {
		// 		//fmt.Println(p.DeclarationMap)
		// 		if value, ok = (*p.LastMeta).DeclarationMap[p.CurrentToken.Value.String]; !ok {
		// 			// FIXME: holy fuck haxorz
		// 			if value.Type == token.IntType {
		// 				//fmt.Printf("fuckthisshit2 %+v\n", p.CurrentToken)
		// 				value.String = next.Value.String
		// 			}
		// 			//fmt.Println((*p.LastMeta).DeclarationMap)
		// 			return token.Value{}, errors.New("Undefined variable reference")
		// 		}
		// 	}
		// 	// // FIXME: holy fuck haxorz
		// 	// if value.Type == token.IntType {
		// 	// 	//fmt.Printf("fuckthisshit %+v\n", p.CurrentToken)
		// 	// 	value.String = next.Value.String
		// 	// }
		// }

		//fmt.Println("wowzerz im gettin that var", p.meta.currentScope)
		variable, ok := p.meta.GetVariable(p.CurrentToken.Value.String)
		if !ok {
			// If we did not find it as a variable, look in the DefinedTypes map
			value2, ok := DefinedTypes[p.CurrentToken.Value.String]
			if !ok {
				return token.Value{}, errors.New("Undefined variable reference " + p.CurrentToken.Value.String)
			}

			variable = NewVariableFromTokenValue(value2)
		}

		refs := p.CurrentToken.Value.String
		if variable.Type == FUNCTION {
			// p.meta.currentVariable.Metadata["from_func"] = true
			functionOpType = "call"
			// FIXME: here we need to look up the return value and make sure it matches
			// FIXME: make sure the args match

			p.Shift()

			// This is a struct declaration
		} else if variable.Type == STRUCT {
			fmt.Println("i am here", variable)

			// Here we need to get the default values from the type map
			if p.NextToken.Type == token.Block {
				inStruct = true
				defer func(typename string) {
					inStruct = false
					value.Metadata["real"] = typename
				}(p.CurrentToken.Value.String)

				/*
					// // Need to make a deep copy function
					// // Extract the body of the struct
					// variableValue := variable.Value.([]token.Value)
					// valuers := []token.Value{}
					// for i := range variableValue {
					// 	if variableValue[i].Type == "struct" {
					// 		// Deep copy again here
					// 	}
					// 	fmt.Println("this is going on", variableValue[i])
					// 	valuers = append(valuers, variableValue[i])
					// }
					// variable.Value = valuers
				*/

				fmt.Println("p.meta.currentScope", p.meta.currentScope)
				for k, v := range p.meta.currentScope {
					fmt.Println("k, v", k, *v)
				}
				gob.Register([]token.Value{})

				// FIXME: check the errors
				vValue := []token.Value{}
				var buf bytes.Buffer
				enc := gob.NewEncoder(&buf)
				err = enc.Encode(variable.Value)
				if err != nil {
					return token.Value{}, err
				}

				dec := gob.NewDecoder(&buf)
				err = dec.Decode(&vValue)
				if err != nil {
					return token.Value{}, err
				}

				fmt.Println("reflecterooni", reflect.TypeOf(variable.Value))

				fmt.Println("variable.Value", variable.Value)
				fmt.Println("vValue", vValue)
				fmt.Println("reflecterooni2", reflect.TypeOf(vValue))
				variable.Value = vValue
				fmt.Println("reflecterooni3", reflect.TypeOf(variable.Value))

				thing := variable.Value.([]token.Value)
				// Deep copy did work, not surprising, but something is fucked
				fmt.Println("thing:", thing)
				// os.Exit(9)

				if len(p.NextToken.Value.True.([]token.Token)) > 0 {
					// fmt.Printf("variable %+v\n", variable)
					// fmt.Println("p.meta.currentScope", p.meta.currentVariable)
					// fmt.Println()

					// if len(p.NextToken.Value.True.([]token.Token)) > 0 {
					// 	os.Exit(9)
					// }

					// Need to make a new scope from the vars inside of the struct
					// pa := New(p.NextToken.Value.True.([]token.Token))
					pa := New([]token.Token{p.NextToken})
					// This gives scoping to the inside of the struct
					pa.meta.NewScopeFromScope(p.meta.currentScope)

					// Unpack all the values from the struct
					variable.Metadata["real"] = p.CurrentToken.Value.String

					block, err := pa.CheckBlock()
					if err != nil {
						return token.Value{}, err
					}

					// For some reason I couldn't get the `currentScope` variable to work here
					// valueBlock := []token.Value{}
					variableTokens, ok := variable.Value.([]token.Value)
					if !ok {
						return token.Value{}, errors.New("Could not assert value of struct during initialization")
					}

					for _, variableValue := range block.True.([]token.Value) {
						for i := 0; i < len(variableTokens); i++ {
							if variableTokens[i].Name == variableValue.Name {
								if variableTokens[i].Type != "var" && variableTokens[i].Type != "object" && variableTokens[i].Type != "struct" {
									if variableTokens[i].True != variableValue.True {
										variableValue.Metadata["default"] = false
									}
								}

								variableTokens[i] = variableValue
								break
							}
						}
					}

					fmt.Println("variable after", variable)

					// For now just shift over the block
					p.Shift()

					return mapVariableToTokenValue(variable), nil
				}

				inStruct = false
			} else if p.NextToken.Type == token.Accessor {
				fmt.Println("WOAH I MADE IT")
				os.Exit(9)
			}
			// FIXME: fix this
			// else {
			// 	return token.Value{}, errors.Errorf("Inaccurate assignment to type, need literal %s", p.NextToken.Type)
			// }

			p.Shift()
		}

		value = mapVariableToTokenValue(variable)
		value.Metadata["refs"] = refs

	case token.Array:
		arrayContents, ok := next.Value.True.([]token.Token)
		if !ok {
			return token.Value{}, errors.New("token.Array: next.Value.True.([]token.Token)")
		}

		var arrayContentsExpressions []token.Token
		for _, piece := range arrayContents {
			//fmt.Println("expression1", piece)
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
		//fmt.Println("arrayContents", arrayContents)
		//fmt.Println("arrayContentsExpressions", arrayContentsExpressions)

		p.meta.currentVariable.Metadata = map[string]interface{}{
			"length": len(arrayContents) - 1,
			"vector": false,
		}

		arrayType := arrayContentsExpressions[0].Value.Type
		if arrayType == "BLOCK" {
			arrayType = "object"
		}

		if len(arrayContents) > 0 {
			for i, arrayValue := range arrayContentsExpressions {

				if arrayValue.Value.Type == "BLOCK" {
					arrayValue.Value.Type = "object"
					arrayContentsExpressions[i].Value.Type = "object"
				}

				if arrayValue.Value.Type != arrayType {
					arrayType = ""
					break
				}
			}
		}

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
			//fmt.Println("hi its me the variable type", arrayType, VariableTypeString(p.meta.currentVariable.ActingType))
			return token.Value{}, errors.New("Error: array elements are of different type than type declaration")
		}

		p.meta.currentVariable.Value = arrayContentsExpressions
		value = mapVariableToTokenValue(p.meta.currentVariable)

		p.Shift()

	case token.Group:
		// Do something here to get the statements
		groupContents, ok := next.Value.True.([]token.Token)
		// FIXME:
		if !ok {
			return token.Value{}, errors.New("token.Group: next.Value.True.([]token.Token)")
		}

		groupTokens := []token.Value{}
		pa := New(groupContents)
		pa.meta.NewScopeFromScope(p.meta.currentScope)
		for pa.NextToken.Type != "" {
			var stmt token.Value
			if functionOpType == "def" {
				stmt, err = pa.GetStatement()
				if err != nil {
					return token.Value{}, err
				}
			} else {
				stmt, err = pa.GetExpression()
				if err != nil {
					return token.Value{}, err
				}
			}

			groupTokens = append(groupTokens, stmt)
			value.True = groupTokens
		}

	case token.Function:
		//fmt.Println("wtf i am here")
		next := p.NextToken
		md := next.Value.Metadata["type"]
		//fmt.Println(next)

		// Unpack tokens from function into new parser
		// Unpack tokens from each in True into a new parser
		// parse group then group, then block
		// FIXME: asserition ok check ???
		unpackedFunctionTokens := p.NextToken.Value.True.([]token.Token)
		// FIXME: size check???
		argsUnpacked := unpackedFunctionTokens[0]

		functionOpType = p.NextToken.Value.Metadata["type"].(string)
		pa := New([]token.Token{argsUnpacked})
		pa.meta.NewScopeFromScope(p.meta.currentScope)
		argExpr, err := pa.GetExpression()
		if err != nil {
			return token.Value{}, err
		}

		p.Shift()

		var returnExpr, block token.Value

		// TODO: if it is not a function defintion or a lambda call,
		// it needs to be validated that the function exists
		// Check for function definition
		if md == "def" {
			// Declare all the variables in the args so that we have them when parsing the
			// return value and the body
			for _, arg := range argExpr.True.([]token.Value) {
				err = p.meta.DeclareVariableFromTokenValue(arg)
				if err != nil {
					return token.Value{}, err
				}
			}

			// TODO: need to fix this stuff to actually check if those are valid indexes
			returnsUnpacked := unpackedFunctionTokens[1]
			bodyUnpacked := unpackedFunctionTokens[2]

			pa = New([]token.Token{returnsUnpacked})
			pa.meta.NewScopeFromScope(p.meta.currentScope)
			returnExpr, err = pa.GetExpression()
			if err != nil {
				return token.Value{}, err
			}

			pa = New([]token.Token{bodyUnpacked})
			pa.meta.NewScopeFromScope(p.meta.currentScope)
			block, err = pa.CheckBlock()
			if err != nil {
				return token.Value{}, err
			}
		}

		functionOpType = ""
		value = token.Value{
			Name:       next.Value.Name,
			AccessType: token.PrivateAccessType,
			Type:       "function",
			True: map[string]token.Value{
				"args":    argExpr,
				"returns": returnExpr,
				"body":    block,
			},
			Metadata: map[string]interface{}{
				"lambda": false,
				"type":   md.(string),
			},
		}

		if md == "def" {
			err = p.meta.DeclareVariableFromTokenValue(value)
			if err != nil {
				return token.Value{}, err
			}
		}

	default:
		return token.Value{}, errors.Errorf("Invalid token placement: %+v", p.NextToken)
	}
	//fmt.Println("value thing again", value)

	// FIXME: TODO: didn't wanna fix right now
	switch p.NextToken.Type {
	// case token.PriOp:
	// 	p.Shift()
	// 	op := p.CurrentToken
	// 	value2, verr := p.GetFactor()
	// 	if verr != nil {
	// 		return token.Value{}, verr
	// 	}
	// 	//fmt.Println("value2thing", value2)

	// value, err = p.EvaluateBinaryOperation(value, value2, op.Value)
	// if err != nil {
	// 	return token.Value{}, err
	// }
	// // FIXME: holy fuck haxorz
	// if value.Type == token.IntType {
	// 	value.String = ""
	// }

	case token.PriOp:
		p.Shift()
		op := p.CurrentToken
		value2, err := p.GetTerm()
		if err != nil {
			return token.Value{}, err
		}

		value, err = p.EvaluateBinaryOperation(value, value2, op.Value)
		if err != nil {
			return token.Value{}, err
		}
		// FIXME: holy fuck haxorz
		if value.Type == token.IntType {
			value.String = ""
		}

	case token.Increment:
		value, err = p.AddOperands(value, token.Value{
			Type: token.IntType,
			True: 1,
		})
		if err != nil {
			return token.Value{}, err
		}
	}

	return value, nil
}

// GetTerm ...
func (p *Parser) GetTerm() (token.Value, error) {
	totalTerm, err := p.GetFactor()
	if err != nil {
		return token.Value{}, err
	}

	for {
		switch p.NextToken.Type {
		// FIXME: TODO: didn't wanna fix right now
		// case token.SecOp:
		// 	p.Shift()
		// 	//fmt.Println("woah i got a secop")
		// 	op := p.CurrentToken
		// 	factor2, ferr := p.GetFactor()
		// 	if ferr != nil {
		// 		return token.Value{}, ferr
		// 	}
		// 	//fmt.Println("factor2", factor2)

		// 	totalTerm, err = p.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
		// 	if err != nil {
		// 		return token.Value{}, err
		// 	}
		// 	// FIXME: holy fuck haxorz
		// 	if totalTerp.Type == token.IntType {
		// 		totalTerp.String = strconv.Itoa(totalTerp.True.(int))
		// 	}

		case token.SecOp:
			value := p.NextToken.Value
			p.Shift()
			factor2, err := p.GetFactor()
			if err != nil {
				return token.Value{}, err
			}

			totalTerm, err = p.EvaluateBinaryOperation(totalTerm, factor2, value)
			if err != nil {
				return token.Value{}, err
			}
			// FIXME: holy fuck haxorz
			if totalTerm.Type == token.IntType {
				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
			}

		// // TODO: need to fix this....
		case token.LThan:
			p.Shift()
			op := p.CurrentToken
			factor2, err := p.GetTerm()
			if err != nil {
				return token.Value{}, err
			}

			totalTermEval, err := p.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
			if err != nil {
				return token.Value{}, err
			}
			//fmt.Println("things totalTermEval", totalTermEval)
			// FIXME: holy fuck haxorz
			// if totalTerp.Type == token.IntType {
			// TODO: should use totalTerm.String here
			//fmt.Printf("factor2before %+v\n", factor2)
			factor2.String = totalTermEval.String

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

			return totalTerm, nil
		}
	}
}

func (p *Parser) ParseExpression(tok token.Token) (token.Value, error) {
	pa := New([]token.Token{
		tok,
	})
	pa.meta.NewScopeFromScope(p.meta.currentScope)
	expression, err := pa.GetExpression()
	if err != nil {
		return token.Value{}, err
	}

	return expression, nil
}

// GetExpression ...
func (p *Parser) GetExpression() (token.Value, error) {
	switch p.NextToken.Type {
	// Will have to experiment on where to put this
	// Might need to put this in factor
	case token.Block:
		//fmt.Println("found a block")
		//fmt.Println("p.meta.currentVariable", p.meta.currentVariable)

		// TODO: Don't know if we need this if we prelod the object values, kinda hacky
		if p.meta.currentVariable.Type == VAR || p.meta.currentVariable.Type == OBJECT {
			inObject = true
			inStruct = false
		} else if p.meta.currentVariable.Type == STRUCT {
			inObject = false
			inStruct = true
		}

		//fmt.Println("inObject, inStruct", inObject, inStruct)

		block, err := p.CheckBlock()
		if err != nil {
			return token.Value{}, err
		}
		return block, nil

	// Assignment Expression
	case token.Assign:
		fmt.Println("currentVariable", p.meta.currentVariable)
		//fmt.Printf("this is an assign %+v\n", p.meta.currentVariable)
		// FIXME: I think this should go in the token.Ident case of GetStatement
		// p.DeclaredName = p.CurrentToken.Value.String
		// p.DeclaredAccessType = p.CurrentToken.Value.Type
		//fmt.Println("p.CurrentToken.Value.String", p.CurrentToken.Value)
		p.meta.currentVariable.Name = p.CurrentToken.Value.String
		p.meta.currentVariable.AccessType = accessTypeFromString(p.CurrentToken.Value.Type)

		switch p.NextToken.Value.Type {
		case "init":
			if p.meta.currentVariable.Type != UNRECOGNIZED {
				return token.Value{}, errors.New("Type specification with init is not valid: " + p.meta.currentVariable.Name)
			} else if p.meta.currentVariable.Type == FUNCTION {
				// FIXME: what is this supposed to do???
				//fmt.Println("function", p.meta.currentVariable)
				//p.meta.currentVariable.Type =
			} else {
				p.meta.currentVariable.Type = SET
			}

			fallthrough

		case "assign":
			if p.meta.currentVariable.Type == UNRECOGNIZED {
				return token.Value{}, errors.New("Undefined reference to variable: " + p.meta.currentVariable.Name)
			}

			p.Shift()
			//fmt.Println("shifted", p.NextToken)
			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}

			// FIXME: wtf is this??
			// TODO: this is where we need to take care of comparing the function return type to the variable
			if expr.Type == token.FunctionType {

			}

			if p.meta.currentVariable.Type == UNRECOGNIZED && !inStruct && !inObject {
				variable, ok := p.meta.GetVariable(p.NextToken.Value.String)
				if ok {
					p.meta.currentVariable.Type = variable.Type
				} else {

					// If we did not find it as a variable, look in the DefinedTypes map
					value2, ok := DefinedTypes[p.NextToken.Value.String]
					if !ok {
						return token.Value{}, errors.Errorf("Type not found for variable: %+v", p.meta.currentVariable)
					}

					variable = NewVariableFromTokenValue(value2)
					p.meta.currentVariable.Type = variable.Type
				}
			} else if p.meta.currentVariable.Type == SET {
				if expr.Type == token.FunctionType {
					funcT, ok := p.meta.GetVariable("someFunction")
					if !ok {
						// FIXME: something goes here???
					}

					// FIXME: check the type assertion on this ...
					p.meta.currentVariable.Type = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)
					p.meta.currentVariable.ActingType = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Acting)

				} else {
					p.meta.currentVariable.Type = variableTypeFromString(expr.Type)
				}

			} else if p.meta.currentVariable.Type == variableTypeFromString("var") {
				if expr.Type != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Type)

				} else if expr.Acting != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Acting)

				} else {
					return token.Value{}, errors.Errorf("Cannot derive variable type from expression; p.meta.currentVariable: %v expr: %v", p.meta.currentVariable, expr)
				}

			} else if p.meta.currentVariable.Type != variableTypeFromString(expr.Type) {
				// FIXME: wtf is this for?
				if expr.Type == token.Block && p.meta.currentVariable.Type == STRUCT {
					// FIXME: wtf is this for ??
					// p.meta.currentVariable.Metadata["real"] = expr.String
				} else if expr.Type == token.FunctionType {

					// FIXME: wtf ??
					os.Exit(9)

				} else if expr.Type != token.ArrayType {
					// TODO: implicit type casting here
					return token.Value{}, errors.Errorf("No implicit type casting as of now: p.meta.currentVariable.Type - %s, expr.Type - %s", VariableTypeString(p.meta.currentVariable.Type), expr.Type)
				}
			}

			if expr.Type == token.FunctionType {
				// Need to change this to an array of token or something so that cpp knows what to do
				p.meta.currentVariable.Value = expr
				// Try to use 'refs' here later

				// FIXME: this def needs to happen elsewhere
				expr.Metadata["from_func"] = true

			} else {
				p.meta.currentVariable.Value = expr.True
			}

			// Copy over all of the metadata
			for k, v := range expr.Metadata {
				p.meta.currentVariable.Metadata[k] = v
			}

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				return token.Value{}, err
			}

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}

			return token.Value{}, errors.New("Undefined variable reference: " + currentName)

		case "set":
			if p.meta.currentVariable.Type != UNRECOGNIZED && !inStruct && !inObject {
				return token.Value{}, errors.New("Type specification with set is not valid: " + p.meta.currentVariable.Name)
			}

			// If we are not in a struct, then allow the set operator to also set the type
			// However, if we are in a struct, then it should only be allowed to set the value
			// within the bounds of the type, typing-attribute, and language defined type-degrades

			// FIXME: we are erroring when declaring a struct that contains an object
			// we need to track the location
			if !inStruct {
				p.meta.currentVariable.Type = SET
			}

			p.Shift()

			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}

			if p.meta.currentVariable.Type == UNRECOGNIZED {
				variable, ok := p.meta.GetVariable(p.NextToken.Value.String)
				if ok {
					p.meta.currentVariable.Type = variable.Type
				} else {

					//fmt.Println("shit fucking shit", p.meta.currentVariable)
					// If we did not find it as a variable, look in the DefinedTypes map
					value2, ok := DefinedTypes[p.CurrentToken.Value.String]
					if !ok {
						// fmt.Println("shit", p.LastToken, value2)
						// fmt.Println("shit", p.CurrentToken, value2)
						// fmt.Println("shit", p.NextToken, value2)
						// fmt.Println("shits and stuff", DefinedTypes)
						return token.Value{}, errors.Errorf("variable still UNRECOGNIZED: %+v", p.meta.currentVariable)
					}

					variable = NewVariableFromTokenValue(value2)
					p.meta.currentVariable.Type = variable.Type
					//fmt.Println("variable2 from token", variable)
				}

			} else if p.meta.currentVariable.Type == SET {
				p.meta.currentVariable.Type = variableTypeFromString(expr.Type)

			} else if p.meta.currentVariable.Type == variableTypeFromString("var") {
				// The acting type should never be `var` so assert the acting type from the expression
				if expr.Type != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Type)

				} else if expr.Acting != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Acting)

				} else {
					return token.Value{}, errors.Errorf("Cannot derive variable type from expression; p.meta.currentVariable: %v expr: %v", p.meta.currentVariable, expr)
				}

			} else if p.meta.currentVariable.Type != variableTypeFromString(expr.Type) {
				if expr.Type == token.Block && p.meta.currentVariable.Type == STRUCT {
					// FIXME: wtf?

				} else if expr.Type == token.FunctionType {

					// fmt.Println("thing", p.CurrentToken)
					// fmt.Println("someFunction")
					funcT, ok := p.meta.GetVariable("someFunction")
					if !ok {
						// FIXME: ???
						// TODO: we couldn't find the function
					}

					p.meta.currentVariable.Type = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)
					p.meta.currentVariable.ActingType = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Acting)

				} else if expr.Type != token.ArrayType {
					// TODO: implicit type casting here
					return token.Value{}, errors.Errorf("No implicit type casting as of now: name - %s, p.meta.currentVariable.Type - %s, expr.Type - %s", p.meta.currentVariable.Name, VariableTypeString(p.meta.currentVariable.Type), expr.Type)
				}
			}

			p.meta.currentVariable.Value = expr.True
			if ref, ok := expr.Metadata["refs"]; ok {
				p.meta.currentVariable.Metadata["refs"] = ref
			}

			if fromFunc, ok := expr.Metadata["from_func"]; ok {
				p.meta.currentVariable.Metadata["from_func"] = fromFunc
			}

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				return token.Value{}, err
			}

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}

			return token.Value{}, errors.New("Undefinded variable reference: " + currentName)
		}

	case token.Increment:
		// FIXME: ???
		//fmt.Println("woah increment brah")
		// variable, ok := p.meta.GetVariable(p.meta.currentVariable.Name)
		// if !ok {
		// 	//fmt.Println("COuld not find variable:", p.meta.currentVariable.Name)
		// 	return token.Value{}, errors.New("shit")
		// }
		// value, err := p.AddOperands(mapVariableToTokenValue(variable), token.Value{
		// 	Type: token.IntType,
		// 	True: 1,
		// })
		// if err != nil {
		// 	return token.Value{}, err
		// }
		//fmt.Println("token.Increment", value)
	default:
		term, err := p.GetTerm()
		if err != nil {
			return term, err
		}

		return term, nil
	}

	return token.Value{}, errors.Errorf("Undefined behavior for token: %+v", p.NextToken)
}

func (p *Parser) ParsePrepositionFor() (token.Value, error) {
	// 1. Always expect an ident after the `for` keyword
	if p.NextToken.Type != token.Ident {
		// TODO:
		return token.Value{}, errors.Errorf("Ident not found after for token: %+v", p.NextToken)
	}

	variableName := p.NextToken.Value.String
	p.meta.currentVariable.Name = variableName
	p.meta.currentVariable.Type = variableTypeFromString(token.SetType)
	p.meta.currentVariable.ActingType = variableTypeFromString(token.SetType)
	p.meta.currentVariable.AccessType = accessTypeFromString(token.PrivateAccessType)
	p.Shift()

	// 2 . NextToken should contain a `prepositional` keyword
	if p.NextToken.Type != token.Keyword {
		return token.Value{}, errors.Errorf("Keyword not found after ident: %+v", p.NextToken)
	}

	// 3. Check the preposition
	extractKey, extractValue := false, false
	switch p.NextToken.Value.String {

	// For loop `key` composition over an iterable
	case "in":
		// 4. Declare `i` as the `key` of the index
		extractKey = true

	// For loop `value` composition over an iterable
	case "of":
		//fmt.Println("of right here rn")
		// 4. Declare `i` as the `value` of the index
		extractValue = true

		// For loop `key-value` composition over an iterable
	case "over":
		// 4. Declare `i` as an `object` containing the `key and value` of the index
		// TODO:
		extractKey, extractValue = true, true

	default:
		return token.Value{}, errors.Errorf("Preposition not found after ident token: %+v", p.NextToken)
	}

	// 5. Parse the `array` literal
	p.Shift()
	arrayExpr, err := p.GetExpression()
	if err != nil {
		// TODO:
		return arrayExpr, err
	}

	p.meta.currentVariable.Type = variableTypeFromString(arrayExpr.Acting)
	p.meta.currentVariable.ActingType = variableTypeFromString(arrayExpr.Acting)
	p.meta.currentVariable.Value = 0
	// arrayValue := p.meta.currentVariable.Value.([]token.Token)

	// FIXME: hold on
	// if extractKey && extractValue {
	// 	// TODO: make an object, put both things in it
	// 	// p.meta.currentVariable.Value = 0
	// } else if extractKey {
	// 	p.meta.currentVariable.Value = 0
	// } else {
	// 	// FIXME: check length here; try with 0 length array literal
	// 	p.meta.currentVariable.Value = arrayValue[0].True
	// }

	// currentVar := p.meta.currentVariable
	err = p.meta.DeclareVariable()
	if err != nil {
		return token.Value{}, err
	}

	// 6. Parse the body
	// p.Shift()
	body, err := p.CheckBlock()
	if err != nil {
		return body, err
	}

	bodyTokens := body.True.([]token.Value)
	varName := variableName + "_" + strconv.FormatInt(int64(time.Now().Unix()), 10)

	extraVars := []token.Value{}
	if extractKey && extractValue {
		// FIXME: ??
		// TODO: make an object, put both things in it
		// p.meta.currentVariable.Value = 0
	} else if extractKey {
		if arrayExpr.Name == variableName {
			arrayExpr.Name = "arrayBoi_" + strconv.FormatInt(int64(time.Now().Unix()), 10)
			extraVars = append(extraVars, arrayExpr)
		}
		p.meta.currentVariable.Value = 0
		extraVars = append(extraVars, token.Value{
			Name:       variableName,
			Type:       arrayExpr.Acting,
			Acting:     arrayExpr.Acting,
			True:       0,
			String:     variableName,
			AccessType: arrayExpr.AccessType,
			Metadata:   map[string]interface{}{},
		})

		bodyTokens = append([]token.Value{
			token.Value{
				Name:       variableName,
				Type:       arrayExpr.Acting,
				Acting:     arrayExpr.Acting,
				True:       0,
				AccessType: arrayExpr.AccessType,
				Metadata: map[string]interface{}{
					"refs":   varName,
					"assign": true,
				},
			},
		}, bodyTokens...)
	} else {
		// // FIXME: check length here; try with 0 length array literal
		if arrayExpr.Name == variableName {
			arrayExpr.Name = "arrayBoi_" + strconv.FormatInt(int64(time.Now().Unix()), 10)
			extraVars = append(extraVars, arrayExpr)
		}
		p.meta.currentVariable.Value = 0
		extraVars = append(extraVars, token.Value{
			Name:   variableName,
			Type:   arrayExpr.Acting,
			Acting: arrayExpr.Acting,
			// TODO: this only works with ints for now
			True:       0,
			String:     variableName,
			AccessType: arrayExpr.AccessType,
			Metadata:   map[string]interface{}{},
		})

		bodyTokens = append([]token.Value{
			token.Value{
				Name:       variableName,
				Type:       arrayExpr.Acting,
				Acting:     arrayExpr.Acting,
				True:       0,
				AccessType: arrayExpr.AccessType,
				Metadata: map[string]interface{}{
					"refs":   arrayExpr.Name + "[" + varName + "]",
					"assign": true,
				},
			},
		}, bodyTokens...)
		// os.Exit(9)
	}
	//fmt.Println("bodyTokens", bodyTokens)

	intLiteralZERO := token.Value{
		Name:       varName,
		Type:       token.IntType,
		True:       0,
		String:     "0",
		AccessType: "private",
		// Metadata: map[string]interface{}{},
	}

	intLiteralONE := intLiteralZERO
	intLiteralONE.True = 1
	intLiteralONE.String = "1"

	intLiteralARRAYLENGTH := intLiteralONE
	intLiteralARRAYLENGTH.True = arrayExpr.Metadata["length"].(int) + 1
	intLiteralARRAYLENGTH.String = fmt.Sprintf("%v", intLiteralARRAYLENGTH.True)

	// 7. Format the token as a normal `for` loop with the right metadata and
	//		variables declared within the loop for the key, value, etc
	md := map[string]interface{}{
		"start":     intLiteralZERO,
		"end":       intLiteralARRAYLENGTH,
		"step":      intLiteralONE,
		"extraVars": extraVars,
	}

	// TODO: don't think we need this for the preposition for loop
	// for k, v := range expr.Metadata {
	// 	md[k] = v
	// }

	// if extractKey && extractValue {
	// 	// TODO: make an object, put both things in it
	// 	// p.meta.currentVariable.Value = 0
	// } else if extractKey {
	// 	p.meta.currentVariable.Value = 0
	// 	// TODO: this should be the one that we want ...
	// 	extraVars[0].True = 0
	// 	extraVars[]
	// } else {
	// 	// FIXME: check length here; try with 0 length array literal
	// 	// p.meta.currentVariable.Value = arrayValue[0].

	// 	// TODO: this should be the one that we want ...
	// 	extraVars[0].True =
	// }
	// p.Shift()

	return token.Value{
		Type:     token.For,
		True:     bodyTokens,
		String:   varName + "<" + intLiteralARRAYLENGTH.String,
		Metadata: md,
	}, nil
	// return token.Value{}, nil
}

func (p *Parser) ParseStandardFor() (token.Value, error) {
	stmt, err := p.GetStatement()
	if err != nil {
		return token.Value{}, err
	}

	expr, err := p.GetExpression()
	if err != nil {
		return token.Value{}, err
	}

	expr2, err := p.GetExpression()
	if err != nil {
		return token.Value{}, err
	}

	step, err := p.SubOperands(expr2, stmt)
	if err != nil {
		return token.Value{}, err
	}

	p.Shift()

	body, err := p.CheckBlock()
	if err != nil {
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

	// //fmt.Println("stuff", expr.Metadata["left"].(token.Value).Name)
	// //fmt.Println("p.meta.currentScope", p.meta.currentScope)
	// delete(p.meta.currentScope, expr.Metadata["left"].(token.Value).Name)
	// //fmt.Println("p.meta.currentScope", p.meta.currentScope)
	expr.Metadata = map[string]interface{}{}

	return token.Value{
		Type:     token.For,
		True:     body.True.([]token.Value),
		String:   expr.String,
		Metadata: md,
	}, nil
}

func (p *Parser) GetKeyword() (token.Value, error) {
	//fmt.Println("LAST666", p.LastToken)
	//fmt.Println("CURRENT666", p.CurrentToken)
	//fmt.Println("NEXT666", p.NextToken)
	keyword := p.NextToken.Value.String
	switch keyword {
	case "for":
		// Make a new meta
		// Get a statement
		// Get an expression
		// Get another expression
		// Sub operands to find step
		p.Shift()
		//fmt.Println("GETTING LOOP")
		p.meta.NewInheritedScope()

		// Save the state from the beginning of the for loop parse
		p.SaveState()

		t, err := p.ParseStandardFor()
		if err != nil {
			// Pop back to the last state
			p.PopState()

			t, err = p.ParsePrepositionFor()
			if err != nil {
				// TODO:
				return t, err
			}
		}
		//fmt.Println("current map", p.meta.currentScope)
		_, err = p.meta.ExitScope()
		if err != nil {
			// TODO:
			return t, err
		}

		return t, nil

	case "if":
		p.Shift()
		expr, err := p.GetExpression()
		if err != nil {
			//fmt.Println("Error: Could not get expresssion")
			return token.Value{}, err
		}

		block, err := p.CheckBlock()
		if err != nil {
			return token.Value{}, err
		}

		expr.Metadata["check"] = expr.String

		return token.Value{
			Type:     token.If,
			True:     block.True.([]token.Value),
			String:   expr.String,
			Metadata: expr.Metadata,
		}, nil

	case "func":
		// Check for ident
		// Check for arguments (group)
		// Check for returns (statement/group)
		// Check for block

		//fmt.Println("wtf this name", p.meta.currentVariable)

		functionOpType = "declaration"
		// Shift away the "func" keyword
		p.Shift()

		//fmt.Println("GETTING FUNC")
		p.meta.NewInheritedScope()

		// Save the state from the beginning of the function parse
		p.SaveState()

		// 1. Always expect an ident after the `func` keyword
		if p.NextToken.Type != token.Ident {
			// TODO:
			return token.Value{}, errors.Errorf("Ident not found after for: %+v", p.NextToken)
		}

		//fmt.Printf("p.NextToken.Value %+v\n", p.NextToken.Value)
		variableName := p.NextToken.Value.String
		p.meta.currentVariable.Name = variableName
		p.meta.currentVariable.Type = variableTypeFromString(token.FunctionType)
		// p.meta.currentVariable.ActingType = variableTypeFromString(token.FunctionType)
		p.meta.currentVariable.AccessType = accessTypeFromString(token.PrivateAccessType)
		p.Shift()

		// Get arguments
		groupExpr, err := p.GetExpression()
		if err != nil {
			return token.Value{}, err
		}
		groupExpr.Name = "args"
		//fmt.Printf("groupExpr %+v, err %v\n", groupExpr, err)
		//fmt.Println("currentScope", p.meta.currentScope)
		// os.Exit(9)

		args, ok := groupExpr.True.([]token.Token)
		//fmt.Println("args, ok", args, ok)
		if ok && len(args) > 0 {
			// os.Exit(9)
			// FIXME: fix this shit
		}

		// err = p.meta.DeclareVariableFromTokenValue()
		// if err != nil {
		// 	return token.Value{}, err
		// }

		p.Shift()

		var groupExpr2 token.Value
		// Get returns
		// Save the state from before the return parse
		p.SaveState()
		if p.NextToken.Type == token.Group {
			groupExpr2, err = p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}
		}
		p.Shift()

		// FIXME: need to ensure that the function returns are what the function header says
		blockToken, err := p.CheckBlock()
		if err != nil {
			return token.Value{}, err
		}
		// p.PopState()

		functionOpType = ""

		// FIXME: We need to add this to the current scope under a "functions" key so that we can check it
		// during function usage
		tv := token.Value{
			Name:       p.meta.currentVariable.Name,
			AccessType: token.PrivateAccessType,
			Type:       "function",
			True: map[string]token.Value{
				"args":    groupExpr,
				"returns": groupExpr2,
				"body":    blockToken,
			},
			Metadata: map[string]interface{}{
				"lambda": false,
				"type":   "def",
			},
		}

		// FIXME: fix this later
		tv.AccessType = "private"

		p.meta.DeclareVariableFromTokenValue(tv)

		return tv, nil

	case "return":
		p.Shift()

		returnExpr, err := p.GetExpression()
		if err != nil {
			return token.Value{}, err
		}
		//fmt.Println("returnExpr", returnExpr)

		// FIXME: fuck it idc about checking the token type for now
		return token.Value{
			Type:   token.Keyword,
			True:   returnExpr,
			String: token.Return,
		}, nil

	case "onexit":
		fallthrough
	case "onreturn":
		fallthrough
	case "onleave":
		fallthrough
	case "defer":
		p.Shift()

		deferExpr, err := p.GetStatement()
		if err != nil {
			return token.Value{}, err
		}
		//fmt.Println("deferExpr", deferExpr)

		// FIXME: fuck it idc about checking the token type for now
		return token.Value{
			Type:   token.Keyword,
			True:   deferExpr,
			String: strings.ToUpper(keyword),
		}, nil

	default:
		return token.Value{}, errors.Errorf("Undefined keyword: %+v", p.NextToken.Value.String)
	}

	return token.Value{}, nil
}

// GetStatement ...
func (p *Parser) GetStatement() (token.Value, error) {
	var tv token.Value

	switch p.NextToken.Type {
	case token.Type:
		//fmt.Println("p.NextToken.Type", p.NextToken.Type)
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
		if p.meta.currentVariable.Type == UNRECOGNIZED {
			//fmt.Println("i am here UNRECOGNIZED")
			// TODO: maybe we should just load the entire variable at this point
			if variable, ok := p.meta.GetVariable(p.NextToken.Value.String); ok {
				variable.Metadata["assign"] = true
				//fmt.Println("FOUND THE VAR", p.NextToken.Value.String)
				p.meta.currentVariable.Type = variable.Type
				p.meta.currentVariable.Metadata = variable.Metadata
			} else {
				p.Shift()

				expr, err := p.GetExpression()
				if err != nil {
					return expr, err
				}

				return expr, nil
			}
			// TODO: make this more general later with the type map later
		} else {
			// *** struct declaration:
			// At this point we know that we are defining a struct, and we should expect
			// that the default value will not be there; if it is then this is an errorr

			if p.meta.currentVariable.Type == STRUCT {
				//fmt.Println("p.meta.currentVariable", p.meta.currentVariable, p.NextToken)
				_, err := getDefaultValueForType(token.StructType, p.NextToken.Value.String)
				if err == nil {
					return token.Value{}, errors.Errorf("Type already declared: %s", p.NextToken.Value.String)
				}

				defer func(name string) {
					DefinedTypes[name] = tv
				}(p.NextToken.Value.String)
			}

			// have the struct, the name of the new type, need to shift over the assignment
			// operator and then collect the block

			// Shift the ident token over
			// p.Shift()

			// // Shift over the assignment token
			// p.Shift()

		}
		//fmt.Println("ASSIGNMENT DECLARED TYPE", p.meta.currentVariable.Type)
		p.Shift()
		//fmt.Println(p.NextToken)
		var err error
		// FIXME: this seems kinda hacky, but w/e fix it later - GetFactor should defer it's judgement
		if p.NextToken.Type == token.Assign {
			tv, err = p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}

			// Since we know we are in a struct, when we recieve an accessor
			// operator, we need to apply the assignment to the struct property
		} else if p.NextToken.Type == token.Accessor {
			structName := p.CurrentToken.Value.String
			variable, ok := p.meta.GetVariable(structName)
			if !ok {
				return token.Value{}, errors.New("Undefined variable reference: " + structName)
			}

			// Create a new scope with only the struct properties declared
			p.meta.NewScopeFromVariable(variable)
			// Ensure that we exit this scope afterwards
			defer p.meta.ExitScope()
			fmt.Println("reflecterooni4", reflect.ValueOf(variable.Value))

			// FIXME: this is changing for some fucky reason: fix it later
			value := variable.Value.([]token.Value)

			fmt.Println("value", value, variable.Value)
			//  else {
			// 	value = value2[0].(map[string]token.Value)
			// }

			// Shift away the accessor
			p.Shift()
			fmt.Println("currentVariable", p.meta.currentVariable, p.CurrentToken)

			// Would be cool if this would work, but the variable doesn't have a type rn

			// Save the current variable and change re initilize it to a fresh variable
			// with the type as SET and then run a GetStatment to ensure the rhs
			currentVariable := *p.meta.currentVariable

			p.meta.currentVariable = &Variable{
				Type: SET,
			}

			fmt.Println("currentToken", p.CurrentToken)
			expr, err := p.GetStatement()
			if err != nil {
				return token.Value{}, err
			}

			fmt.Println("expr", expr)

			// Change the variable back
			p.meta.currentVariable = &currentVariable

			if !reflect.DeepEqual(expr, token.Value{}) {
				found := false
				for _, key := range value {
					fmt.Println("key", key)
					if key.Name == expr.Name {
						// Might need to do something here for dynamic types
						// This is also where we would need to implement some sort of 'isAssignable'
						// function that will allow us to utilize this logic elsewhere
						if expr.Type != key.Type {
							return token.Value{}, errors.Errorf("Struct property of %s is not the same type as right hand side", structName)
						}

						//value[i] = expr
						found = true
						break
					}
				}

				if found != true {
					return token.Value{}, errors.Errorf("Struct %s of type %s does not contain property %s", structName, variable.Metadata["real"], expr.Name)
				}
			}

			fmt.Println("where tf am i")

			// We have to return something here due to the way that GetStatement works
			// return token.Value{
			// 	Name: structName + "[\"" + expr.Name + "\"]",
			// 	Type: expr.Type,
			// 	True: expr.True,
			// 	Metadata: map[string]interface{}{
			// 		"assign": true,
			// 	},
			// }, nil
			return token.Value{}, nil

		} else {
			p.meta.currentVariable.Name = p.CurrentToken.Value.String
			actingTypeName := p.CurrentToken.Value.Acting
			if p.meta.currentVariable.Type == STRUCT {
				p.meta.currentVariable.Name = p.NextToken.Value.String
				// TODO: need to set the actual struct type here
				// p.meta.currentVariable.ActingType = p.CurrentToken.Value.String
				p.meta.currentVariable.Metadata["real"] = p.CurrentToken.Value.String
				actingTypeName = p.CurrentToken.Value.String
			}

			baseValue, err := getDefaultValueForType(VariableTypeString(p.meta.currentVariable.Type), actingTypeName)
			if err != nil {
				return token.Value{}, err
			}

			if VariableTypeString(p.meta.currentVariable.Type) == "var" {
				p.meta.currentVariable.ActingType = OBJECT
			}

			p.meta.currentVariable.Value = baseValue
			//fmt.Println("typemap", DefinedTypes)
			if p.meta.currentVariable.Type == STRUCT {
				fmt.Println("i am here", p.meta.currentVariable)
				fmt.Println("baseValue", baseValue)
				something := baseValue.(token.Value)
				fmt.Println("something", something)
				fmt.Println("something type", reflect.TypeOf(something))
				fmt.Println("baseValue type", reflect.TypeOf([]token.Value{baseValue.(token.Value)}))
				p.meta.currentVariable.Value = baseValue.(token.Value).True
				fmt.Println("baseValue2 type", reflect.TypeOf(p.meta.currentVariable.Value))
				// fmt.Println("baseValue1", baseValue)
				defer p.Shift()
			}

			p.meta.currentVariable.AccessType = accessTypeFromString(p.CurrentToken.Value.AccessType)
			// if it's still not set, just make it private because it's a literal or something
			if p.meta.currentVariable.AccessType < 1 {
				p.meta.currentVariable.AccessType = 1
			}

			tv = mapVariableToTokenValue(p.meta.currentVariable)
			fmt.Printf("tv %+v\n", tv)
			fmt.Println("tv type", reflect.TypeOf(tv.True))
			err = p.meta.DeclareVariable()
			if err != nil {
				return tv, err
			}

			fmt.Println("tv typeof", reflect.TypeOf(tv.True))
		}

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
		// TODO: not sure if we should return something else here
		p.Shift()
		// return token.Value{}, nil

	case token.Block:
		//fmt.Println("blockboi")
		p.Shift()
		block, err := p.CheckBlock()
		if err != nil {
			return token.Value{}, err
		}

		p.Shift()

		return block, nil

	case token.Function:
		fv, err := p.GetExpression()
		if err != nil {
			return token.Value{}, err
		}

		return fv, nil

	case "":
		// FIXME: ??
		return token.Value{}, nil

	default:
		return token.Value{}, errors.Errorf("Undefined behavior for token: %+v", p.NextToken)
	}

	return token.Value{}, nil
}

// CheckBlock ...
func (p *Parser) CheckBlock() (token.Value, error) {
	p.Shift()

	var tokensFromBlock = []token.Token{}
	var ok bool
	if p.CurrentToken.Value.True != nil {
		tokensFromBlock, ok = p.CurrentToken.Value.True.([]token.Token)
		if !ok {
			return token.Value{}, errors.Errorf("Error: Current token does not contain an array of tokens for block %+v", p.CurrentToken)
		}
	}

	// FIXME: TODO: we need to fix this hacky shit
	// works for now, but hacky as shit
	pNew := New(tokensFromBlock)
	pNew.meta.NewScopeFromScope(p.meta.currentScope)

	blockTokens := []token.Value{}

	for {

		stmt, err := pNew.GetStatement()
		if err != nil {
			return token.Value{}, err
		}

		if pNew.Index > pNew.Length()-1 && reflect.DeepEqual(stmt, token.Value{}) {
			return token.Value{
				Type: token.Block,
				True: blockTokens,
			}, nil
		}

		// // This is by-passing the blank "{}" token that is
		// // produced from the comma somtimes; need to solve
		// // it more elegantly
		// if reflect.DeepEqual(stmt, token.Value{}) {
		// 	// FIXME: fix error
		// 	return token.Value{}, errors.New("Could not get statement")
		// }

		blockTokens = append(blockTokens, stmt)
	}
}

// Semantic ...
func (p *Parser) Semantic() (token.Value, error) {
	block, err := p.CheckBlock()
	if err != nil {
		// TODO:
		return token.Value{}, err
	}

	return block, nil
}
