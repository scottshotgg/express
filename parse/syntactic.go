package parse

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/scottshotgg/express/token"
)

// ParseString parses a string literal. Anything surrounded by quotes.
func (p *Parser) ParseString() token.Token {
	stringLiteral := ""
	for {
		p.ShiftWithWS()
		//fmt.Println("current", p.CurrentToken)

		// FIXME: stop doing hacky shit, purge this shit, need to preserve whitespaces in the lexer
		stringLiteral += p.CurrentToken.Value.String
		if p.NextToken.Value.String == "\"" {

			p.ShiftWithWS()

			return token.Token{
				Type: token.Literal,
				Value: token.Value{
					Type:   "string",
					True:   stringLiteral,
					String: stringLiteral,
				},
			}
		}
		// Getting the last 'separating' character; aka a whitespace that was separating the tokens
	}
}

// ParseGroup parses a grouping of items; tuple, function arguments, function returns. Anything encapsulated in parenthesis.
func (p *Parser) ParseGroup() token.Token {
	groupTokens := []token.Token{}

	for {
		p.Shift()

		current := p.CurrentToken

		switch current.Type {
		case token.RParen:
			return token.Token{
				ID:   1,
				Type: token.Group,
				Value: token.Value{
					Type: token.Group,
					True: groupTokens,
				},
			}

		case token.PriOp:
			fallthrough
		case token.SecOp:
			fallthrough
		case token.Literal:
			groupTokens = append(groupTokens, current)

		case token.Type:
			peek := p.NextToken
			switch peek.Type {
			case token.LBracket:
				//fmt.Println("found array")
				os.Exit(8)

			case token.Ident:
				groupTokens = append(groupTokens, p.CurrentToken)
				p.Shift()
				//fmt.Println("GROUP PARSE IDENT")
				p.ParseIdent(&groupTokens, p.CurrentToken)

			case token.Literal:
				groupTokens = append(groupTokens, p.CurrentToken)

				p.Shift()
				p.CurrentToken.Type = token.Ident
				groupTokens = append(groupTokens, p.CurrentToken)
			default:
				os.Exit(7)
			}

		case token.Ident:
			p.ParseIdent(&groupTokens, p.CurrentToken)

		case token.Separator:
			continue

		case token.DQuote:
			groupTokens = append(groupTokens, p.ParseString())

		case token.LBrace:
			groupTokens = append(groupTokens, p.ParseBlock())

		case token.LBracket:
			groupTokens = append(groupTokens, p.ParseArray())

		// New grouping
		case token.LParen:
			groupTokens = append(groupTokens, p.ParseGroup())

		case token.Block:
			groupTokens = append(groupTokens, current)

		case token.Group:
			//fmt.Println("CURRENT SHIT", current)
			currentTrue := current.Value.True.([]token.Token)
			//fmt.Println("CURRENT STUFFS", currentTrue)
			newCurrent := append(
				[]token.Token{token.Token{
					Type: token.LParen,
					Value: token.Value{
						Type: "op_3",
					},
				}},
				append(currentTrue,
					token.Token{
						Type: token.RParen,
						Value: token.Value{
							Type: "op_3",
						},
					})...)
			//fmt.Println("NEW CURRENT", newCurrent)
			if currentTrue != nil {
				newParser := New(newCurrent)
				newParser.Shift()
				//fmt.Println("newParser.LastToken", p.LastToken)
				//fmt.Println(newParser.CurrentToken)
				//fmt.Println(newParser.NextToken)

				iterationCurrent := newParser.ParseGroup()

				//fmt.Println("ITERATION CURRENT", iterationCurrent)

				groupTokens = append(groupTokens, iterationCurrent.Value.True.([]token.Token)...)
			} else {
				//fmt.Println("ELSE CURRENT", current)
				groupTokens = append(groupTokens, current)
			}

		default:
			//fmt.Printf("ERROR: Unrecognized group token; current: %+v\n meta: %+v\n\n", current, p)
			//fmt.Println("p.LastToken", p.LastToken)
			//fmt.Println(p.CurrentToken)
			//fmt.Println(p.NextToken)
			os.Exit(8)
		}
	}
}

// ParseArray parses an array of items. Anything encapulated in square brackets except for attributes.
func (p *Parser) ParseArray() token.Token {
	arrayTokens := []token.Token{}

	for {
		//fmt.Println("arrayTokens", arrayTokens)
		p.Shift()

		switch p.CurrentToken.Type {
		case token.Separator:
			//fmt.Println("found separator")
			continue

		case token.Ident:
			p.ParseIdent(&arrayTokens, p.CurrentToken)

		case token.DQuote:
			arrayTokens = append(arrayTokens, p.ParseString())

		case token.Literal:
			arrayTokens = append(arrayTokens, p.CurrentToken)

		case token.LParen:
			arrayTokens = append(arrayTokens, p.ParseGroup())

		case token.LBrace:
			arrayTokens = append(arrayTokens, p.ParseBlock())

		case token.LBracket:
			// arrayTokens = append(arrayTokens, p.ParseArray())

		case token.RBracket:
			//fmt.Println("arrayTokens2", arrayTokens)
			return token.Token{
				ID:   1,
				Type: token.Array,
				Value: token.Value{
					Type: token.ArrayType,
					True: arrayTokens,
				},
			}

		case token.Assign:
			arrayTokens = append(arrayTokens, p.CurrentToken)

		case token.SecOp:
			arrayTokens = append(arrayTokens, p.CurrentToken)

		case "":
			//fmt.Println("we got nothing", arrayTokens)
			// return arrayTokens,
			os.Exit(9)

		default:
			//fmt.Println("ERROR: Unrecognized array token", p.CurrentToken, p)
			os.Exit(8)
		}

		// // FIXME: This should throw an error
		// if p.NextToken == (token.Token{}) {
		// 	//fmt.Println("nextToken array", arrayTokens)
		// 	return token.Token{}
		// }
	}
}

func (p *Parser) ParseNamedFuncDef() token.Token {
	// Append the ident token
	nameToken := p.NextToken
	functionTokens := []token.Token{}
	p.Shift()
	//fmt.Println("wowee", p.NextToken)
	p.Shift()
	args := p.ParseGroup()
	functionTokens = append(functionTokens, args)
	//fmt.Println("wowe2e", p.NextToken)
	p.Shift()
	returns := p.ParseGroup()
	functionTokens = append(functionTokens, returns)
	//fmt.Println("args", args)
	//fmt.Println("wowee3", p.NextToken)
	//fmt.Println("returns", returns)
	p.Shift()
	body := p.ParseBlock()
	functionTokens = append(functionTokens, body)
	// bodyTokens := body.Value.True.([]token.Token)
	// //fmt.Println("bodyTokens", bodyTokens)
	// //fmt.Println("body", body)

	// if p.NextToken.Type == token.Group {
	// 	//fmt.Println("woah we hit a group")
	// 	os.Exit(9)
	// }

	// FIXME: fix this later
	if len(nameToken.Value.String) < 1 {
		//fmt.Println("wtf that name")
		os.Exit(9)
	}

	accessType := token.PrivateAccessType
	if nameToken.Value.String[0] > 64 || nameToken.Value.String[0] < 91 {
		accessType = token.PublicAccessType
	}

	return token.Token{
		ID:   1,
		Type: token.Function,
		Value: token.Value{
			Name:       nameToken.Value.String,
			AccessType: accessType,
			Type:       token.FunctionType,
			// True: map[string]token.Token{
			// 	"args":    args,
			// 	"returns": returns,
			// 	"body":    body,
			// },
			True: functionTokens,
			Metadata: map[string]interface{}{
				"lambda": false,
				"type":   "def",
			},
		},
	}
}

func (p *Parser) ParseFunctionDef() token.Token {
	next := p.NextToken

	//fmt.Println("woah a function def", next)

	switch next.Type {
	case token.Ident:
		// Named function case
		return p.ParseNamedFuncDef()

	case token.LParen:
		// Lambda function case

	default:
		//fmt.Println("did not find an L_PAREN")
		os.Exit(9)
	}

	return token.Token{
		Type: "BROKEN",
	}
}

func (p *Parser) ParseFunctionCall() token.Token {
	// next := p.NextToken

	//fmt.Println("woah i got a function call", p.CurrentToken)

	// Append the ident token
	nameToken := p.CurrentToken
	functionTokens := []token.Token{}
	//fmt.Println("wowee", p.NextToken)
	p.Shift()
	args := p.ParseGroup()
	functionTokens = append(functionTokens, args)
	//fmt.Println("wowe2e", p.NextToken)

	// FIXME: fix this later
	if len(nameToken.Value.String) < 1 {
		//fmt.Println("wtf that name")
		os.Exit(9)
	}

	accessType := token.PrivateAccessType
	if nameToken.Value.String[0] > 64 || nameToken.Value.String[0] < 91 {
		accessType = token.PublicAccessType
	}

	ft := token.Token{
		ID:   1,
		Type: token.Function,
		Value: token.Value{
			Name:       nameToken.Value.String,
			AccessType: accessType,
			Type:       token.FunctionType,
			True:       functionTokens,
			Metadata: map[string]interface{}{
				"lambda": false,
				"type":   "call",
			},
		},
	}

	//fmt.Println("ft", ft)
	return ft
}

// ParseIdent parses an identifier
func (p *Parser) ParseIdent(blockTokens *[]token.Token, peek token.Token) {
	if blockTokens == nil {
		fmt.Println("ERROR: blockTokens is nil")
		os.Exit(5)
	}

	identTokens := []token.Token{}

	identSplit := strings.Split(peek.Value.String, ".")
	// TODO: change this to a traditional for to right before the length
	for i, ident := range identSplit {
		identTokens = append(identTokens, token.Token{
			ID:   0,
			Type: token.Ident,
			// Expected:
			Value: token.Value{
				Type: func() string {
					if len(ident) > 0 && ident[0] > 64 && ident[0] < 91 {
						return "public"
					}

					return "private"
				}(),
				// True: ,
				String: ident,
			},
		})

		if i < len(identSplit)-1 {
			identTokens = append(identTokens, token.TokenMap["."])
		}
	}

	// If we have an `ident` and then a brace (starting of a `block`)
	if p.NextToken.Type == token.LBrace {
		// Once we get here we need to determine if this is an `ident`
		// and then an anonymous scope (i.e, `block`) or is this a `struct`

		// Test this by unshifting and then shifting with whitespace
		p.Unshift()
		p.ShiftWithWS()

		// If there is no newline separating the `ident` and `block`,
		// then this is a `struct`
		if p.NextToken.Value.Type != "newline" { // && p.NextToken.Value.Type != "BLOCK" {
			if len(identTokens) != 1 {
				// no return value here to give error?
				// fmt.Println("more than 1 ident token in struct declaration was not coded for")
				os.Exit(9)
			}

			// // We don't want the whitespace to propogate down so
			// // unshift again and then reshift what was there before
			// p.Unshift()
			// p.Shift()
			// p.Shift()

			// // TODO: we need to do something for declarations
			// // identTokens[0].Type = token.Struct
			// identTokens[0].Type = "STRUCT"
			// identTokens[0].Value.Acting = identTokens[0].Value.String

		} else {
			// If theres a newline then put a separator
			identTokens = append(identTokens, token.Token{
				ID:   0,
				Type: token.Separator,
			})
		}
		// We don't want the whitespace to propogate down so
		// unshift again and then reshift what was there before
		p.Unshift()
		p.Shift()

	} else if p.NextToken.Type == token.LParen {
		// Check if the last token was function
		if p.LastToken.Type == token.Function {
		} else {
			identTokens = append(identTokens[:len(identTokens)-1], p.ParseFunctionCall())
		}
	}

	*blockTokens = append(*blockTokens, identTokens...)
}

// ParseBlock parses the center piece of the language; the block. Anything encapulated in curly braces.
func (p *Parser) ParseBlock() token.Token {

	// FIXME: could do something fancy with another meta and then use that but w/e
	blockTokens := []token.Token{}

	for {
		p.Shift()

		current := p.CurrentToken
		// fmt.Println("token", current)

		switch current.Type {
		// TODO: this needs to change to PRI_OP
		case token.PriOp:
			//fmt.Println("found a pri_op")
			blockTokens = append(blockTokens, current)

		case token.SecOp:
			//fmt.Println("found a sec_op")
			if p.NextToken.Type == current.Type {
				p.Shift()
				if t, ok := token.TokenMap[current.Value.String+p.CurrentToken.Value.String]; ok {
					blockTokens = append(blockTokens, t)
				} else {
					//fmt.Println("wtf happened here: ", current.Value.String+p.CurrentToken.Value.String)
					// os.Exit(9)
					return token.Token{}
				}
			} else {
				blockTokens = append(blockTokens, current)
			}
			// blockTokens = append(blockTokens, current)

		case token.Array:
			//fmt.Println("found an array")
			blockTokens = append(blockTokens, current)

		case token.Keyword:
			if p.CurrentToken.Value.String == "func" {
				// if p.NextToken.Type != token.Ident {
				// 	//fmt.Println("wtf im here??")
				// 	os.Exit(9) // for now
				// }

				// //fmt.Println("IMPLEMENT p.ParseFunctionDef here")
				function := p.ParseFunctionDef()
				//fmt.Println("function", function)
				// os.Exit(9)
				blockTokens = append(blockTokens, function)

			} else {
				blockTokens = append(blockTokens, p.CurrentToken)
			}
			// switch current.Value.Type {
			// case token.SQL:
			// 	//fmt.Println("found a sql keyword")
			// }
			// os.Exit(9)

		case token.GThan:
			//fmt.Println("found a greater than")
			blockTokens = append(blockTokens, current)

		case token.LThan:
			//fmt.Println("found a greater than")
			blockTokens = append(blockTokens, current)

		case token.Increment:
			//fmt.Println("found an increment")
			blockTokens = append(blockTokens, current)

		case token.At:
			//fmt.Println("found an at")
			blockTokens = append(blockTokens, current)

		// TODO: put all of these at the bottom
		// Don't do anything with these for now except append them
		// FIXME: hack to fix the repitition
		case token.Block:
			// blockTokens = append(blockTokens, //p.ParseBlock())
			blockTokens = append(blockTokens, current)
		case token.Init:
			fallthrough
		case token.Attribute:
			fallthrough
		case token.Function:
			//fmt.Println("woah its a function")

			// loop over len(trueValue)
			// 	Create a new parser
			//		parse each group
			// 		replace the group with the output

			// currentTrue := current.Value.True.([]token.Token)

			// currentCopy := current
			// currentCopy.Value.True = New(currentTrue).ParseGroup()
			// //fmt.Println("CURRENT COPY", currentCopy)
			blockTokens = append(blockTokens, current)

		case token.Group:
			//fmt.Println("\nGOTAGROUP")

			// Getting an ident for the lastToken
			//fmt.Println(p.LastToken)
			//fmt.Println(p.CurrentToken)
			//fmt.Println(p.NextToken)

			// FIXME: just assume for now that groups are only used in functions
			// so if we get one we need to determine what kind of function we
			// are in

			functionTokens := []token.Token{current}

			peek := p.NextToken
			//fmt.Println("stuff and yeah", peek)
			// TODO: FIXME: for now we are going to assume that two groups only appear in sequence for a function
			switch peek.Type {
			case token.Group:
				// blockTokens = append(blockTokens, //p.ParseFunctionDef(current))
				p.Shift()
				//fmt.Println(p.LastToken)
				//fmt.Println(p.CurrentToken)

				// Getting block for the nextToken
				//fmt.Println(p.NextToken)
				//fmt.Println("stuff and yeah", p.NextToken)
				functionTokens = append(functionTokens, p.CurrentToken)

				if p.NextToken.Type == token.Block {
					p.Shift()
					functionToken := token.Token{
						ID:   4,
						Type: token.Function,
						Value: token.Value{
							Type: "def",
							True: append(functionTokens, p.CurrentToken),
						},
					}
					//fmt.Println("functionToken", functionToken)
					blockTokens = append(blockTokens, functionToken)
				}

			case token.Block:
				p.Shift()

				// TODO: could make a change here to instead just put it as a group but w/e
				// if p.LastCollectedToken.Type == token.Keyword {

				// }

				blockTokens = append(blockTokens, token.Token{
					ID:   4,
					Type: token.Function,
					Value: token.Value{
						Type: "def",
						True: append(functionTokens, p.CurrentToken),
					},
				})

			default:
				//fmt.Printf("%+v\n%+v", p.CurrentToken, p.NextToken)
				blockTokens = append(blockTokens, p.CurrentToken)
				// //fmt.Printf("wtf peek following group %+v \n%+v\n", peek, p)
				// os.Exit(8)

			}

		case token.Hash:
			//blockTokens = append(blockTokens, //p.ParseAttribute())

		case token.Separator:
			fallthrough

		case token.EOS:
			// TODO: this will need to check the last and next token type later to determine wtf to do
			blockTokens = append(blockTokens, p.CurrentToken)

		case token.Whitespace:
			continue

		case token.Type:
			blockTokens = append(blockTokens, p.CurrentToken)
			peek := p.NextToken
			switch peek.Type {
			case token.Array:
				blockTokens = append(blockTokens, peek)

			case token.Ident:
				p.Shift()
				p.ParseIdent(&blockTokens, p.CurrentToken)

			case token.Literal:
				blockTokens = append(blockTokens, p.CurrentToken)
				p.Shift()
				p.CurrentToken.Type = token.Ident
				blockTokens = append(blockTokens, p.CurrentToken)

			case token.LBracket:
				//fmt.Println("found array", current)
				p.Shift()
				if p.NextToken.Type != token.RBracket {
					//fmt.Println("syntax ERROR: missing ] after type declaration")
					os.Exit(8)
				}

				// later on this should be changed to anything + [] becomes array
				// FIXME: fix this and make the ok check
				arrayToken, ok := token.TokenMap[current.Value.String+peek.Value.String+p.NextToken.Value.String]
				if !ok {
					//fmt.Println("TokenMap check failed on", current.Value.String+peek.Value.String+p.NextToken.Value.String)
					os.Exit(9)
				}

				p.Shift()
				// p.Shift()
				// //fmt.Println("arrayToken", arrayToken)
				// parsedArrayToken := p.ParseArray()
				// var arrayTokens []token.Token
				// // var ok bool
				// arrayTokens, ok = parsedArrayToken.Value.True.([]token.Token)
				// arrayToken.Value.True = arrayTokens
				// //fmt.Println("arrayToken", arrayToken, ok)
				// // blockTokens = append(blockTokens, arrayToken)
				blockTokens[len(blockTokens)-1] = arrayToken

				//fmt.Println()
				//fmt.Println("blockTokens", blockTokens)
				//fmt.Println()
				// p.Shift()
				// blockTokens = append(blockTokens, //p.ParseArray())
				// p.Shift()
				// //fmt.Println("p.Current shit", p.CurrentToken)

				// if p.CurrentToken.Type != token.Ident {
				// 	//fmt.Println("syntax error: no ident after array type declaration")
				// 	os.Exit(8)
				// }
				// //p.ParseIdent(&blockTokens, p.CurrentToken)

			default:
				//fmt.Printf("meta %+v\n", p)
				//fmt.Println("ERROR after type declaration: peek, current", peek, current)
				os.Exit(77)
			}

		case token.Assign:
			//fmt.Println("ASSIGN", current)
			//fmt.Printf("CURRENTVALUETYPE %+v\n", current)
			switch current.Value.Type {
			case "set":
				peek := p.NextToken
				//fmt.Println("PEEK", peek)
				switch peek.Type {
				case token.Assign:
					//fmt.Println("FOUND :=", current.Value.String+peek.Value.String)
					if t, ok := token.TokenMap[current.Value.String+peek.Value.String]; ok {
						blockTokens = append(blockTokens, t)
						p.Shift()
					}
				default:
					blockTokens = append(blockTokens, p.CurrentToken)
				}

			case "assign":
				fallthrough
			case "init":
				blockTokens = append(blockTokens, current)

			default:
				// blockTokens = append(blockTokens, current)
				// continue
				//fmt.Println("ERROR, how did we get in here without an assign type token", current)
				os.Exit(9)
			}

		case token.Ident:
			peek := p.NextToken

			if peek.Type == token.LParen {
				//fmt.Println("IMPLEMENT p.ParseFunctionCall")
				blockTokens = append(blockTokens, p.ParseFunctionCall())
			} else {
				// fmt.Println("parseIDENT")
				p.ParseIdent(&blockTokens, p.CurrentToken)
			}

			// TODO: this case might need to move to the Syntactic part of the parser
		case token.Literal:
			// TODO: this may cause some problems
			// TODO: this is causing some problems
			// switch p.PeekLastCollectedToken().Type {
			// case "SET":
			// 	fallthrough

			// case token.Assign:
			// 	fallthrough

			// case token.Init:
			// 	blockTokens = append(blockTokens, p.CurrentToken)
			// }
			blockTokens = append(blockTokens, p.CurrentToken)

		case token.LParen:
			blockTokens = append(blockTokens, p.ParseGroup())

		case token.RParen:
			// FIXME: why

		case token.LBracket:
			blockTokens = append(blockTokens, p.ParseArray())

		case token.LBrace:
			blockTokens = append(blockTokens, p.ParseBlock())

		case token.RBrace:
			return token.Token{
				ID:   0,
				Type: token.Block,
				// Expected: TODO: do the same thing that we did on the array but use the meta tokens
				Value: token.Value{
					Type: token.Block,
					True: blockTokens,
					// String: TODO: do the same thing that we did on array
				},
			}

		case token.DQuote:
			blockTokens = append(blockTokens, p.ParseString())

		case "":
			//fmt.Println("got nothing")

		default:
			// fmt.Println("IDK WTF TO DO with this token", p.CurrentToken)
			os.Exit(6)
		}
		//fmt.Println(current, p.NextToken)

		if reflect.DeepEqual(p.NextToken, token.Token{}) {
			//fmt.Println()
			//fmt.Println("nextToken block", blockTokens)
			//fmt.Println()
			// //fmt.Println("blockTokens", blockTokens)
			return token.Token{
				ID:   0,
				Type: token.Block,
				// Expected: TODO: do the same thing that we did on the array but use the meta tokens
				Value: token.Value{
					Type: token.Block,
					True: blockTokens,
					// String: TODO: do the same thing that we did on array
				},
			}
		}
	}
}

// Syntactic begins the parsing process for a passes set of tokens
func (p *Parser) Syntactic() ([]token.Token, error) {
	// fmt.Println("p.source", p.source)

	block := p.ParseBlock()
	//fmt.Println("parseBlock", block)

	blockTokens := block.Value.True.([]token.Token)
	//fmt.Println(blockTokens)

	//fmt.Println("going again")
	// Make a new parser and syntactically parse the file
	syntacticTokens := New(blockTokens).ParseBlock()

	//fmt.Println("got stuff", syntacticTokens)

	return []token.Token{
		syntacticTokens,
	}, nil
}
