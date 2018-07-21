package parse

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

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
)

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
	}

	arrayString += arrayType + " " + t.Name + "[] = { "

	for i, v := range trueValue {
		sprintString := "%v"
		if v.Value.Type == "string" {
			sprintString = "\"" + sprintString + "\""
		}
		arrayString += fmt.Sprintf(sprintString, v.Value.True)
		if i != len(trueValue)-1 {
			arrayString += ", "
		}
	}

	arrayString += " };\n"

	return arrayString, nil
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

	switch t.Type {
	case "var":
		// int abc = 5;
		// Any zyx = Any{ "int", &abc };
		varName := t.Name + strconv.Itoa(int(r.Uint32()))
		variableString += strings.Join([]string{t.Acting, varName, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		variableString += "Any " + t.Name + " = Any{ \"" + t.Acting + "\", &" + varName + " };\n"
		// fmt.Println(thing)
		// _, err = f.Write([]byte(thing))
		// if err != nil {
		// 	fmt.Println("error writing to file")
		// 	os.Exit(9)
		// }
		return variableString, nil

	case "object":
		// translateObject(t)
		return variableString, nil

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

	fmt.Println("wtf")
	fmt.Printf("t %+v\n", t)

	// _, err = f.Write([]byte(fmt.Sprintf("if (%s) ", t.String)))
	// if err != nil {
	// 	fmt.Println("error writing to file")
	// 	os.Exit(9)
	// }
	controlString += fmt.Sprintf("if (%s) ", t.String)

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

	controlString += blockString + "\n"

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

	loopString += fmt.Sprintf("{\nint %s=%d;\nwhile (%s<%d) {\n",
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

	loopString += blockString + fmt.Sprintf("%s+=%d;}\n}\n", t.Metadata["start"].(token.Value).Name, t.Metadata["step"].(token.Value).True.(int))
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
	blockString := "{\n"

	insideBlock, ok := tv.True.([]token.Value)
	if !ok {
		return "", errors.New("Could not assert block")
	}
	fmt.Println("insideBlock", len(insideBlock))

	for _, t := range insideBlock {
		fmt.Println("insideBlock t", t)

		variableString, err := translateVariableStatement(t)
		if err != nil {
			fmt.Println("i am here translateVariableStatement", err)
			loopString, err := translateLoop(t)
			if err != nil {
				ifString, err := translateIf(t)
				if err != nil {
					// TODO:
				}
				return blockString + ifString + "}", nil
			}
			return blockString + loopString + "}", nil
		}
		return blockString + variableString + "}", nil
	}

	return "", errors.New("something")
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
	var f string

	f += "#include <string>\n"
	f += "#include \"json.hpp\"\n"
	f += "int main()"

	blockString, err := translateBlock(block)
	if err != nil {
		// TODO:
		fmt.Println("error getting block", err)
		return "", err
	}

	f += blockString

	return f, nil
}
