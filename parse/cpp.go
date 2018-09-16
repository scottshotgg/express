package parse

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

/*
	This should really be an interface that can be implemented with
	certain functions so that other transpilations can be implemeneted
*/

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	// f          string
	r          *rand.Rand
	insideLoop bool

	// TODO: FIXME: this is causing functions to be compiled in every single file
	structStrings = "\n"

	LibBase = ""

	blockDepth = 0
)

func (p *Parser) TranslateFunctionCall(t token.Value) (string, error) {
	functionString := "\n"

	trueValue, ok := t.True.(map[string]token.Value)
	if !ok {
		return "", errors.New("Could not assert value of function call")
	}

	functionString += t.Name + "("
	argsInterface := trueValue["args"].True
	if argsInterface != nil {
		args, ok := argsInterface.([]token.Value)
		if !ok {
			return "", errors.New("Could not assert list of arguments in function call")
		}

		// FIXME: change this loop to standard for loop until the end and add commas
		// everytime, w/e
		for i, arg := range args {
			ref, ok := arg.Metadata["refs"]
			if ok {
				assert, ok := ref.(string)
				if !ok {
					return "", errors.New("Could not assert `ref` from metadata")
				}
				functionString += assert

			} else if arg.Type == "BLOCK" || arg.Type == "object" || arg.Type == "var" {
				arg.Name = arg.Name + "_" + RandStringBytesMaskImprSrc(10)

				objectString, err := p.TranslateObject(arg, "")
				if err != nil {
					return "", err
				}
				functionString = objectString + functionString + arg.Name
			} else if arg.Type == "function" {
				innerFunctionCall, err := p.TranslateFunctionCall(arg)
				if err != nil {
					return "", err
				}
				// Cut off the newline and semicolon that is on the end
				functionString += innerFunctionCall[:len(innerFunctionCall)-2]
			} else if arg.Type == "string" || arg.Type == "char" {
				// The extra `, ""` is a hack to get strings to work in the template
				// functions in the C++ bindings

				functionString += fmt.Sprintf("\"%+v\"", arg.True)
			} else {
				functionString += fmt.Sprintf("%+v", arg.True)
			}

			// Append a comma if this is not the last arg
			if i != len(args)-1 {
				functionString += ","
			}
		}

		if len(args) == 1 && args[0].Type == "string" || args[0].Type == "char" {
			functionString += ", \"\""
		}
	}

	// Close the fucntion
	functionString += ");\n"

	return functionString, nil
}

func (p *Parser) TranslateFunctionDef(t token.Value) (string, error) {
	functionString := "\n"

	trueValue, ok := t.True.(map[string]token.Value)
	if !ok {
		return "", errors.New("Could not assert value of function definition")
	}

	// FIXME: later on we need to create a `statement` structure that we can call .String() on that will create the C++ statement

	// Parse the returns first since that is the first value of a C++ function
	returnsInterface := trueValue["returns"].True

	// If the function does not return anything, the equivalent C++ return type is `void`
	if returnsInterface == nil {
		functionString += "void "

		// Otherwise figure out the return type
	} else {
		returns, ok := returnsInterface.([]token.Value)
		if !ok {
			return "", errors.Errorf("Could not assert `returns[0]` to token.Value: %+v", returns[0])
		}
		if len(returns) == 0 {
			return "", errors.New("Returns was non-nil, but len(returns) was 0")
		}
		// FIXME: fix multi returns here; only use the first one for now
		firstReturn := returns[0]

		// Ensure that the return has a type
		if firstReturn.Type == "" {
			return "", errors.Errorf("No return type found on `returns[0]`: %+v", returns[0])
		}

		// If we are returning a `BLOCK`, `object`, or `var` we need to use the equivalent C++ `var` type
		if firstReturn.Type == "BLOCK" || firstReturn.Type == "object" || firstReturn.Type == "var" {
			functionString += "var "

			// Otherwise just return the type as all other Express types match C++ types right now
		} else {
			functionString += firstReturn.Type + " "
		}
	}

	// Append the name of the function
	functionString += t.Name + "("

	// Parse the arguments
	argsInterface := trueValue["args"].True

	// If there are arguments then we need to check them
	if argsInterface != nil {
		args, ok := argsInterface.([]token.Value)
		if !ok {
			return "", errors.New("Could not assert value of arguments in function definition")
		}

		// FIXME:  change this loop to standard for loop until the end and add commas
		// everytime, w/e
		for i, arg := range args {
			// If the argument type is a `BLOCK`, `object`, or `var` type, then we need to use the equivalent C++ `var` type
			if arg.Type == "BLOCK" || arg.Type == "object" || arg.Type == "var" {
				functionString += "var " + arg.Name

				// Otherwise the C++ arg type is the same as all other Express types match C++ types right now
			} else {
				functionString += arg.Type + " " + arg.Name
			}

			// Append a comma if this isn't the last argument
			if i != len(args)-1 {
				functionString += ","
			}
		}
	}

	// Close the function
	functionString += ")"

	// Translate the body of the function
	bodyString, err := p.TranslateBlock(trueValue["body"])
	if err != nil {
		return "", err
	}

	// Add the onreturn defer stack as the first declaration
	functionString += "{\ndefer onReturnFuncs;\n" + bodyString + "}"

	return functionString, nil
}

func (p *Parser) TranslateArray(t token.Value) (string, error) {
	arrayString := ""

	trueValue, ok := t.True.([]token.Token)
	if !ok {
		return "", errors.New("Could not assert array value")
	}

	// FIXME: assuming only single type arrays until I have time to do multi type arrays in C++
	arrayType := t.Acting

	// TODO: was lazy, could do this with the defers, need to fix
	if arrayType == "string" {
		arrayType = "std::" + arrayType
	} else if arrayType == "object" {
		arrayType = "var"
	}

	// Start the array declaraion off
	arrayString += arrayType + " " + t.Name + "[] = { "

	// If is is a var then close the array
	if arrayType == "var" {
		arrayString += " };\n"
	}

	// Range over each index of the array
	for i, v := range trueValue {
		sprintString := "%v"
		// FIXME: for some reason array composition using only one object variable causes
		// each object to change when one does because of the pointers and stuff
		// It's because the memory is not copied when putting the var into the map, so we need
		// to create a copy constructor for the `var` type in the runtime,
		// however I don't wanna deal with wrangling C++ memory shit right now brah so we gonna do da hackz0rz

		ref, ok := v.Value.Metadata["refs"]

		// If the Type is an object and if references a previously declared variable then set the value in the array
		if v.Value.Type != "object" && ok {
			arrayString += ref.(string)

			// If the type is a string then we need to quote it
		} else if v.Value.Type == "string" {
			sprintString = "\"" + sprintString + "\""
			arrayString += fmt.Sprintf(sprintString, v.Value.True)

			// Again check if it is an object (I know... this is fucky, but im only one superhero, comrade ¯\_(ツ)_/¯)
		} else if v.Value.Type == "object" {
			// So if it's an object, instead of being able to copy it, we need to make a new `var` type in C++ so use a anonymized name for it
			v.Value.Name = v.Value.Name + "_" + RandStringBytesMaskImprSrc(10)

			// Translate the object
			objectString, err := p.TranslateObject(v.Value, "")
			if err != nil {
				return "", err
			}
			arrayString += "{" + objectString
			// Add it to the array
			arrayString += fmt.Sprintf("%s[%d] = %s;\n}\n", t.Name, i, v.Value.Name)
			continue

			// Otherwise just add the value to the array
		} else {
			arrayString += fmt.Sprintf(sprintString, v.Value.True)
		}

		// FIXME: Change the loop to not require this; loop till before the last one
		if i != len(trueValue)-1 {
			arrayString += ", "
		}
	}

	// TODO: try to remember why this is here and working
	if arrayType != "var" {
		arrayString += " };\n"
	}

	return arrayString, nil
}

// FIXME: just use var for now, but later we will try to not use var
// FIXME: ideally we want to store this in a "symbol map" with the defaults already there
func (p *Parser) TranslateObject(t token.Value, objName string) (string, error) {

	// If the name of the object that was passed in was not empty, use that
	// This is for passing an anonymized name into the p.Translate object
	tName := t.Name
	if objName != "" {
		tName = objName
	}

	_, ok := t.Metadata["from_func"]

	// If the object is being returned from a function
	if ok {
		asserted, ok := t.True.(token.Value)
		// If we are able to assert the value
		if !ok {
			return "", errors.New("Could not assert value of object")
		}

		// p.Translate the function
		funcCall, err := p.TranslateFunctionCall(asserted)
		if err != nil {
			return "", err
		}

		// might have to make a custom tName
		// [1:] to shave off the newline from the translation: remove if it causes problems
		return "var " + tName + " = " + funcCall[1:], nil
	}

	// Objects a the `var` type in the C++ runtime
	objectString := "var " + tName + " = {};\n"

	tTrues, ok := t.True.([]token.Value)
	if !ok {
		return "", errors.New("Could not assert true value of object")
	}

	// Range over the statements inside the object
	for _, v := range tTrues {
		objectValue := v.True

		var fromFuncOk bool

		_, ok := v.True.([]token.Value)
		if !ok {
			// If we can't assert the value, check if it references something; if it does set the value
			ref, ok := v.Metadata["refs"]
			if ok {
				objectValue = ref
			}

			// Check if this value is from a function; if it is, parse the function
			_, fromFuncOk = v.Metadata["from_func"]
			if ok && objectValue != ref {
				asserted, ok := v.True.(token.Value)
				if ok {
					// Translate the function call
					funcCall, err := p.TranslateFunctionCall(asserted)
					if err != nil {
						return "", err
					}

					// might have to make a custom tName
					// [1:] to shave off the newline: remove if it causes problems
					objectString += tName + " = " + funcCall[1:]
					continue
				}
			}
		}

		// If it is an object or a struct inside the object, we need to make anonymize the name and p.Translate it
		if v.Type == "object" || v.Type == "struct" {
			vName := v.Name + "_" + RandStringBytesMaskImprSrc(10)
			anotherObjectString, err := p.TranslateObject(v, vName)
			if err != nil {
				return "", err
			}
			objectString += anotherObjectString + tName + "[\"" + v.Name + "\"] = " + vName + ";\n"

			// If it is a string, assign it normally
		} else if v.Type == "string" {
			objectString += tName + fmt.Sprintf("[\"%s\"] = \"%v\";\n", v.Name, objectValue)

			// If it is a var that is not from a function, then we need to copy the contents - again, no copy constructor in the runtime yet
		} else if v.Type == "var" && !fromFuncOk {
			var varStmt string
			var err error

			// Switch the real and acting types of the var and then p.Translate it using the existing translation mechanisms
			vName := v.Name
			vType := v.Type

			v.Type = v.Acting
			v.Acting = vType

			// If the var is an object or struct anonymize the name and p.Translate it
			if v.Type == "object" || v.Type == "struct" {
				vName = v.Name + "_" + RandStringBytesMaskImprSrc(10)
				varStmt, err = p.TranslateObject(v, vName)
				varStmt += tName + "[\"" + v.Name + "\"] = " + vName + ";\n"

				// If the shadow (acting) type of the var is a var, then we have an error.
				// The compiler MUST be able to identify the shadow type of any dynamic variable
				// at compile time in order to provide predictable behavior at runtime.
			} else if v.Type == "var" {
				fmt.Println("wtf still got var after switching acting and type")
				return "", errors.Errorf("Dynamic variable must have a static shadow type: %v", v)

				// Otherwise anonymize the name and just p.Translate the statement normally
			} else {
				vName = v.Name + "_" + RandStringBytesMaskImprSrc(10)
				varStmt, err = p.TranslateVariableStatement(v)
				if err != nil {
					return "", err
				}
			}

			// After translating, assign it to the object
			varStmt += tName + "[\"" + v.Name + "\"] = " + vName + ";\n"
			objectString += strings.Join(
				[]string{vType, vName, "=", strings.Join(strings.SplitAfter(varStmt, "=")[1:], "")}, " ",
			) + "\n"

			// If the object key is an array type
		} else if v.Type == "array" {
			// TODO: think about this when it comes up; Javascript/Python supports this feature in their dynamic variables
			// But we dont wanna be spewin memory everywhere and polluting the heap and destroying the ozone... wait wrong project
			// I am not supporting arrays for now, will have to debate how to
			// do this later. By definition, if objects are just map[string]<var>
			// and objects should be able to have keys with array values, then
			// <var> has to be able to containerize an array.
			// FIXME: the underlying C++ var could hold an array, but Express
			// could only allow its usage in arrays
			continue

			// If it is a zero valued float type make sure it has 0.0 and not 0; this is because the
			// runtime will initalize the variable's shadow type to be an int if we dont
		} else {
			if v.Type == token.FloatType && v.True.(float64) == 0 {
				objectString += tName + fmt.Sprintf("[\"%s\"] = %f;\n", v.Name, float64(0.0))
			} else {
				objectString += tName + fmt.Sprintf("[\"%s\"] = %v;\n", v.Name, objectValue)
			}
		}
	}

	return objectString, nil
}

// TODO: make this more idiomatic later
func (p *Parser) TranslateStruct(t token.Value) (string, error) {
	// call p.TranslateObject to get the object foot print
	// put that object into the map
	structDef, err := p.TranslateObject(t, "")
	if err != nil {
		return "", err
	}
	// structDef += "structMap[\"" + t.Name + "\"] = " + t.Name + ";\n"
	// structDef := "var " + t.Name + " = genStruct(\"" + t.Metadata["real"].(string) + "\");\n"

	return structDef, nil
	// return "var " + t.Name + " = genStruct(\"" + t.Acting + "\");\n", nil
}

func (p *Parser) TranslateVariableStatement(t token.Value) (string, error) {
	variableString := ""

	// if the token type is var make a var statement in C
	if insideLoop {
		if ref, ok := t.Metadata["refs"]; ok {
			t.True = ref
		}
	}

	tType := t.Type
	if _, ok := t.Metadata["assign"]; ok {
		tType = ""
	}

	fmt.Println("metadata", t.Metadata)
	if _, ok := t.Metadata["from_func"]; ok {
		funcCall, err := p.TranslateFunctionCall(t.True.(token.Value))
		if err != nil {
			return "", err
		}

		// might have to make a custom tName
		// [1:] to shave off the newline: remove if it causes problems
		if tType == "object" {
			tType = "var"
		}
		return tType + " " + t.Name + " = " + funcCall[1:], nil
	}

	switch t.Type {
	case "var":
		fmt.Println("around the world", t)
		var varStmt string
		var err error
		tName := t.Name

		tType = t.Type
		t.Type = t.Acting
		t.Acting = tType
		fmt.Println("t.Type", t.Type)
		fmt.Println("t.Acting", t.Acting)

		if t.Type == "object" || t.Type == "struct" {
			tName = t.Name + "_" + RandStringBytesMaskImprSrc(10)
			varStmt, err = p.TranslateObject(t, tName)
		} else if t.Type == "var" {
			fmt.Println("wtf still got var after switching acting and type")
			os.Exit(9)
		} else {
			fmt.Println("else", t.Type, t.Acting)
			varStmt, err = p.TranslateVariableStatement(t)
			if err != nil {
				return "", err
			}
		}

		_, ok := t.Metadata["assign"]
		if ok {
			_, ok = t.True.([]token.Value)
			fmt.Println("t, ok", t.Name, ok)
			if !ok {
				tType = ""
			}
		}

		variableString += strings.Join(
			[]string{tType, tName, "=", strings.Join(strings.SplitAfter(varStmt, "=")[1:], "")}, " ",
		) + "\n"
		return variableString, nil

	case "object":
		objectString, err := p.TranslateObject(t, "")
		if err != nil {
			return "", err
		}
		return variableString + objectString, nil

	case "struct":
		// real := t.Metadata["real"]
		var structString string
		var err error

		// struct declaration has no 'real' type
		// if real == nil {
		structString, err = p.TranslateStruct(t)
		// fmt.Println("structString", structString)
		if err != nil {
			return "", err
		}
		// } else {
		// 	// TODO: Just make structs essentially an object in the backend for now
		// 	// structString, err = p.TranslateObject(t)
		// 	// fmt.Println("structString2", structString)
		// 	// if err != nil {
		// 	// 	// TODO:
		// 	// 	return "", err
		// 	// }

		// 	// This was how I was doing the genStruct usage with the above in an if
		// 	structString = "var " + t.Name + " = genStruct(\"" + t.Metadata["real"].(string) + "\");\n"
		// 	objstr, err := p.TranslateObject(t, true)
		// 	if err != nil {
		// 		return "", err
		// 	}

		// 	structString += objstr

		// 	// TODO: Just make structs essentially an object in the backend for now
		// 	// if t.Metadata["refs"] != nil && t.Metadata["refs"] != "" {
		// 	// 	return variableString + structString, nil
		// 	// }
		// 	// structStrings += structString
		// 	// return variableString, nil
		// }
		return variableString + structString, nil

	case "array":
		arrayString, err := p.TranslateArray(t)
		if err != nil {
			// TODO:
			return "", err
		}

		return variableString + arrayString, nil

	// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
	// could just use that json library for now but wtf
	// //fmt.Println("std::map<string, " + +"> " + t.Name)
	case "string":
		variableString += "std::" + strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("\"%v\"", t.True)}, " ") + ";\n"
		// //fmt.Println(thing)
		// _, err = f.Write([]byte(thing))
		// if err != nil {
		// 	//fmt.Println("error writing to file")
		// 	os.Exit(9)
		// }
		// f += thing
		return variableString, nil

	case "int":
		fallthrough
	case "bool":
		fallthrough
	case "char":
		variableString += strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		return variableString, nil
	case "float":
		if t.True.(float64) == 0 {
			variableString += strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("%f", 0.0)}, " ") + ";\n"
		} else {
			variableString += strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		}
		// //fmt.Println(thing)
		// _, err = f.Write([]byte(thing))
		// if err != nil {
		// 	//fmt.Println("error writing to file")
		// 	os.Exit(9)
		// }
		// f += thing
		return variableString, nil

	default:
		//fmt.Println("am i an error ???")
		// return "", errors.Errorf("i am not nil %+v", t)

		// FIXME: This allowed us to compile when doing the accessor
		// shit but is pretty much a silent error
		return "", nil
	}
}

func (p *Parser) TranslateIf(t token.Value) (string, error) {
	controlString := ""

	if t.Type != token.If {
		return "", errors.New("blah")
	}

	//fmt.Println("wtf")
	//fmt.Printf("t %+v\n", t)

	// _, err = f.Write([]byte(fmt.Sprintf("if (%s) ", t.String)))
	// if err != nil {
	// 	//fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }

	// metadata, ok := t.Metadata
	// if !ok {
	// 	//fmt.Println("omfg error")
	// 	os.Exit(9)
	// }
	//fmt.Println("metadata", t.Metadata)

	//fmt.Println("t.True in p.Translateif", t)
	// evalString := ""
	// var leftInterface, opInterface, rightInterface interface{}
	// var left, op, right token.Value

	// if leftInterface = t.Metadata["left"]; leftInterface == nil {
	// 	return "", errors.New("Left was nil")
	// }
	// left = leftInterface.(token.Value)
	// evalString += left.String

	// if opInterface = t.Metadata["op"]; opInterface == nil {
	// 	return "", errors.New("Op was nil")
	// }
	// op = opInterface.(token.Value)
	// evalString += op.String

	// if rightInterface = t.Metadata["right"]; rightInterface == nil {
	// 	return "", errors.New("Right was nil")
	// }
	// right = rightInterface.(token.Value)
	// evalString += right.String

	blockString, err := p.TranslateBlock(token.Value{
		Type: token.Block,
		True: t.True,
	})
	if err != nil {
		// TODO:
		return "", err
	}

	controlString = fmt.Sprintf("if (%s) ", t.String) + blockString + "\n"

	return controlString, nil
}

func (p *Parser) TranslateLoop(t token.Value) (string, error) {
	loopString := ""

	// Turn on the loop var
	insideLoop = true
	// Turn off the loop var at the end
	defer func() {
		insideLoop = false
	}()

	//fmt.Printf("t: %+v", t)
	if t.Type != "for" {
		return "", errors.New("blah")
	}

	//fmt.Println()
	//fmt.Println("t.Metadata", t.Metadata)
	// os.Exit(9)

	// tValue, ok := t.Metadata)
	// if !ok {
	// 	return errors.New("not the type")
	// }
	// //fmt.Println("tValue", tValue)

	// for k, v := range t.Metadata {
	// 	//fmt.Println("k, v", k, v)
	// }

	loopString += "\n{\n"

	if extraVarsInterface, ok := t.Metadata["extraVars"]; ok {
		if extraVars, ok := extraVarsInterface.([]token.Value); ok {
			for _, extraVar := range extraVars {
				variableString, err := p.TranslateVariableStatement(extraVar)
				if err != nil {
					return "", err
				}

				loopString += variableString
			}
		}
	}

	// fmt.Println("endValue", t.Metadata["end"].(token.Value))
	// // end := t.Metadata["end"].(token.Value).True
	// endValue := t.Metadata["end"].(token.Value)
	// fmt.Println("endValue", endValue)
	// if _, ok := endValue.Metadata["refs"]; ok {
	// 	fmt.Println("wtf")
	// 	// end = endRef
	// 	os.Exit(9)
	// }

	loopString += fmt.Sprintf("int %s=%d;\nwhile (%s) {\n",
		t.Metadata["start"].(token.Value).Name,
		t.Metadata["start"].(token.Value).True.(int),
		t.String,
	)
	// //fmt.Println("loop", loop)
	// _, err = f.Write([]byte(loop))
	// if err != nil {
	// 	//fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	// f += loop

	//fmt.Println("wtf is this")
	blockString, err := p.TranslateBlock(token.Value{
		Type: token.Block,
		True: t.True,
	})
	if err != nil {
		// TODO:
		return "", err
	}

	loopString += blockString + fmt.Sprintf("%s+=%d;\n}\n}\n", t.Metadata["start"].(token.Value).Name, t.Metadata["step"].(token.Value).True.(int))
	// //fmt.Println(loopEnding)
	// _, err = f.Write([]byte(loopEnding))
	// if err != nil {
	// 	//fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	// f += loopEnding

	return loopString, nil
}

// FIXME: this will not work for lambda functions but we don't even have
// that supported yet so idc
func (p *Parser) TranslateOnExit(t token.Value) (string, error) {
	// TODO: look into what is the best way to do defer: might have to change whether we use:
	//		[=] - value
	//		[&] - reference
	onExitString := "onExitFuncs.deferStack.push([=](...){\n"

	onExitValue, ok := t.True.(token.Value)
	if !ok {
		// FIXME:
		return "", errors.Errorf("wtf happened %+v", t)
	}

	functionCallString, err := p.TranslateFunctionCall(onExitValue)

	onExitString += functionCallString + "});\n"
	return onExitString, err
}

// FIXME: this will not work for lambda functions but we don't even have
// that supported yet so idc
func (p *Parser) TranslateOnReturn(t token.Value) (string, error) {
	// TODO: look into what is the best way to do defer: might have to change whether we use:
	//		[=] - value
	//		[&] - reference
	onReturnString := "onReturnFuncs.deferStack.push([=](...){\n"

	onReturnValue, ok := t.True.(token.Value)
	if !ok {
		// FIXME:
		return "", errors.Errorf("wtf happened %+v", t)
	}

	functionCallString, err := p.TranslateFunctionCall(onReturnValue)

	onReturnString += functionCallString + "});\n"
	return onReturnString, err
}

// FIXME: this will not work for lambda functions but we don't even have
// that supported yet so idc
func (p *Parser) TranslateOnLeave(t token.Value) (string, error) {
	// TODO: look into what is the best way to do defer: might have to change whether we use:
	//		[=] - value
	//		[&] - reference
	onLeaveString := "onLeaveFuncs.deferStack.push([=](...){\n"

	onLeaveValue, ok := t.True.(token.Value)
	if !ok {
		// FIXME:
		return "", errors.Errorf("wtf happened %+v", t)
	}

	functionCallString, err := p.TranslateFunctionCall(onLeaveValue)

	onLeaveString += functionCallString + "});\n"
	return onLeaveString, err
}

// FIXME: this will not work for lambda functions but we don't even have
// that supported yet so idc
func (p *Parser) TranslateDefer(t token.Value) (string, error) {
	// TODO: look into what is the best way to do defer: might have to change whether we use:
	//		[=] - value
	//		[&] - reference
	stackName := "onExitFuncs"
	if blockDepth > 1 {
		stackName = "onReturnFuncs"
	}

	deferString := stackName + ".deferStack.push([=](...){\n"

	deferValue, ok := t.True.(token.Value)
	if !ok {
		// FIXME:
		return "", errors.Errorf("wtf happened %+v", t)
	}

	functionCallString, err := p.TranslateFunctionCall(deferValue)

	deferString += functionCallString + "});\n"
	return deferString, err
}

func (p *Parser) TranslateKeyword(t token.Value) (string, error) {
	if t.Type != token.Keyword {
		return "", errors.Errorf("not a keyword %+v", t)
	}

	switch t.String {
	case token.Return:
		return p.TranslateReturn(t)

	case token.OnExit:
		return p.TranslateOnExit(t)

	case token.OnReturn:
		return p.TranslateOnReturn(t)

	case token.OnLeave:
		return p.TranslateOnLeave(t)

	case token.Defer:
		return p.TranslateDefer(t)

	default:
		return "", errors.Errorf("idk wtf this is %+v", t)
	}
}

func (p *Parser) TranslateReturn(t token.Value) (string, error) {
	returnString := "return "

	returnValue, ok := t.True.(token.Value)
	if !ok {
		// FIXME:
		return "", errors.Errorf("wtf happened %+v", t)
	}

	if returnValue.Metadata["refs"] != nil {
		returnString += returnValue.Name + ";\n"
		// FIXME: do something about the char later
	} else if returnValue.Type == "string" || returnValue.Type == "char" {
		returnString += fmt.Sprintf("\"%+v\";\n", returnValue.True)
	} else if returnValue.Type == "BLOCK" || returnValue.Type == "object" || returnValue.Type == "var" {
		returnValue.Name = returnValue.Name + "_" + RandStringBytesMaskImprSrc(10)
		//fmt.Println("NAME_BYTES", returnValue.Name)

		objectString, err := p.TranslateObject(returnValue, "")
		if err != nil {
			return "", err
		}
		returnString = objectString + returnString + returnValue.Name + ";"
		// returnString += "{};\n"
	} else {
		returnString += fmt.Sprintf("%+v;\n", returnValue.True)
	}

	return returnString, nil
}

func (p *Parser) TranslateBlock(tv token.Value) (string, error) {
	blockDepth++
	defer func() { blockDepth-- }()
	// _, err = f.Write([]byte("{\n"))
	// if err != nil {
	// 	//fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	//fmt.Println("TVTVTVTV", tv)

	blockString := "\n"
	if blockDepth != 1 {
		blockString = "{" + blockString + "defer onLeaveFuncs;\n"
	}

	insideBlock, ok := tv.True.([]token.Value)
	if !ok {
		return "", errors.New("Could not assert block")
	}

	for _, t := range insideBlock {
		//fmt.Println("insideBlock t", t)
		blockString += "\n"

		var resultString string
		var err error

		switch t.Type {
		case "for":
			resultString, err = p.TranslateLoop(t)

		case "if":
			resultString, err = p.TranslateIf(t)

		case "function":
			switch t.Metadata["type"] {
			case "call":
				resultString, err = p.TranslateFunctionCall(t)

			case "def":
				var functionString string
				functionString, err = p.TranslateFunctionDef(t)
				if err != nil {
					return "", err
				}
				p.FunctionStrings += functionString
				continue

			default:
				return "", errors.Errorf("Not a function type: %s", t.Metadata["type"])
			}

		default:
			resultString, err = p.TranslateKeyword(t)
			if err != nil {
				resultString, err = p.TranslateVariableStatement(t)
			}
		}

		if err != nil {
			return "", err
		}

		blockString += resultString
	}

	if blockDepth != 1 {
		blockString += "\n}"
	}

	return blockString, nil
}

// genStructFactory was being used because in my head the grammar flow is:
// BLOCK + name = object
// object + name = struct (typed object)
// And therefore you can envision structs as just a type-safe object factory,
// like you have in JavaScript, but you don't have to use JavaSript...
// and it is incorporated into the language as a static compile time check to ensure type-safety
func genStructFactory() string {
	// structFactory :=
	return `
		std::map<std::string, var> structMap;
		var genStruct(std::string structName) {
			var structValue = structMap[structName];
			return structValue;
		}
		`

	// return structFactory
}

// Transpile starts the process of generating C++ code
func (p *Parser) Transpile(block token.Value) (string, error) {
	//fmt.Println("yo waddup")

	//fmt.Println("block", block)

	// //fmt.Println(p.source)

	// //fmt.Println("tokens", len(p.source))
	// for _, value := range  {
	// 	//fmt.Println()
	// 	fmt.Printf("value %+v\n", value)
	// }
	// f, err = os.Create("../test/output/cpp/main.expr.cpp")
	// if err != nil {
	// 	//fmt.Println("got an err creating file", err)
	// 	os.Exit(9)
	// }

	// TODO: check all f.Write errors I guess
	// f+="#include <map>\n#include <string>\n"
	// f+="struct Any { std::string type; void* data; };\n"
	r = rand.New(rand.NewSource(time.Now().Unix()))

	p.FunctionStrings = "\n"

	var f string

	baseLibs := []string{
		"#include <string>",
	}

	f += strings.Join(baseLibs, "\n")

	// Make the semantic parser declare these as it parses later
	// For now just include everything that we could ever need:
	//	- var
	//	- std
	//	- defer
	extraLibs := []string{
		"var.cpp",
		"std.cpp",
		"file.cpp",
		"defer.cpp",
	}

	for k := range extraLibs {
		extraLibs[k] = "#include " + strconv.Quote(LibBase+"/"+extraLibs[k])
	}

	f += "\n" + strings.Join(extraLibs, "\n")
	f += "\ndefer onExitFuncs;"

	blockString, err := p.TranslateBlock(block)
	if err != nil {
		// TODO:
		//fmt.Println("error getting block", err)
		return "", err
	}

	f += genStructFactory() + "\n" +
		structStrings + "\n" +
		p.FunctionStrings + "\nint main() {" +
		blockString + "}\n"

	return f, nil
}
