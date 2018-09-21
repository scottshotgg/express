package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express-ast"
	"github.com/scottshotgg/express/token"
)

type ASTBuilder struct {
	Tokens []token.Token
	Index  int
}

func (a *ASTBuilder) GetFactor() (ast.Expression, error) {
	currentToken := a.Tokens[a.Index]

	switch currentToken.Type {
	case token.Ident:

	case token.Literal:
		switch currentToken.Value.Type {
		case "int":
			return ast.NewInt(ast.Token{}, currentToken.Value.True.(int)), nil
		}

	default:
		return nil, errors.Errorf("Could not parse factor from token: %+v", currentToken)
	}

	return nil, nil
}

func (a *ASTBuilder) GetTerm() (ast.Expression, error) {
	fmt.Println("a again", a.Tokens[a.Index])

	factor, err := a.GetFactor()
	if err != nil {
		return nil, err
	}

	return factor, nil
}

func (a *ASTBuilder) GetExpression() (ast.Expression, error) {
	fmt.Println("a", a.Tokens[a.Index])

	term, err := a.GetTerm()
	if err != nil {
		return nil, err
	}

	// FIXME: should probably check for secondary operations right here

	return term, nil
}

func (a *ASTBuilder) GetStatement() (ast.Statement, error) {
	// An example assignment statement of:
	// bool i = true
	ident, err := ast.NewBoolIdent(ast.Token{}, "i")
	if err != nil {
		return nil, err
	}

	// TODO: could make a new boolean assignment here?
	as, err := ast.NewAssignment(ast.Token{}, ident, ast.Equals, ast.NewBool(ast.Token{}, true))
	if err != nil {
		return nil, err
	}

	// NEED to switch and capture these
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

	typeOf := ""
	nameOf := ""
	currentToken := a.Tokens[a.Index]
	switch currentToken.Type {
	case token.Type:
		// Look for an ident as the next thing for now
		// fallthrough to the next block for now
		typeOf = currentToken.Value.String

	case token.Ident:
		// Here we will want to look at what is next and handle it
		// If it is an assignment statment then we are looking for an expression afterwards
		nameOf = currentToken.Value.String
		a.Index++
		if a.Tokens[a.Index].Type == "ASSIGN" {
			a.Index++
			expr, err := a.GetExpression()
			if err != nil {
				return nil, err
			}
			fmt.Println("expr", expr)

			// FIXME: need to implement Type() so that we can get the var type
			ident, err := ast.NewIdent(Token{}, expr.Type(), nameOf)

			// TODO: could make a new boolean assignment here?
			as, err := ast.NewAssignment(ast.Token{}, ident, ast.Init, expr)
			if err != nil {
				return nil, err
			}

			// TODO: add statement here later
			return as, nil

		}
		return nil, errors.Errorf("Expected assignment token, got %+v", a.Tokens[a.Index])

		fmt.Println("nameOf", nameOf)

	case token.Block:
		// Here we will want to recursively call GetStatement()
		// however, a block should be able to be parsed for an expression as well

		// This one will have to be figured out when parsing the ident
	// case token.Call:
	case token.Function:
		// Next things we look for after the Function token is:
		//	[ ident ] [ group ] { group } [ block ]

		// TODO: create this token
	// case token.If:

	// This needs to switch to token.Loop later on
	case token.For:
		// Look at how we did the for loop parsing in `semantic.go`

	case token.Return:
		// For now just look for a single expression afterwards

	default:
		return nil, errors.Errorf("Could not get statement from token: %+v", currentToken)
	}

	fmt.Println("typeOf", typeOf)

	return as, nil
}

// BuildAST builds an AST from the tokens provided by the lexer
func (a *ASTBuilder) BuildAST(lexTokens []token.Token) (*ast.Program, error) {
	p := ast.NewProgram()

	// Spoof this name for now
	file := ast.NewFile("main.expr")

	for {
		// We know that the file can only consist of statements
		stmt, err := a.GetStatement()
		if err != nil {
			return nil, err
		}

		file.AddStatement(stmt)

		a.Index++

		if a.Index > len(lexTokens)-1 {
			break
		}
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
