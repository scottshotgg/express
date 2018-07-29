package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/parse"
	"github.com/scottshotgg/express/token"
)

var (
	jsonIndent = "\t"
)

func main() {
	argLen := len(os.Args)

	if argLen < 2 {
		fmt.Println("ERROR: You must provide an input program")
		return
	}

	input, err := ioutil.ReadFile(os.Args[argLen-1])
	if err != nil {
		fmt.Printf("ERROR: Cannot read input program: %s\n", os.Args[argLen-1])
		os.Exit(9)
	}

	l := lex.New(string(input))

	lexTokens, err := l.Lex()
	if err != nil {
		fmt.Println("error lexing", err)
		os.Exit(9)
	}
	token.PrintTokens(lexTokens, jsonIndent)
	fmt.Println("\n\n")

	p := parse.New(lexTokens)
	// tokens, err := p.Parse()
	// if err != nil {
	// 	fmt.Println("error in syntactic parsing", err)
	// 	os.Exit(9)
	// }
	// fmt.Println("tokens", tokens)
	// // PrintTokens(tokens, jsonIndent)
	// // fmt.Println("\n\n")

	// // p = parse.New(syntacticTokens)

	syntacticTokens, err := p.Syntactic()
	if err != nil {
		fmt.Println("error in syntactic parsing", err)
		os.Exit(9)
	}
	fmt.Println("Syntactic Tokens")
	token.PrintTokens(syntacticTokens, jsonIndent)
	fmt.Println("\n\n")

	pNew := parse.New(syntacticTokens)
	semanticTokens, err := pNew.Semantic()
	if err != nil {
		fmt.Println("error in semantic parsing", err)
		os.Exit(9)
	}
	fmt.Println("Semantic Tokens")
	// token.PrintValues(semanticTokens, jsonIndent)
	fmt.Println(semanticTokens)
	fmt.Println("\n\n")
}
