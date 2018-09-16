package lex

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/golang-collections/collections/stack"
	"github.com/scottshotgg/express/token"
)

// TokenType ...
type TokenType int

const (
	TYPE TokenType = iota
	ASSIGN
	SPACE
	LIT
)

var (
	lexSymbols = map[string]TokenType{
		"int": TYPE,
		"=":   ASSIGN,
		" ":   SPACE,
	}
)

type Lexeme struct {
	Type string `json:"type,omitempty"`
	// ParseType parse.VariableType `json:"parseType,omitempty"`
	Value string      `json:"value,omitempty"`
	True  interface{} `json:"true,omitempty"`
}

// Lexer ...
type Lexer struct {
	source      string
	Accumulator string
	Escaped     bool
	Tokens      []token.Token
	LastToken   token.Token

	// We can just make our own stack later
	Enclosers *stack.Stack
}

func New(source string) *Lexer {
	return &Lexer{
		source:    source,
		Enclosers: stack.New(),
	}
}

func NewFromFile(path string) (*Lexer, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("openErr", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("openErr", err)
		return nil, err
	}

	return New(string(data)), nil
}

func (meta *Lexer) LexLiteral() token.Token {
	// var err error

	// Make a token and set the default value to bool; this is just because its the
	// first case in the switch and everything below sets it, so it makes the code a bit
	// cleaner
	// We COULD do this with tokens in the tokenMap for true and false
	t := token.Token{
		ID:   0,
		Type: token.Literal,
		Value: token.Value{
			True: false,
			Type: token.BoolType,
		},
	}

	switch meta.Accumulator {
	// Default value is false, we only need to catch the case
	case "false":

	// Check if its true
	case "true":
		t.Value.True = true

	// Else move on and figure out what kind of number it is (or an ident)
	default:
		// Figure out from the two starting characters
		if len(meta.Accumulator) > 2 {
			t.Value.String = meta.Accumulator[2:]
			switch meta.Accumulator[:2] {
			// Binary
			case "0b":
				value, err := strconv.ParseInt(t.Value.String, 2, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "binary"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t

			// Octal
			case "0o":
				value, err := strconv.ParseInt(t.Value.String, 8, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "octal"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t

			// Hexadecimal
			case "0x":
				value, err := strconv.ParseInt(t.Value.String, 16, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "hexadecimal"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t
			}
		}
		// Clear the string value
		t.Value.String = ""

		// Attempt to parse an int from the accumulator
		value, err := strconv.ParseInt(meta.Accumulator, 0, 0)
		if err != nil {
			// TODO:
		}
		t.Value.True = int(value)
		t.Value.Type = token.IntType

		// TODO: need to make something for scientific notation with carrots and e
		// If it errors, check to see if it is an int
		if err != nil {
			// Attempt to parse a float from the accumulator
			t.Value.True, err = strconv.ParseFloat(meta.Accumulator, 0)
			t.Value.Type = token.FloatType
			if err != nil {
				// leave this checking for the semantic
				// 	identSplit := strings.Split(meta.Accumulator, ".")
				// 	if len(identSplit) > 1 {
				// 		for _, ident := range identSplit {

				// 		}
				// 	}

				// need to check whether it is a type/keyword in the map
				keyword, ok := token.TokenMap[meta.Accumulator]
				if ok {
					t = keyword
				} else {
					// If it errors, assume that it is an ident (for now)
					t.Type = token.Ident
					t.Value = token.Value{
						String: meta.Accumulator,
					}
				}
			}
		}
	}

	return t
}

// Lex attemps to lex the token
func (meta *Lexer) Lex() ([]token.Token, error) {
	for index := 0; index < len(meta.source); index++ {
		char := meta.source[index]
		if string(char) == " " || string(char) == "\n" {
			if meta.Accumulator != "" {
				if lexemeToken, ok := token.LexemeMap[meta.Accumulator]; ok {
					meta.Tokens = append(meta.Tokens, lexemeToken)
				} else {
					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
				}
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			} else if string(char) == " " || string(char) == "\n" {
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}

			continue

		} else {
			if lexemeToken, ok := token.LexemeMap[string(char)]; ok {
				// Filter out the comments
				switch lexemeToken.Value.Type {
				case "div":
					index++
					if index < len(meta.source)-1 {
						switch meta.source[index] {
						case '/':
							for {
								index++
								if index == len(meta.source) || meta.source[index] == '\n' {
									break
								}
							}

						case '*':
							for {
								index++
								if index == len(meta.source) || (meta.source[index] == '*' && meta.source[index+1] == '/') {
									index++
									break
								}
							}

						default:
							meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
						}
					}
					continue

				// Use the lexer to parse strings
				case "squote":
					fmt.Println("found an squote")
					fallthrough

				case "dquote":
					stringLiteral := ""

					index++
					for string(meta.source[index]) != lexemeToken.Value.String {
						stringLiteral += string(meta.source[index])
						index++
					}

					varType := token.StringType
					if len(stringLiteral) < 2 {
						varType = token.CharType
					}

					meta.Tokens = append(meta.Tokens, token.Token{
						ID:   0,
						Type: token.Literal,
						Value: token.Value{
							Type:   varType,
							True:   stringLiteral,
							String: stringLiteral,
						},
					})

					continue
				}

				if meta.Accumulator != "" {
					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
					meta.Accumulator = ""
				}

				meta.Tokens = append(meta.Tokens, lexemeToken)
				meta.Accumulator = ""

				continue
			} else if string(char) == " " || string(char) == "\n" {
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
		}

		meta.Accumulator += string(char)

		if char == 0 {
			break
		}
	}

	return meta.Tokens, nil
}
