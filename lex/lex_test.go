package lex_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/scottshotgg/ExpressRedo/lex"
	"github.com/scottshotgg/ExpressRedo/token"
)

var (
	l *lex.Lexer
	// TODO: one thing the old lexer architecture fixed was the space at the end
	simpleTest = "int i = 5 "
)

func TestNew(t *testing.T) {
	fmt.Println("TestNew")

	l = lex.New(simpleTest)
	fmt.Printf("Lexer: %+v\n", l)
}

func TestLex(t *testing.T) {
	TestNew(t)

	fmt.Println("TestLex")
	lexemes, err := l.Lex()
	if err != nil {
		// TODO:
	}

	// fmt.Printf("lexemes: %+v\n", lexemes)
	for i, lexeme := range lexemes {
		fmt.Println("i", i, lexeme)
	}
	lexemeJSON, err := json.Marshal(lexemes)
	if err != nil {
		fmt.Println("jsonErr", err)
	}
	fmt.Println(string(lexemeJSON))
}

func TestNewFromFile(t *testing.T) {
	f, err := os.Open("../samples/if.expr")
	if err != nil {
		fmt.Println("openErr", err)
		os.Exit(9)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("openErr", err)
		os.Exit(9)
	}

	lexTokens, err := lex.New(string(data)).Lex()
	if err != nil {
		fmt.Println("lexErr", err)
		os.Exit(9)
	}

	token.PrintTokens(lexTokens, "\t")
}
