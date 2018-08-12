package parse

import (
	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/token"
)

type Token struct {
	TokenClass string
	Type       string
	Lexemes    []lex.Lexeme
	Value      map[string]interface{}
}

type Parser struct {
	source []token.Token
	length int
	Index  int
	// CurrentLexeme lex.Lexeme
	// CurrentToken     Token
	meta             *Meta
	Output           []Token
	IgnoreWhiteSpace bool

	// LastLexeme    lex.Lexeme
	// CurrentLexeme lex.Lexeme
	// NextLexeme    lex.Lexeme

	// FIXME: This should be a stack
	// ProcessedTokens []token.Token
	ProcessedTokens *Stack

	LastToken    token.Token
	CurrentToken token.Token
	NextToken    token.Token

	States []Parser
}

func (p *Parser) Length() int { return p.length }

// Make this return an error
func New(tokens []token.Token) *Parser {
	// err := InitLogger()
	// if err != nil {
	// 	return nil, err
	// }

	p := &Parser{
		// source:           append(append([]token.Token{token.TokenMap["{"]}, tokens...), token.TokenMap["}"]),
		source:           tokens,
		length:           len(tokens),
		meta:             NewMeta(),
		IgnoreWhiteSpace: true,
		ProcessedTokens:  NewStack(),
	}
	p.InitLogger()
	p.Shift()
	return p
}

func (p *Parser) Parse() (token.Value, error) {
	syntacticTokens, err := p.Syntactic()
	if err != nil {
		return token.Value{}, err
	}

	pNew := New(syntacticTokens)
	semanticToken, err := pNew.Semantic()
	if err != nil {
		return token.Value{}, err
	}

	return semanticToken, nil
}
