package parse

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/scottshotgg/ExpressRedo/token"
)

/*
	This should really be an interface that can be implemented with
	certain functions so that other transpilations can be implemeneted
*/

var (
	// f          string
	r          *rand.Rand
	err        error
	insideLoop bool

	// TODO: FIXME: this is causing functions to be compiled in every single file
	functionStrings = ""

	libBase = ""
)

func translateFunction(t token.Value) (string, error) {
	functionString := ""

	fmt.Printf("%+v\n", t)
	trueValue, ok := t.True.(map[string]token.Value)
	if !ok {
		fmt.Println("shit look at t")
		return "", errors.New("not ok")
	}

	// TODO: only supporting void type for now
	if trueValue["return"].True == nil {
		functionString += "void "
	}

	// Append the name of the function
	functionString += t.Name

	argsInterface := trueValue["args"].True
	if argsInterface == nil {
		functionString += "()"
	} else {
		args, ok := argsInterface.([]token.Token)
		if !ok {
			fmt.Println("shit look at args interface")
			return "", errors.New("not ok")
		}
		functionString += "("

		for i, arg := range args {
			functionString += arg.Type + " " + arg.Value.Name

			if i != len(args)-1 {
				functionString += ","
			}
		}
		functionString += ") "

	}

	fmt.Println("trueValue", trueValue)

	bodyString, err := translateBlock(trueValue["body"])
	if err != nil {
		// TODO:
	}

	functionString += bodyString

	return functionString, nil
}

func translateArray(t token.Value) (string, error) {
	arrayString := ""

	fmt.Printf("%+v\n", t)
	trueValue, ok := t.True.([]token.Token)
	if !ok {
		fmt.Println("shit look at t")
		fmt.Println("trueValue", trueValue)
		return "", errors.New("not ok")
	}

	// assuming only single type arrays until I have time to do multi type arrays in C
	arrayType := t.Acting

	if arrayType == "string" {
		arrayType = "std::" + arrayType
	} else if arrayType == "object" {
		arrayType = "map<string,var>"
	}

	arrayString += arrayType + " " + t.Name + "[] = { "

	for i, v := range trueValue {

		sprintString := "%v"
		// FIXME: for some reason array composition using only one object variable causes
		// each object to change when one does because of the pointers and stuff
		// It's because the memory is not copied (i.e, constructor is not called) when putting the var into the
		// map
		if ref, ok := v.Value.Metadata["refs"]; v.Value.Type != "object" && ok {
			arrayString += ref.(string)
		} else if v.Value.Type == "string" {
			sprintString = "\"" + sprintString + "\""
			arrayString += fmt.Sprintf(sprintString, v.Value.True)
		} else if v.Value.Type == "object" {
			arrayString += makeMap(v.Value)
		} else {
			arrayString += fmt.Sprintf(sprintString, v.Value.True)
		}

		if i != len(trueValue)-1 {
			arrayString += ", "
		}
	}

	arrayString += " };\n"

	return arrayString, nil
}

func makeMap(t token.Value) string {
	mapString := "map<string,var>{\n"
	for _, v := range t.True.([]token.Value) {
		// fmt.Println("k, v", k, v)
		if v.Type == "object" {
			mapString += fmt.Sprintf("{ \"%s\", %v },\n", v.Name, makeMap(v))
		} else if v.Type == "string" {
			mapString += fmt.Sprintf("{ \"%s\", \"%v\" },\n", v.Name, v.True)

		} else if v.Type == "array" {
			continue
		} else {
			mapString += fmt.Sprintf("{ \"%s\", %v },\n", v.Name, v.True)
		}
	}

	return mapString + "}"
}

// FIXME: just use var for now, but later we will try to not use var
// FIXME: ideally we want to store this in a "symbol map" with the defaults already there
func translateObject(t token.Value) (string, error) {
	return "var " + t.Name + " = " + makeMap(t) + "\n;", nil
}

func translateVariableStatement(t token.Value) (string, error) {
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

	fmt.Println("translating shit")

	switch t.Type {
	case "var":
		// int abc = 5;
		// Any zyx = Any{ "int", &abc };
		// varName := t.Name + strconv.Itoa(int(r.Uint32()))
		// variableString +=
		sprintfVar := "%v"
		if t.Acting == "string" {
			sprintfVar = "\"%v\""
		}
		variableString += strings.Join([]string{tType, t.Name, "=", fmt.Sprintf(sprintfVar, t.True)}, " ") + ";\n"

		// variableString += strings.Join([]string{t.Acting, varName, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		// variableString += "Any " + t.Name + " = Any{ \"" + t.Acting + "\", &" + varName + " };\n"
		// // fmt.Println(thing)
		// // _, err = f.Write([]byte(thing))
		// // if err != nil {
		// // 	fmt.Println("error writing to file")
		// // 	os.Exit(9)
		// // }
		return variableString, nil

	case "object":
		objectString, err := translateObject(t)
		if err != nil {
			// TODO:
			return "", err
		}
		return variableString + objectString, nil

	case "array":
		arrayString, err := translateArray(t)
		if err != nil {
			// TODO:
			return "", err
		}

		return variableString + arrayString, nil

	// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
	// could just use that json library for now but wtf
	// fmt.Println("std::map<string, " + +"> " + t.Name)
	case "string":
		variableString += "std::" + strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("\"%v\"", t.True)}, " ") + ";\n"
		// fmt.Println(thing)
		// _, err = f.Write([]byte(thing))
		// if err != nil {
		// 	fmt.Println("error writing to file")
		// 	os.Exit(9)
		// }
		// f += thing
		return variableString, nil

	case "int":
		fallthrough
	case "float":
		fallthrough
	case "bool":
		variableString += strings.Join([]string{tType, t.Name, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		// fmt.Println(thing)
		// _, err = f.Write([]byte(thing))
		// if err != nil {
		// 	fmt.Println("error writing to file")
		// 	os.Exit(9)
		// }
		// f += thing
		return variableString, nil

	default:
		fmt.Println("am i an error ???")
		return "", errors.New("i am not nil")
	}
	return "", errors.New("why am i here")
}

func translateIf(t token.Value) (string, error) {
	controlString := ""

	if t.Type != token.If {
		return "", errors.New("blah")
	}

	fmt.Println("wtf")
	fmt.Printf("t %+v\n", t)

	// _, err = f.Write([]byte(fmt.Sprintf("if (%s) ", t.String)))
	// if err != nil {
	// 	fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }

	// metadata, ok := t.Metadata
	// if !ok {
	// 	fmt.Println("omfg error")
	// 	os.Exit(9)
	// }
	fmt.Println("metadata", t.Metadata)

	fmt.Println("t.True", t.True)

	blockString, err := translateBlock(token.Value{
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

func translateLoop(t token.Value) (string, error) {
	loopString := ""

	// Turn on the loop var
	insideLoop = true
	// Turn off the loop var at the end
	defer func() {
		insideLoop = false
	}()

	fmt.Printf("t: %+v", t)
	if t.Type != "for" {
		return "", errors.New("blah")
	}

	fmt.Println()
	fmt.Println("t.Metadata", t.Metadata)
	// os.Exit(9)

	// tValue, ok := t.Metadata)
	// if !ok {
	// 	return errors.New("not the type")
	// }
	// fmt.Println("tValue", tValue)

	for k, v := range t.Metadata {
		fmt.Println("k, v", k, v)
	}

	loopString += "\n{\n"

	if extraVarsInterface, ok := t.Metadata["extraVars"]; ok {
		if extraVars, ok := extraVarsInterface.([]token.Value); ok {
			for _, extraVar := range extraVars {
				variableString, err := translateVariableStatement(extraVar)
				if err != nil {
					return "", err
				}

				loopString += variableString
			}
		}
	}

	loopString += fmt.Sprintf("int %s=%d;\nwhile (%s<%d) {\n",
		t.Metadata["start"].(token.Value).Name, t.Metadata["start"].(token.Value).True.(int), t.Metadata["start"].(token.Value).Name,
		t.Metadata["end"].(token.Value).True.(int))
	// fmt.Println("loop", loop)
	// _, err = f.Write([]byte(loop))
	// if err != nil {
	// 	fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	// f += loop

	fmt.Println("wtf is this")
	blockString, err := translateBlock(token.Value{
		Type: token.Block,
		True: t.True,
	})
	if err != nil {
		// TODO:
		return "", err
	}

	loopString += blockString + fmt.Sprintf("%s+=%d;\n}\n}\n", t.Metadata["start"].(token.Value).Name, t.Metadata["step"].(token.Value).True.(int))
	// fmt.Println(loopEnding)
	// _, err = f.Write([]byte(loopEnding))
	// if err != nil {
	// 	fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	// f += loopEnding

	return loopString, nil
}

func translateBlock(tv token.Value) (string, error) {
	// _, err = f.Write([]byte("{\n"))
	// if err != nil {
	// 	fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	fmt.Println("TVTVTVTV", tv)

	blockString := "{\n"

	insideBlock, ok := tv.True.([]token.Value)
	if !ok {
		return "", errors.New("Could not assert block")
	}

	for _, t := range insideBlock {
		fmt.Println("insideBlock t", t)

		variableString, err := translateVariableStatement(t)
		if err != nil {
			fmt.Println("i am here translateVariableStatement", err)
			loopString, err := translateLoop(t)
			if err != nil {
				fmt.Println("loop translation err", err)
				ifString, err := translateIf(t)
				if err != nil {
					// TODO:
					fmt.Println("if translate err", err)
					functionString, err := translateFunction(t)
					if err != nil {
						fmt.Println("function translate err", err)
						return "", err
					}
					functionStrings += functionString + "\n"
					continue
				}
				blockString += ifString
				continue
			}
			blockString += loopString
			continue
		}
		blockString += variableString
		continue
	}

	return blockString + "}\n", nil
}

func (p *Parser) Transpile(block token.Value) (string, error) {
	fmt.Println("yo waddup")

	fmt.Println("block", block)

	// fmt.Println(p.source)

	// fmt.Println("tokens", len(p.source))
	// for _, value := range  {
	// 	fmt.Println()
	// 	fmt.Printf("value %+v\n", value)
	// }
	// f, err = os.Create("../test/output/cpp/main.expr.cpp")
	// if err != nil {
	// 	fmt.Println("got an err creating file", err)
	// 	os.Exit(9)
	// }

	// TODO: check all f.Write errors I guess
	// f+="#include <map>\n#include <string>\n"
	// f+="struct Any { std::string type; void* data; };\n"
	r = rand.New(rand.NewSource(time.Now().Unix()))

	libBase, err = filepath.Abs("../lib/")
	if err != nil {
		os.Exit(9)
	}

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
		// "defer.cpp",
	}

	for k := range extraLibs {
		extraLibs[k] = "#include " + strconv.Quote(libBase+"/"+extraLibs[k])
	}

	f += "\n" + strings.Join(extraLibs, "\n")

	blockString, err := translateBlock(block)
	if err != nil {
		// TODO:
		fmt.Println("error getting block", err)
		return "", err
	}

	f += "\n" + functionStrings + "\nint main()" + blockString

	return f, nil
}
