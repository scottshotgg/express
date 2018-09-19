package parse

import (
	"fmt"

	"github.com/scottshotgg/express-ast"
	"github.com/scottshotgg/express/token"
)

func GetStatement() (ast.Statement, error) {
	// An example assignment statement of:
	// i := true
	ident, err := ast.NewIdent(ast.Token{}, ast.NewBoolType(), "i")
	if err != nil {
		return nil, err
	}

	as, err := ast.NewAssignment(ast.Token{}, ident, ast.Init, ast.NewBool(ast.Token{}, true))
	if err != nil {
		return nil, err
	}

	// Statements can be:
	//	- assignment
	//		- type
	//		- ident
	//	- block
	//	- call
	//		- ident
	//	- func / fn
	//	- if/else
	//	- loop
	//	- return

	return as, nil
}

// BuildAST builds an AST from the tokens provided by the lexer
func BuildAST(lexTokens []token.Token) (*ast.Program, error) {
	p := ast.NewProgram()

	// Spoof this name for now
	file := ast.NewFile("main.expr")

	i := 0
	for {
		// We know that the file can only consist of statements
		stmt, err := GetStatement()
		if err != nil {
			return nil, err
		}

		file.AddStatement(stmt)

		if i > len(lexTokens)-1 {
			break
		}

		i++
	}

	p.AddFile(file)

	return p, nil
}

func CompressTokens(lexTokens []token.Token) ([]token.Token, error) {
	compressedTokens := []token.Token{}

	alreadyChecked := false

	for i := 0; i < len(lexTokens)-1; i++ {
		fmt.Println("i", lexTokens[i])

		// This needs to be simplified
		if lexTokens[i].Type == "ASSIGN" || lexTokens[i].Type == "SEC_OP" || lexTokens[i].Type == "PRI_OP" && lexTokens[i+1].Type == "ASSIGN" || lexTokens[i+1].Type == "SEC_OP" || lexTokens[i+1].Type == "PRI_OP" {
			compressedToken, ok := token.TokenMap[lexTokens[i].Value.String+lexTokens[i+1].Value.String]
			fmt.Println("added \"" + lexTokens[i].Value.String + lexTokens[i+1].Value.String + "\"")
			if ok {
				compressedTokens = append(compressedTokens, compressedToken)
				i++

				// If we were able to combine the last two tokens and make a new one, mark it
				if i == len(lexTokens)-1 {
					alreadyChecked = true
				}

				continue
			}
		}

		// Filter out the white space
		if lexTokens[i].Type == "WS" {
			continue
		}

		compressedTokens = append(compressedTokens, lexTokens[i])
	}

	// If it hasn't been already checked and the last token is not a white space, then append it
	if !alreadyChecked && lexTokens[len(lexTokens)-1].Type != "WS" {
		compressedTokens = append(compressedTokens, lexTokens[len(lexTokens)-1])
	}

	return compressedTokens, nil
}
