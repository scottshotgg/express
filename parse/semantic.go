package parse

import (
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

		//fmt.Println("holy shit gettin that var")
		variable, ok := p.meta.GetVariable(p.CurrentToken.Value.String)
		//fmt.Println(variable, ok)
		if !ok {
			// If we did not find it as a variable, look in the DefinedTypes map
			value2, ok := DefinedTypes[p.CurrentToken.Value.String]
			if !ok {
				fmt.Println()
				fmt.Println("p.meta.currentScope", p.meta.currentScope)
				return token.Value{}, errors.New("Undefined variable reference " + p.CurrentToken.Value.String)
			}

			variable = NewVariableFromTokenValue(value2)
			fmt.Println("variable from token", variable)
		}

		refs := p.CurrentToken.Value.String
		if variable.Type == FUNCTION {
			// p.meta.currentVariable.Metadata["from_func"] = true
			functionOpType = "call"
			// FIXME: here we need to look up the return value and make sure it matches
			// FIXME: make sure the args match

			// fmt.Println(p.meta.GetVariable())

			fmt.Println("I AM A CALL", variable)
			fmt.Println("last2", p.LastToken)
			fmt.Println("current2", p.CurrentToken)
			fmt.Println("next2", p.NextToken)

			p.Shift()

			// This is a struct declaration
		} else if variable.Type == STRUCT {
			// Here we need to get the default values from the type map

			fmt.Println("p.NextToken", p.NextToken)
			// os.Exit(9)

			if p.NextToken.Type == token.Block {
				inStruct = true
				defer func(typename string) {
					inStruct = false
					value.Metadata["real"] = typename
				}(p.CurrentToken.Value.String)

				if len(p.NextToken.Value.True.([]token.Token)) > 0 {
					// fmt.Printf("variable %+v\n", variable)
					// fmt.Println("p.meta.currentScope", p.meta.currentVariable)
					// fmt.Println()

					// if len(p.NextToken.Value.True.([]token.Token)) > 0 {
					// 	os.Exit(9)
					// }
					anotherVariable := *variable
					anotherVariable.Name = p.meta.currentVariable.Name

					// Need to make a new scope from the vars inside of the struct
					// pa := New(p.NextToken.Value.True.([]token.Token))
					pa := New([]token.Token{p.NextToken})
					// This gives scoping to the inside of the struct
					pa.meta.NewScopeFromScope(p.meta.currentScope)

					// Unpack all the values from the struct
					valuers := []token.Value{}
					for _, valuer := range anotherVariable.Value.([]token.Value) {
						pa.meta.DeclareVariableFromTokenValue(valuer)
						valuers = append(valuers, valuer)
					}
					anotherVariable.Value = valuers
					anotherVariable.Metadata["real"] = p.CurrentToken.Value.String

					block, err := pa.CheckBlock()
					if err != nil {
						return token.Value{}, err
					}
					// fmt.Println("block", block)

					// fmt.Println("pa.meta.currentScope", pa.meta.currentScope)

					// For some reason I couldn't get the `currentScope` variable to work here
					// valueBlock := []token.Value{}
					anotherVariableTokens := anotherVariable.Value.([]token.Value)
					for _, variableValue := range block.True.([]token.Value) {
						for i := 0; i < len(anotherVariableTokens); i++ {
							if anotherVariableTokens[i].Name == variableValue.Name {
								if anotherVariableTokens[i].Type != "var" && anotherVariableTokens[i].Type != "object" && anotherVariableTokens[i].Type != "struct" {
									if anotherVariableTokens[i].True != variableValue.True {
										variableValue.Metadata["default"] = false
										// os.Exit(9)
									}
								}

								anotherVariableTokens[i] = variableValue
								break
							}
						}
					}
					// fmt.Println()
					// fmt.Println("variable", variable)
					// fmt.Println("another", anotherVariable)
					// fmt.Println()
					// // anotherVariableTokens[0].True = 89

					// For now just shift over the block
					p.Shift()
					value = mapVariableToTokenValue(&anotherVariable)
					return value, nil
					// fmt.Println("value", value)
					// os.Exit(9)

					// fmt.Println("anotherVariable", anotherVariable)
					// os.Exit(9)
				}
				// else {
				// 	return token.Value{}, errors.New("struct without initialization does not fulfill rhs")
				// }

				inStruct = false
			}
			// TODO: fix this
			// else {
			// 	return token.Value{}, errors.Errorf("Inaccurate assignment to type, need literal %s", p.NextToken.Type)
			// }

			p.Shift()
		}

		// return p.GetExpression()
		value = mapVariableToTokenValue(variable)
		// fmt.Println("p.CurrentToken.Value.String", refs)
		value.Metadata["refs"] = refs
		// fmt.Println("value", value.Metadata)

	// case token.Group:
	// 	// //fmt.Println("getting group", next.Value)
	// 	// groupContents, ok := next.Value.True.([]token.Token)
	// 	// //fmt.Println("groupContents, ok", groupContents, ok)

	// 	// for _, groupee := range groupContents {
	// 	// 	//fmt.Println("groupee", groupee)
	// 	// }

	// 	// os.Exit(9)

	// 	value = next.Value

	// // case "":
	// // 	//fmt.Println("we at the end?")
	// // 	os.Exit(8)

	case token.Array:
		//fmt.Println("ayy rayy")
		arrayContents, ok := next.Value.True.([]token.Token)
		if !ok {
			//fmt.Println("wtf no arrray stuffs", next)
			// os.Exit(9)
			return token.Value{}, err
		}
		//fmt.Println("current", p.CurrentToken)
		//fmt.Println("next", next.Value)

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
		// p.meta.currentVariable.Metadata["vector"] = false
		// //fmt.Println("metadata", p.meta.currentVariable.Metadata["length"])
		// arrayType := VariableTypeString(p.meta.currentVariable.ActingType)
		arrayType := arrayContentsExpressions[0].Value.Type
		if arrayType == "BLOCK" {
			arrayType = "object"
		}

		if len(arrayContents) > 0 {
			//fmt.Println("arrayType", arrayType)
			for i, arrayValue := range arrayContentsExpressions {

				if arrayValue.Value.Type == "BLOCK" {
					arrayValue.Value.Type = "object"
					arrayContentsExpressions[i].Value.Type = "object"
				}

				// //fmt.Println("arrayType", arrayType, arrayValue.Value.Type)
				//fmt.Println("arrayType.Value.Type", arrayValue.Value.Type)
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
			//fmt.Println("hi its me the variable type", arrayType, VariableTypeString(p.meta.currentVariable.ActingType))
			return token.Value{}, errors.New("Error: array elements are of different type than type declaration")
		}
		// actingType := VariableTypeString(p.meta.currentVariable.ActingType)
		// if arrayValue.Value.Type != arrayType {
		// 	//fmt.Println("NOT EQUAL", arrayValue.Value.Type, arrayType)
		// }

		p.meta.currentVariable.Value = arrayContentsExpressions

		//fmt.Println("p.meta.currentVariable", p.meta.currentVariable)
		value = mapVariableToTokenValue(p.meta.currentVariable)
		//fmt.Printf("next: %+v\n", p.meta.currentVariable)
		//fmt.Printf("value: %+v\n", value)
		//fmt.Println("value", value)
		p.Shift()

	case token.Group:
		// Do something here to get the statements
		//fmt.Println("grouperooni")
		groupContents, ok := next.Value.True.([]token.Token)
		// FIXME:
		if !ok {
			//fmt.Println("wtf no group stuffs", next)
			// os.Exit(9)
			return token.Value{}, err
		}
		//fmt.Println("current", p.CurrentToken)
		//fmt.Println("next", next.Value)
		//fmt.Println("groupContents", groupContents)

		// // var groupContentsExpressions []token.Token
		// for _, piece := range groupContents {
		// 	//fmt.Println("expression1", piece)
		// 	//fmt.Println(p.ParseExpression(p.NextToken))
		// 	os.Exit(9)
		// }

		groupTokens := []token.Value{}
		pa := New(groupContents)
		pa.meta.NewScopeFromScope(p.meta.currentScope)
		for pa.NextToken.Type != "" {
			var stmt token.Value
			if functionOpType == "def" {
				//fmt.Println("FUCKTION OP TYPE", functionOpType)
				stmt, err = pa.GetStatement()
			} else {
				//fmt.Println("FUNCTION OP TYPE", functionOpType)
				stmt, err = pa.GetExpression()
				fmt.Println(p.meta.currentVariable)
				fmt.Println("last", p.LastToken)
				fmt.Println("current", p.CurrentToken)
				fmt.Println("next", p.NextToken)
				fmt.Println("I AM A FUNCTION CALL", stmt)
			}
			if err != nil {
				//fmt.Println("Error: could not parse expression inside group")
				//fmt.Println(err.Error())
				return token.Value{}, err
			}

			// fmt.Println("inside group expression", stmt)

			groupTokens = append(groupTokens, stmt)
			value.True = groupTokens
		}

		// os.Exit(9)

	case token.Function:
		fmt.Println("wtf i am here")
		next := p.NextToken
		md := next.Value.Metadata["type"]
		fmt.Println(next)

		// Unpack tokens from function into new parser
		// Unpack tokens from each in True into a new parser
		// parse group then group, then block
		unpackedFunctionTokens := p.NextToken.Value.True.([]token.Token)
		// functionTokens := []token.Value{}
		//fmt.Printf("unpackedFunctionTokens %+v\n", unpackedFunctionTokens)
		// args
		// returns
		// body
		//fmt.Println()
		argsUnpacked := unpackedFunctionTokens[0]
		//fmt.Printf("argsUnpacked %+v\n\n", argsUnpacked)

		functionOpType = p.NextToken.Value.Metadata["type"].(string)
		pa := New([]token.Token{argsUnpacked})
		pa.meta.NewScopeFromScope(p.meta.currentScope)
		argExpr, err := pa.GetExpression()
		if err != nil {
			//fmt.Println("Error: could not parse expression inside group2")
			//fmt.Println(err.Error())
			return token.Value{}, err
		}

		//fmt.Println("argExpr", argExpr)
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

			returnsUnpacked := unpackedFunctionTokens[1]
			// fmt.Printf("returnsUnpacked %+v\n\n", returnsUnpacked)

			bodyUnpacked := unpackedFunctionTokens[2]
			//fmt.Printf("bodyUnpacked %+v\n\n", bodyUnpacked)

			pa = New([]token.Token{returnsUnpacked})
			pa.meta.NewScopeFromScope(p.meta.currentScope)
			returnExpr, err = pa.GetExpression()
			if err != nil {
				//fmt.Println("Error: could not parse expression inside group3")
				//fmt.Println(err.Error())
				return token.Value{}, err
			}

			pa = New([]token.Token{bodyUnpacked})
			pa.meta.NewScopeFromScope(p.meta.currentScope)
			block, err = pa.CheckBlock()
			if err != nil {
				return token.Value{}, err
			}
			//fmt.Println("last after block", p.LastToken)
			//fmt.Println("current after block", p.CurrentToken)
			//fmt.Println("Next after block", p.NextToken)
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
			p.meta.DeclareVariableFromTokenValue(value)
		}

		fmt.Println("woah water fack")
		fmt.Println(p.meta.GetVariable("someFunction"))

	default:
		//fmt.Println("last2", p.LastToken)
		//fmt.Println("current2", p.CurrentToken)
		//fmt.Println("next2", p.NextToken)
		return token.Value{}, errors.Errorf("default %+v", p.NextToken)
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
		value2, verr := p.GetTerm()
		if verr != nil {
			return token.Value{}, verr
		}
		//fmt.Println("value2thing", value2)

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
		// p.Shift()`
	}

	// FIXME: holy fuck haxorz
	// if value.Type == token.IntType {
	// 	value.String = next.Value.String
	// }
	//fmt.Println("returning", value)
	return value, nil
}

// GetTerm ...
func (p *Parser) GetTerm() (token.Value, error) {
	//fmt.Println("GetTerm")

	totalTerm, err := p.GetFactor()
	if err != nil {
		return token.Value{}, err
	}
	//fmt.Println("totalTERM", totalTerm)

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
			p.Shift()
			//fmt.Println("woah i got a secop")
			op := p.CurrentToken
			factor2, ferr := p.GetFactor()
			if ferr != nil {
				return token.Value{}, ferr
			}
			//fmt.Println("factor2", factor2)

			totalTerm, err = p.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
			if err != nil {
				return token.Value{}, err
			}
			// FIXME: holy fuck haxorz
			if totalTerm.Type == token.IntType {
				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
			}

		// // TODO: need to fix this....
		case token.LThan:
			//fmt.Println("in the lthan")
			// ident := p.LastToken
			// nextTokenOpString := p.NextToken.Value.String
			p.Shift()
			op := p.CurrentToken
			factor2, ferr := p.GetTerm()
			if ferr != nil {
				return token.Value{}, ferr
			}
			//fmt.Println("lthan totalTerm", totalTerm)
			//fmt.Println("lthan factor2", factor2)
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
			// }
			//fmt.Printf("totalTerm %+v\n", totalTerm)
			//fmt.Printf("totalTermEval %+v\n", totalTermEval)
			//fmt.Println("factor2", factor2)
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
			//fmt.Println("i am here", p.NextToken)
			//fmt.Println("totalTerm", totalTerm)
			return totalTerm, nil
		}
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
		//fmt.Println("Error: could not parse expression inside array")
		//fmt.Println(err.Error())
		return token.Value{}, err
	}

	//fmt.Println("insidexpression", expression)

	return expression, nil
}

// GetExpression ...
func (p *Parser) GetExpression() (token.Value, error) {
	//fmt.Println("GetExpression")
	//fmt.Printf("p.NextToken %+v\n", p.NextToken)

	switch p.NextToken.Type {
	// Will have to experiment on where to put this
	// Might need to put this in factor
	case token.Block:
		fmt.Println("found a block")
		fmt.Println("p.meta.currentVariable", p.meta.currentVariable)

		// TODO: Don't know if we need this if we prelod the object values, kinda hacky
		if p.meta.currentVariable.Type == VAR || p.meta.currentVariable.Type == OBJECT {
			inObject = true
			inStruct = false
		} else if p.meta.currentVariable.Type == STRUCT {
			inObject = false
			inStruct = true
		}

		fmt.Println("inObject, inStruct", inObject, inStruct)

		block, err := p.CheckBlock()
		if err != nil {
			//fmt.Println("waddup blockboi", err)
			return token.Value{}, err
		}
		return block, nil

	// Assignment Expression
	case token.Assign:
		//fmt.Printf("this is an assign %+v\n", p.meta.currentVariable)
		// FIXME: I think this should go in the token.Ident case of GetStatement
		// p.DeclaredName = p.CurrentToken.Value.String
		// p.DeclaredAccessType = p.CurrentToken.Value.Type
		fmt.Println("p.CurrentToken.Value.String", p.CurrentToken.Value)
		p.meta.currentVariable.Name = p.CurrentToken.Value.String
		p.meta.currentVariable.AccessType = accessTypeFromString(p.CurrentToken.Value.Type)

		switch p.NextToken.Value.Type {
		case "init":
			if p.meta.currentVariable.Type != UNRECOGNIZED {
				return token.Value{}, errors.New("Type specification with init is not valid: " + p.meta.currentVariable.Name)
			} else if p.meta.currentVariable.Type == FUNCTION {
				fmt.Println("function", p.meta.currentVariable)
				//p.meta.currentVariable.Type =
			} else {
				p.meta.currentVariable.Type = SET
			}

			// if p.meta.currentVariable.Name == "something" {
			// 	fmt.Println("something:", p.meta.currentVariable)
			// 	os.Exit(9)
			// }
			fallthrough

		case "assign":
			if p.meta.currentVariable.Type == UNRECOGNIZED {
				return token.Value{}, errors.New("Undefined reference to variable")
			}

			p.Shift()
			//fmt.Println("shifted", p.NextToken)
			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}
			fmt.Printf("expr in assign %+v\n", expr)

			// TODO: this is where we need to take care of comparing the function return type to the variable
			if expr.Type == token.FunctionType {
				// p.meta.currentVariable.Metadata["from_func"] = true
				fmt.Println("currentVar", p.meta.currentVariable)

				// fmt.Println(p.meta.GetVariable("someFunction"))

				// we are going to need to deal with different things here:
				//	can variables be assigned to a function such that:
				//	something := someFunc
				//	we also need to deal with allowing function definitions here
				//	but for now we will only do calls

				// funcT, ok := p.meta.GetVariable(expr.Name)
				// if !ok {
				// 	return token.Value{}, errors.New("Could not find function")
				// }

				// fmt.Println("expr", expr)
				// expr = mapVariableToTokenValue(funcT)
				// fmt.Println("expr after", expr)
				// os.Exit(9)

			}

			if p.meta.currentVariable.Type == UNRECOGNIZED && !inStruct && !inObject {
				variable, ok := p.meta.GetVariable(p.NextToken.Value.String)
				if ok {
					p.meta.currentVariable.Type = variable.Type
				} else {

					// If we did not find it as a variable, look in the DefinedTypes map
					value2, ok := DefinedTypes[p.NextToken.Value.String]
					if !ok {
						return token.Value{}, errors.Errorf("variable still UNRECOGNIZED: %+v", p.meta.currentVariable)
					}

					variable = NewVariableFromTokenValue(value2)
					p.meta.currentVariable.Type = variable.Type
					fmt.Println("variable2 from token", variable)
				}
			} else if p.meta.currentVariable.Type == SET {
				// if variable, ok := p.meta.GetVariable(p.meta.currentVariable.Name); ok {
				// 	//fmt.Println("variable", variable, "ok", ok)
				// }
				//  else {
				// 	//fmt.Println("wtf am i doing here", p.meta.currentVariable)
				// 	os.Exit(9)
				// }
				if expr.Type == token.FunctionType {

					fmt.Println("thing", p.CurrentToken)
					fmt.Println("someFunction")
					funcT, ok := p.meta.GetVariable("someFunction")
					if !ok {
						// TODO: we couldn't find the function
					}

					fmt.Println(p.meta.currentVariable)
					fmt.Println("returns", funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)

					p.meta.currentVariable.Type = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)
					p.meta.currentVariable.ActingType = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Acting)

					fmt.Println("butts", p.meta.currentVariable)

				} else {

					p.meta.currentVariable.Type = variableTypeFromString(expr.Type)

				}
			} else if p.meta.currentVariable.Type == variableTypeFromString("var") {

				fmt.Println("woah its me", expr.Type, expr.Acting, p.meta.currentVariable.Type, p.meta.currentVariable.ActingType)
				if expr.Type != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Type)

				} else if expr.Acting != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Acting)

				} else {
					return token.Value{}, errors.Errorf("Not sure how to assert variable type; p.meta.currentVariable: %v\nexpr: %v", p.meta.currentVariable, expr)
				}

			} else if p.meta.currentVariable.Type != variableTypeFromString(expr.Type) {
				// FIXME: wtf is this for?
				if expr.Type == token.Block && p.meta.currentVariable.Type == STRUCT {
					// p.meta.currentVariable.Metadata["real"] = expr.String
				} else if expr.Type == token.FunctionType {
					fmt.Println("FUNCTION", expr)

					fmt.Println("scope", p.meta.currentScope)

					os.Exit(9)

				} else if expr.Type != token.ArrayType {
					//fmt.Println(VariableTypeString(p.meta.currentVariable.Type), expr.Type)
					// TODO: implicit type casting here
					fmt.Println("expr", expr)
					return token.Value{}, errors.Errorf("No implicit type casting as of now: p.meta.currentVariable.Type - %s, expr.Type - %s", VariableTypeString(p.meta.currentVariable.Type), expr.Type)
				}
			}
			// else {
			// 	//fmt.Printf("wtf typerooni: %+v\n", p.meta.currentVariable)
			// }

			if expr.Type == token.FunctionType {
				// Need to change this to an array of token or something so that cpp knows what to do
				p.meta.currentVariable.Value = expr
				// Try to use 'refs' here later

				// FIXME: this def needs to happen elsewhere
				// newVar := *p.meta.currentVariable
				// newVar.Metadata["from_func"] = true
				// p.meta.currentVariable = &newVar
				expr.Metadata["from_func"] = true

				// p.meta.currentVariable.Metadata["refs"] =
			} else {
				p.meta.currentVariable.Value = expr.True
			}

			// Copy over all of the metadata
			for k, v := range expr.Metadata {
				p.meta.currentVariable.Metadata[k] = v
			}
			// if ref, ok := expr.Metadata["refs"]; ok {
			// 	//fmt.Println("there was a ref")
			// 	p.meta.currentVariable.Metadata["refs"] = ref
			// }
			//fmt.Println("p.meta.currentVariable2", p.meta.currentVariable)

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				// TODO:
				//fmt.Println("declareVariable error", err.Error())
				// os.Exit(9)
				return token.Value{}, err
			}
			// fmt.Println("p.meta.currentScope after", p.meta.currentScope)

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}

			return token.Value{}, errors.New("Could not find variable: " + currentName)

		case "set":
			if p.meta.currentVariable.Type != UNRECOGNIZED && !inStruct && !inObject {
				return token.Value{}, errors.New("Type specification with set is not valid")
			}

			// If we are not in a struct, then allow the set operator to also set the type
			// However, if we are in a struct, then it should only be allowed to set the value
			// within the bounds of the type, typing-attribute, and language defined type-degrades
			fmt.Println("inStruct", inStruct)

			// FIXME: we are erroring when declaring a struct that contains an object
			// we need to track the location
			if !inStruct {
				p.meta.currentVariable.Type = SET
			}

			// if variable, ok := p.meta.currentScope[p.meta.currentVariable.Name]; ok {
			// 	//fmt.Println("Error: Variable already declared in this scope", variable)
			// 	return token.Value{}, errors.New("Error: Variable already declared in this scope")
			// }

			// // Old set stuff
			// p.Shift()
			// //fmt.Println("what do", p.NextToken)
			// expr, err := p.GetExpression()
			// if err != nil {
			// 	return token.Value{}, err
			// }
			// //fmt.Printf("expr in set %+v\n", expr)

			// p.meta.currentVariable.Type = variableTypeFromString(expr.Type)
			// p.meta.currentVariable.Value = expr.True
			// p.meta.currentVariable.Metadata = expr.Metadata

			// // TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			// currentName := p.meta.currentVariable.Name
			// err = p.meta.DeclareVariable()
			// if err != nil {
			// 	// TODO:
			// 	//fmt.Println("declareVariable error", err.Error())
			// 	return token.Value{}, err
			// }

			// if variable, ok := p.meta.GetVariable(currentName); ok {
			// 	// Map it over to a token for now
			// 	return mapVariableToTokenValue(variable), nil
			// }
			// return token.Value{}, errors.New("Could not find variable: " + currentName)

			p.Shift()
			// //fmt.Println("what do", p.NextToken)
			expr, err := p.GetExpression()
			if err != nil {
				return token.Value{}, err
			}

			if p.meta.currentVariable.Type == UNRECOGNIZED {
				variable, ok := p.meta.GetVariable(p.NextToken.Value.String)
				if ok {
					p.meta.currentVariable.Type = variable.Type
				} else {

					fmt.Println("shit fucking shit", p.meta.currentVariable)
					// If we did not find it as a variable, look in the DefinedTypes map
					value2, ok := DefinedTypes[p.CurrentToken.Value.String]
					if !ok {
						fmt.Println("shit", p.LastToken, value2)
						fmt.Println("shit", p.CurrentToken, value2)
						fmt.Println("shit", p.NextToken, value2)
						fmt.Println("shits and stuff", DefinedTypes)
						return token.Value{}, errors.Errorf("variable still UNRECOGNIZED: %+v", p.meta.currentVariable)
					}

					variable = NewVariableFromTokenValue(value2)
					p.meta.currentVariable.Type = variable.Type
					fmt.Println("variable2 from token", variable)
				}
			} else if p.meta.currentVariable.Type == SET {
				// if variable, ok := p.meta.GetVariable(p.meta.currentVariable.Name); ok {
				// 	//fmt.Println("variable", variable, "ok", ok)
				// }
				//  else {
				// 	//fmt.Println("wtf am i doing here", p.meta.currentVariable)
				// 	os.Exit(9)
				// }

				p.meta.currentVariable.Type = variableTypeFromString(expr.Type)

			} else if p.meta.currentVariable.Type == variableTypeFromString("var") {
				// The acting type should never be `var` so assert the acting type from the expression
				if expr.Type != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Type)
				} else if expr.Acting != "var" {
					p.meta.currentVariable.ActingType = variableTypeFromString(expr.Acting)
				} else {
					return token.Value{}, errors.New("wtf")
				}

			} else if p.meta.currentVariable.Type != variableTypeFromString(expr.Type) {
				if expr.Type == token.Block && p.meta.currentVariable.Type == STRUCT {
					// FIXME: wtf?
					// p.meta.currentVariable.Metadata["real"] = expr.String

				} else if expr.Type == token.FunctionType {

					fmt.Println("thing", p.CurrentToken)
					fmt.Println("someFunction")
					funcT, ok := p.meta.GetVariable("someFunction")
					if !ok {
						// TODO: we couldn't find the function
					}

					fmt.Println(p.meta.currentVariable)
					fmt.Println("returns", funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)

					p.meta.currentVariable.Type = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Type)
					p.meta.currentVariable.ActingType = variableTypeFromString(funcT.Value.(map[string]token.Value)["returns"].True.([]token.Value)[0].Acting)

					fmt.Println("butts2", p.meta.currentVariable)

				} else if expr.Type != token.ArrayType {
					//fmt.Println(VariableTypeString(p.meta.currentVariable.Type), expr.Type)
					// TODO: implicit type casting here
					fmt.Println("expr2", expr, p.meta.currentVariable)
					return token.Value{}, errors.Errorf("No implicit type casting as of now: p.meta.currentVariable.Type - %s, expr.Type - %s", VariableTypeString(p.meta.currentVariable.Type), expr.Type)
				}
			}
			// else {
			// 	//fmt.Printf("wtf typerooni: %+v\n", p.meta.currentVariable)
			// }

			p.meta.currentVariable.Value = expr.True
			if ref, ok := expr.Metadata["refs"]; ok {
				//fmt.Println("there was a ref")
				p.meta.currentVariable.Metadata["refs"] = ref
			}
			if fromFunc, ok := expr.Metadata["from_func"]; ok {
				//fmt.Println("there was a ref")
				p.meta.currentVariable.Metadata["from_func"] = fromFunc
			}
			//fmt.Println("p.meta.currentVariable2", p.meta.currentVariable)

			// TODO: doing this to ensure that it is in the map and findable ... not sure if we need to or should
			currentName := p.meta.currentVariable.Name
			err = p.meta.DeclareVariable()
			if err != nil {
				// TODO:
				//fmt.Println("declareVariable error", err.Error())
				// os.Exit(9)
				return token.Value{}, err
			}
			//fmt.Println("p.meta.currentScope after", p.meta.currentScope)

			if variable, ok := p.meta.GetVariable(currentName); ok {
				// Map it over to a token for now
				return mapVariableToTokenValue(variable), nil
			}

			return token.Value{}, errors.New("Could not find variable: " + currentName)
		}

	// case token.LThan:
	// 	//fmt.Println("wtf")
	// 	//fmt.Println("current", p.CurrentToken)
	// 	//fmt.Println("next", p.NextToken)
	// 	p.Shift()
	// 	term, err := p.GetTerm()
	// 	if err != nil {
	// 		return token.Value{}, err
	// 	}
	// 	return term, nil

	case token.Increment:
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
		return p.GetTerm()
	}

	return token.Value{}, errors.Errorf("default %+v", p.NextToken)
}

func (p *Parser) ParsePrepositionFor() (token.Value, error) {
	//fmt.Println("ParsePrepositionFor")

	// 1. Always expect an ident after the `for` keyword
	if p.NextToken.Type != token.Ident {
		// TODO:
		return token.Value{}, errors.Errorf("Ident not found after for: %+v", p.NextToken)
	}

	//fmt.Printf("p.NextToken.Value %+v\n", p.NextToken.Value)
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
		return token.Value{}, errors.Errorf("Preposition not found: %+v", p.NextToken)
	}

	//fmt.Println("p.meta.currentVariable", p.meta.currentVariable)
	//fmt.Println("extractKey, extractValue:", extractKey, extractValue)

	// 5. Parse the `array` literal
	p.Shift()
	arrayExpr, err := p.GetExpression()
	if err != nil {
		// TODO:
		return arrayExpr, err
	}
	//fmt.Println("arrayExpr", arrayExpr)

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
		//fmt.Println("Could not check block")
		return body, err
	}
	//fmt.Println("body", body)

	bodyTokens := body.True.([]token.Value)
	//fmt.Println("bodyTokens", bodyTokens)

	varName := variableName + "_" + strconv.FormatInt(int64(time.Now().Unix()), 10)

	extraVars := []token.Value{}
	if extractKey && extractValue {
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
		//fmt.Println("inside the value thing")
		// // FIXME: check length here; try with 0 length array literal
		// p.meta.currentVariable.Value = arrayValue[0].Value.True
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
		// TODO:
		// os.Exit(9)
		return token.Value{}, err
	}

	expr, err := p.GetExpression()
	if err != nil {
		//fmt.Println("Error: Could not get expression")
		// os.Exit(9)
		return token.Value{}, err
	}

	expr2, err := p.GetExpression()
	if err != nil {
		//fmt.Println("Error: Could not get expression2")
		// os.Exit(9)
		return token.Value{}, err
	}
	//fmt.Println("stmt", stmt)
	//fmt.Println("expr1", expr)
	//fmt.Println("expr2", expr2)
	// os.Exit(9)

	step, err := p.SubOperands(expr2, stmt)
	if err != nil {
		//fmt.Println("Could not sub operands")
		return token.Value{}, err
	}

	p.Shift()
	body, err := p.CheckBlock()
	if err != nil {
		//fmt.Println("Could not check block")
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
			//fmt.Println("Error: Could not get block")
			return token.Value{}, err
		}

		//fmt.Printf("expr %+v\n", expr)
		//fmt.Printf("block %+v\n", block)
		//fmt.Printf("next555 %+v\n", p.NextToken)

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

		fmt.Println("wtf this name", p.meta.currentVariable)

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
		//fmt.Println("woah idk", p.NextToken)
		// os.Exit(9)
		return token.Value{}, err
	}

	return token.Value{}, nil
}

// GetStatement ...
func (p *Parser) GetStatement() (token.Value, error) {
	var tv token.Value
	//fmt.Println("GetStatement")
	//fmt.Println("p.NextToken", p.NextToken)
	// p.Shift()
	// fmt.Println("p.meta.currentVariable in GetStatement", p.meta.currentVariable)
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
		//fmt.Println("idnet p.meta.currentVariable", p.meta.currentVariable)
		//fmt.Println("ident", p.NextToken)
		//fmt.Println("declaredMap", p.meta.currentScope)
		if p.meta.currentVariable.Type == UNRECOGNIZED {
			//fmt.Println("i am here UNRECOGNIZED")
			// TODO: maybe we should just load the entire variable at this point
			if variable, ok := p.meta.GetVariable(p.NextToken.Value.String); ok {
				variable.Metadata["assign"] = true
				//fmt.Println("FOUND THE VAR", p.NextToken.Value.String)
				p.meta.currentVariable.Type = variable.Type
				p.meta.currentVariable.Metadata = variable.Metadata
			} else {
				// If its unrecognized and we cant find it, it doesn't exist
				//fmt.Println("in the elser")
				//fmt.Println("ASSIGNMENT DECLARED VALUE", m.DeclaredValue)
				p.Shift()
				expr, err := p.GetExpression()
				//fmt.Printf("THIS IS THE EXPRESSION %+v %s\n", expr, err)
				// return p.GetExpression()
				//fmt.Println("expr, err", expr, err)

				return expr, err
			}
			// TODO: make this more general later with the type map later
		} else {
			// *** struct declaration:
			// At this point we know that we are defining a struct, and we should expect
			// that the default value will not be there; if it is then this is an errorr

			if p.meta.currentVariable.Type == STRUCT {
				fmt.Println("p.meta.currentVariable", p.meta.currentVariable, p.NextToken)
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
		if p.NextToken.Type == "ASSIGN" {
			tv, err = p.GetExpression()
			// fmt.Println("nofind THIS IS THE EXPRESSION", tv, err)
			if err != nil {
				//fmt.Println("getExpressionErr", err)
				// os.Exit(9)
				return token.Value{}, err
			}
			//fmt.Println("TVTVTV", tv)
			//fmt.Println("another", p.NextToken)
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

			fmt.Println("typemap", DefinedTypes)
			fmt.Println("baseValue", baseValue)
			if p.meta.currentVariable.Type == STRUCT {
				baseValue = baseValue.(token.Value).True
				defer p.Shift()
			}
			p.meta.currentVariable.Value = baseValue
			p.meta.currentVariable.AccessType = accessTypeFromString(p.CurrentToken.Value.AccessType)
			// if it's still not set, just make it private because it's a literal or something
			if p.meta.currentVariable.AccessType < 1 {
				p.meta.currentVariable.AccessType = 1
			}
			//fmt.Printf("p.CurrentToken %+v\n", p.CurrentToken)
			//fmt.Printf("else p.meta.currentVariable %+v\n", p.meta.currentVariable)
			tv = mapVariableToTokenValue(p.meta.currentVariable)
			err = p.meta.DeclareVariable()
			if err != nil {
				return tv, err
			}
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
		// //fmt.Println("should we have gotten this here?")
		// os.Exit(9)
		// TODO: not sure if we should return something else here
		p.Shift()
		// return token.Value{}, nil

	case token.Block:
		//fmt.Println("blockboi")
		p.Shift()
		block, err := p.CheckBlock()
		if err != nil {
			//fmt.Println("waddup blockboi", err)
			// os.Exit(9)
			return token.Value{}, err
		}
		p.Shift()
		return block, err

	case token.Function:
		//fmt.Println("hey i found a function")

		return p.GetExpression()

	case "":
		//fmt.Println("im fuckin here")
		return token.Value{}, nil

	default:
		// TODO: this causes infinite loops when you cant parse
		//fmt.Println("hey its me the default", p.NextToken)
		// os.Exit(9)
		return token.Value{}, err
	}

	return token.Value{}, nil
}

// CheckBlock ...
func (p *Parser) CheckBlock() (token.Value, error) {
	//fmt.Printf("CheckBlock %+v\n", p.CurrentToken)

	p.Shift()
	// fmt.Printf("CheckBlock2 %+v\n", p.CurrentToken)
	// logger.Debug("CheckBlock", zap.String("p.CurrentToken", fmt.Sprintf("%+v", p.CurrentToken)))
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

		//fmt.Println(pNew.NextToken)
		// if reflect.DeepEqual(pNew.NextToken, token.Token{}) {
		if pNew.Index > pNew.Length()-1 && reflect.DeepEqual(stmt, token.Value{}) {
			// //fmt.Println(p.meta.GetVariable("a"))
			// p.meta.NewScopeFromScope(pNew.meta.currentScope)
			// fmt.Println("blockTokens", blockTokens)

			return token.Value{
				Type: token.Block,
				True: blockTokens,
			}, nil
		}

		// This is by-passing the blank "{}" token that is
		// produced from the comma somtimes; need to solve
		// it more elegantly
		if reflect.DeepEqual(stmt, token.Value{}) {
			// return token.Value{
			// 	Type: token.Block,
			// 	True: blockTokens,
			// }, nil
			return token.Value{}, errors.New("Could not get statement")
		}

		blockTokens = append(blockTokens, stmt)

		// p.meta.NewVariable()
		//fmt.Println("CheckBlock currentScope: ", pNew.meta.currentScope)
		//fmt.Println()
	}
}

// Semantic ...
func (p *Parser) Semantic() (token.Value, error) {
	// defer func() {
	// 	for k, v := range p.meta.currentScope {
	// 		logger.Debug("p.meta.currentScope", zap.String("scope", fmt.Sprintf("%+v", p.meta.currentScope)))
	// 	}
	// }()

	block, err := p.CheckBlock()
	if err != nil {
		// TODO:
		return token.Value{}, err
	}
	//fmt.Println("block", block)
	//fmt.Println()

	//fmt.Println("End currentScope: ", p.meta.currentScope)
	//fmt.Println()

	return block, nil
}
