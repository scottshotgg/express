package libc

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type RustStruct struct {
}

type RustVariable struct {
	Type    string
	Name    string
	Pointer int
}

type RustFunction struct {
	RustString string
	ReturnType string
	Name       string
	Args       []RustVariable
}

func New(function string) RustFunction {
	return RustFunction{
		RustString: function,
	}
}

func (rf *RustFunction) Rust() string {
	return rf.RustString
}

func (rf *RustFunction) C() string {
	return fmt.Sprintf("%s %s(%s)",
		rf.ReturnType, rf.Name, func(args []RustVariable) (vars string) {
			i := 0
			for ; i < len(args)-1; i++ {
				vars += args[i].Type + " " + args[i].Name + ", "
			}
			vars += args[i].Type + " " + args[i].Name
			return
		}(rf.Args))
}

func (rf *RustFunction) Express() string {
	// TODO: for now we have to put the parens there
	return fmt.Sprintf("func %s(%s) (%s)",
		rf.Name, func(args []RustVariable) (vars string) {
			i := 0
			for ; i < len(args)-1; i++ {
				vars += args[i].Type + " " + args[i].Name + ", "
			}
			vars += args[i].Type + " " + args[i].Name
			return
		}(rf.Args), rf.ReturnType)
}

func ParseRustFunction() {
	// function := `pub unsafe extern "C" fn aio_read(aiocbp: *mut aiocb) -> c_int`
	// Test const with this
	function := `pub fn aio_suspend(aiocb_list: *const *const aiocb, nitems: ::c_int, timeout: *const ::timespec) -> ::c_int;`
	// function :=
	// 	`pub fn process_vm_writev(pid: ::pid_t,
	// 	local_iov: *const ::iovec,
	// 	liovcnt: ::c_ulong,
	// 	remote_iov: *const ::iovec,
	// 	riovcnt: ::c_ulong,
	// 	flags: ::c_ulong) -> isize;`

	rustFunc := New(function)

	fmt.Println("original:", function)

	// Sift out the \n, ::, const and mut for now
	function = strings.Replace(function[0:len(function)-1], "\n", "", -1)
	function = strings.Replace(function, "::", "", -1)
	function = strings.Replace(function, "const ", "", -1)
	function = strings.Replace(function, "c_", "", -1)
	function = strings.Replace(function, "mut ", "", -1)

	whitespaceRegex := regexp.MustCompile(`,([\n|\s])*`)
	functionByte := whitespaceRegex.ReplaceAll([]byte(function), []byte(","))
	function = string(functionByte)
	fmt.Println("function after everything:", function)

	// Parse the pub later if we need to
	relevantFnInfo := strings.SplitAfter(function, "fn ")

	// Need to check if it is more than one
	fmt.Println("relevantFnInfo[1]:", relevantFnInfo[1])

	// split by "->"; this will give us the return type
	returnType := strings.Split(relevantFnInfo[1], " -> ")
	fmt.Println("splitReturn:", returnType[1])

	rustFunc.ReturnType = returnType[1]

	// fmt.Println(returnType[0])

	// parenRegex := regexp.MustCompile("([A-z]*)")
	// parenRegexSplit := parenRegex.FindAll([]byte(returnType[0]), -1)

	// rustFunc.Name = parenRegexSplit[0]

	// for _, match := range parenRegexSplit[] {
	// 	if string(match) != "" {
	// 		fmt.Println(string(match))
	// 	} else if strings.Contains(string(match), "mut") {
	// 		switch strings.Count(string(match), "*") {
	// 		case 0:
	// 			// Non-pointer

	// 		case 1:
	// 			// How can we handle pointers?

	// 		default:
	// 			fmt.Println("had more than two asterisks")
	// 		}
	// 	}
	// }

	// Split by '('
	lParenSplit := strings.Split(returnType[0], "(")
	// [0] -> function name
	// [1] -> args
	rustFunc.Name = lParenSplit[0]
	fmt.Println("function name:", lParenSplit[0])

	args := strings.Split(strings.Replace(lParenSplit[1], ")", "", 1), ",")
	fmt.Println("args", len(args), args)

	for _, arg := range args {
		argSplit := strings.Split(arg, ": ")
		fmt.Println("argSplit:", argSplit)

		// TODO: Make sure that the type is actually defined here
		rustFunc.Args = append(rustFunc.Args, RustVariable{
			Name:    argSplit[0],
			Type:    strings.Replace(argSplit[1], "*", "", -1),
			Pointer: strings.Count(argSplit[1], "*"),
		})
	}

	fmt.Println("rustFunc", rustFunc)
	eJSON, _ := json.Marshal(rustFunc)
	fmt.Println(string(eJSON))

	fmt.Println(rustFunc.C())
	fmt.Println(rustFunc.Express())

	// fmt.Println("parenRegexSplit", parenRegexSplit)
}

// func ParseRustStruct() {

// }
