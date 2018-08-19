package libc

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// type ExpressType string
// type CType string

// // Express types
// const (
// 	INT    ExpressType = "int"
// 	STRING ExpressType = "string"
// 	BOOL   ExpressType = "bool"
// 	FLOAT  ExpressType = "float"
// 	CHAR   ExpressType = "char"
// 	STRUCT ExpressType = "struct"

// 	// Don't really want to use these, but we might have to
// 	VAR    ExpressType = "var"
// 	OBJECT ExpressType = "object"
// )

// // dont forget const

// char-> myCharPointer

// // C types
// const (
// 	VOID CType = "void"
// 	INT    CType = "int"
// 	STRING CType = "string"
// 	BOOL   CType = "bool"
// 	FLOAT  CType = "float"
// 	CHAR   CType = "char"
// 	STRUCT CType = "struct"

// 	VOIDP CType = "void *"
// 	INTP    CType = "int *"
// 	STRINGP CType = "string *"
// 	BOOLP   CType = "bool *"
// 	FLOATP  CType = "float *"
// 	CHARP   CType = "char *"
// 	STRUCTP CType = "struct *"

// 	SIZE_T CType = "size_t"
// 	SSIZE_T CType = "ssize_t"

// 	// Make a struct for this
// 	ERROR_T CType = "error_t"
// )

var (
	basicTypes = []string{"int", "float", "bool", "string"} //, "char", "struct", "enum"}

	// For now any function that contains a type that is not in the map will be dropped
	// This means:
	// - pointers
	// - any odd type
	// - unsigned/constant types
	// - size_t and ssize_t
	// - long x
	// - double
	// - complex
	// - struct: until we can figure out how to create them
	// - char
	// - enum
	// - anything with time
	// - "wint_t", "wctrans_t", "wchar_t"
	// - "void": will have to just translate this to nothing

	// For now leave const out
	// CType -> ExpressType
	typeMap = map[string]string{
		// // "const": "",
		// // "uint64_t": "int",
		// // "void": ""
		// "long": "int",
		// "long": "short",
		// "double": "float",
	}

	// otherType = []string{"wint_t", "wctrans_t", "wchar_t"}
	// wchar_t is i32
	// https://github.com/rust-lang/libc/blob/master/src/cloudabi/x86_64.rs
	// use rust docs as help

	// for _, t := range {

	// }

	cFuncs = []*CFunction{}
)

type CVariable struct {
	Type string
	Name string
}

type CFunction struct {
	CVariable

	Args []*CVariable
}

func ParseLibCFunctions() {
	for _, basicType := range basicTypes {
		typeMap[basicType] = basicType
	}

	fmt.Println("typeMap", typeMap)

	libcFunctions, err := ioutil.ReadFile("function")
	if err != nil {
		// TODO:
		os.Exit(9)
	}
	// fmt.Println("strings.Split(string(libcFunctions), \"\n\")", len(strings.Split(string(libcFunctions), "\n")))

	total := 0
	functions := strings.Split(string(libcFunctions), "\n")
	functions = []string{"int isalpha (int C)", "int isdigit (int C)"}

	var ok bool
	for _, function := range functions {
		cFunc := &CFunction{}

		funcSplit := strings.Split(function, " ")

		cFunc.Type, ok = typeMap[funcSplit[0]]

		// We know the first one is always a type so check it before we even go on
		if !ok {
			// For now we are just ignoring the ones that
			// don't have a compatible type defined in Express
			continue
		}

		fmt.Println(function)
		fmt.Println("funcSplit", funcSplit)
		fmt.Println("got type", typeMap[funcSplit[0]])

		// FIXME: need to make a var to track the index
		// For now we are ignoring pointers
		if funcSplit[1] == "*" {
			// Change cFunc.Type to be the pointer type
			//cFunc = funcSplit[2]
			continue
		} else {
			// If an asterisk isn't next, it has to be the name
			cFunc.Name = funcSplit[1]
			fmt.Println("got name function", funcSplit[1])
		}

		dont := false
		// TODO: we will need to check this index
		// Loop over the arguments now until the last one
		// These 2's need to be tha same index as above
		for i := 2; i < len(funcSplit[2:]); i++ {
			cVar := &CVariable{}
			fmt.Println("i brah", funcSplit[i]) // Loop over everything until you get to something that doesn't have
			// a comma inside the string indicating it is the last piece and is the name
			for strings.Count(funcSplit[i], ",") == 0 || strings.Count(funcSplit[i], ")") == 0 {
				fmt.Println("innerType", funcSplit[i])
				// We will have to do some extra checking for struct later
				expressType, ok := typeMap[strings.Replace(funcSplit[i], "(", "", -1)]
				if !ok {
					// For now we are just ignoring the ones that
					// don't have a compatible type defined in Express
					dont = true
					break
				}
				// Keep appending what it can; 'unsigned', 'constant', etc
				fmt.Println("got a type2", funcSplit[i])
				cVar.Type += expressType

				i++
			}

			fmt.Println("got to the last one", funcSplit[i])

			// It looks like for the args and returns, they are including the asterisks
			// as part of the variable name and not the type declaration
			if !dont && strings.Count(funcSplit[i], "*") > 0 {
				dont = true
				break
			}

			cVar.Name = funcSplit[i]
			fmt.Println("got the last ones name", funcSplit[i])
			i++

			cFunc.Args = append(cFunc.Args, cVar)
			fmt.Println()
		}

		if dont {
			continue
		}

		cFuncs = append(cFuncs, cFunc)
		total++
	}

	fmt.Println("total:", total)
	fmt.Printf("%+v\n", *cFuncs[0])
}

// RipLibC takes a libc.info and rips all relevant parts and filenames out of it
// TODO: FIXME: might just be better and easier to rip glibc all together
func RipLibC() {
	libcContents, err := ioutil.ReadFile("libc.info")
	if err != nil {
		// TODO:
		os.Exit(9)
	}

	fmt.Println("len(libcContents)", len(libcContents))

	// accumulation := []byte{}
	// capture := []byte{}

	// for _, char := range libcContents {

	// }

	// zp := regexp.MustCompile("`(.*\.h)`")

	// Split the file everywhere we see that phrase
	// String split max 1 by "." and capture the filename
	// In the rest of the text capture everything that is like '--'
	// Also get is defined in FILENAME from the last period

	// The following functions are declared in the header file `search.h`.
	// The following xyzs are declared in the header file `abc.h`.

	declaredInSplit := strings.Split(string(libcContents), "declared in the header file")
	fmt.Println(len(declaredInSplit))

	// fmt.Println(declaredInSplit[1])
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	filenameSplit := strings.SplitN(declaredInSplit[1], "`.", 2)

	// Shave off the preceding backtick left by the regex
	filename := filenameSplit[0][2:]
	fmt.Println("Filename:", filename)

	// -- Variables, -- Data Types, etc
	dashdash := regexp.MustCompile(`-- [A-z]:*.*\n`)

	// fmt.Println(filenameSplit[1][0:50])

	// Find all dashdash items
	joinedDashDashString := strings.Join(func(shit [][]byte) (notshit []string) {
		for _, k := range shit {
			notshit = append(notshit, string(k))
		}

		return
	}(dashdash.FindAll([]byte(filenameSplit[1]), -1)), "")

	fmt.Println(joinedDashDashString)

	// for _, declared := range declaredInSplit {
	// }

	// -- [A-z]*.*
	// \n\n`[A-Z]*`
}
