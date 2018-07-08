package lex_test

import (
	"encoding/json"
	"fmt"
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
	lexer, err := lex.NewFromFile("../samples/really_simple.expr")
	if err != nil {
		fmt.Println("NewFromFile", err)
		os.Exit(9)
	}

	lexTokens, err := lexer.Lex()
	if err != nil {
		fmt.Println("LexErr", err)
		os.Exit(9)
	}

	token.PrintTokens(lexTokens, "\t")
}
