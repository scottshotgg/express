package lex

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strconv"

// 	"github.com/golang-collections/collections/stack"
// )

// // TokenType ...
// type TokenType int

// const (
// 	TYPE TokenType = iota
// 	ASSIGN
// 	SPACE
// 	LIT
// )

// var (
// 	lexSymbols = map[string]TokenType{
// 		"int": TYPE,
// 		"=":   ASSIGN,
// 		" ":   SPACE,
// 	}
// )

// type Lexeme struct {
// 	Type string `json:"type,omitempty"`
// 	// ParseType parse.VariableType `json:"parseType,omitempty"`
// 	Value string      `json:"value,omitempty"`
// 	True  interface{} `json:"true,omitempty"`
// }

// // Lexer ...
// type Lexer struct {
// 	source      string
// 	Accumulator string
// 	Escaped     bool

// 	// We can just make our own stack later
// 	Enclosers *stack.Stack
// }

// // New ...
// func New(source string) *Lexer {
// 	return &Lexer{
// 		source:    source,
// 		Enclosers: stack.New(),
// 	}
// }

// func lexLiteral(accumulator string) Lexeme {
// 	// var err error

// 	// Make a token and set the default value to bool; this is just because its the
// 	// first case in the switch and everything below sets it, so it makes the code a bit
// 	// cleaner
// 	// We COULD do this with tokens in the tokenMap for true and false

// 	t := token.Token{
// 		ID:   0,
// 		Type: token.Literal,
// 		Value: token.Value{
// 			True: false,
// 			Type: token.BoolType,
// 		},
// 	}

// 	switch accumulator {
// 	// Default value is false, we only need to catch the case
// 	case "false":

// 	// Check if its true
// 	case "true":
// 		t.True = true

// 	// Else move on and figure out what kind of number it is (or an ident)
// 	default:
// 		t.Value = "int"
// 		var err error
// 		var tValue string

// 		if accumulator[:1] == "\"" && accumulator[len(accumulator)-1:] == "\"" {
// 			fmt.Println("I GOT HERE")
// 			return Lexeme{
// 				Type:  "lit",
// 				Value: "string",
// 				True:  accumulator[1 : len(accumulator)-1],
// 			}
// 		}

// 		// Figure out from the two starting characters
// 		if len(accumulator) > 2 {
// 			tValue = accumulator[2:]
// 			switch accumulator[:2] {
// 			// Binary
// 			case "0b":
// 				t.True, err = strconv.ParseInt(tValue, 2, 64)
// 				if err != nil {
// 					fmt.Println("ERROR", err)
// 				}

// 			// Octal
// 			case "0o":
// 				t.True, err = strconv.ParseInt(tValue, 8, 64)
// 				if err != nil {
// 					fmt.Println("ERROR", err)
// 				}

// 			// Hexadecimal
// 			case "0x":
// 				t.True, err = strconv.ParseInt(tValue, 16, 64)
// 				if err != nil {
// 					fmt.Println("ERROR", err)
// 				}
// 			}
// 		}
// 		if t.True != false {
// 			return t
// 		}

// 		// Attempt to parse an int from the accumulator
// 		t.True, err = strconv.ParseInt(accumulator, 0, 0)
// 		// TODO: need to make something for scientific notation with carrots and e
// 		// If it errors, check to see if it is an int
// 		if err != nil {
// 			// Attempt to parse a float from the accumulator
// 			t.True, err = strconv.ParseFloat(accumulator, 0)
// 			t.Value = "float"
// 			if err != nil {
// 				// leave this checking for the semantic
// 				// 	identSplit := strings.Split(accumulator, ".")
// 				// 	if len(identSplit) > 1 {
// 				// 		for _, ident := range identSplit {

// 				// 		}
// 				// 	}

// 				// need to check whether it is a type/keyword in the map
// 				// keyword, ok := token.TokenMap[accumulator]
// 				// if ok {
// 				// 	t = keyword
// 				// } else {
// 				// If it errors, assume that it is an ident (for now)

// 				t = Lexeme{
// 					Type:  "ident",
// 					Value: accumulator,
// 				}
// 			}
// 		}
// 	}

// 	return t
// }

// // Lex ...
// func (l *Lexer) Lex() (lexemes []Lexeme, err error) {
// 	file, err := os.Open("../lexemes.json")
// 	if err != nil {
// 		fmt.Println("error and stuff", err)
// 		os.Exit(9)
// 	}

// 	buf, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println("error and stuff", err)
// 		os.Exit(9)
// 	}

// 	lexMap := map[string]Lexeme{}
// 	err = json.Unmarshal(buf, &lexMap)
// 	if err != nil {
// 		fmt.Println("hey its me", err)
// 		os.Exit(9)
// 	}

// 	fmt.Printf("%+v\n", lexMap)

// 	for i, v := range l.source {
// 		fmt.Printf("i: %d, v: %s\n", i, string(v))
// 		fmt.Printf("accumulator: %s\n", l.Accumulator)

// 		// If we find that the current symbol is in the symbol map we need to decide what to do
// 		if tokenType, ok := lexMap[string(v)]; ok {
// 			fmt.Println("Found something in the symbol map", v, tokenType)

// 			switch string(v) {
// 			case "\"":
// 				fallthrough
// 			case "'":
// 				// TODO: need to check current stack peek
// 				// fallthrough
// 				if string(v) == l.Enclosers.Peek() && !l.Escaped {
// 					l.Enclosers.Pop()
// 					litLexeme := lexLiteral(l.Accumulator + string(v))
// 					lexemes = append(lexemes, litLexeme)
// 					l.Accumulator = ""

// 					continue
// 				} else {
// 					l.Enclosers.Push(string(v))
// 				}

// 			}

// 			// If we are enclosed by something, use the stack and figure out what to do
// 			if l.Enclosers.Len() != 0 {
// 				fmt.Println("stack peek", l.Enclosers.Peek())
// 				if l.Enclosers.Peek() == "\"" || l.Enclosers.Peek() == "'" {
// 					l.Accumulator += string(v)
// 				}

// 				continue
// 			}

// 			if len(l.Accumulator) > 0 {
// 				if accumulatorTokenType, ok := lexMap[l.Accumulator]; ok {
// 					fmt.Println("Found something in the symbol map again", l.Accumulator, accumulatorTokenType)

// 					lexemes = append(lexemes, Lexeme{
// 						Type:  accumulatorTokenType.Type,
// 						Value: l.Accumulator,
// 					})
// 				} else {
// 					// Not in the map but there is something there
// 					// Assume it is a literal
// 					litLexeme := lexLiteral(l.Accumulator)
// 					// if litLexeme.Value == "int" && string(v) == "." {
// 					// 	l.Accumulator += string(v)
// 					// 	continue
// 					// }
// 					lexemes = append(lexemes, litLexeme)
// 					l.Accumulator = ""
// 				}
// 			}
// 			l.Accumulator = ""

// 			lexemes = append(lexemes, Lexeme{
// 				Type:  tokenType.Type,
// 				Value: string(v),
// 			})

// 			// If its not in the symbol map
// 		} else {
// 			l.Accumulator += string(v)
// 		}
// 	}

// 	if len(l.Accumulator) > 0 {
// 		if accumulatorTokenType, ok := lexMap[l.Accumulator]; ok {
// 			fmt.Println("Found something in the symbol map again", l.Accumulator, accumulatorTokenType)

// 			lexemes = append(lexemes, Lexeme{
// 				Type:  accumulatorTokenType.Type,
// 				Value: l.Accumulator,
// 			})
// 			l.Accumulator = ""
// 		} else {
// 			// Not in the map but there is something there
// 			// Assume it is a literal
// 			litLexeme := lexLiteral(l.Accumulator)
// 			lexemes = append(lexemes, litLexeme)
// 			l.Accumulator = ""
// 		}
// 	}

// 	lexemes = append(lexemes, lexMap["end"])

// 	return
// }
