package parse

import (
	"github.com/pkg/errors"
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

	States          []Parser
	FunctionStrings string
	BlockDepth      int

	DefinedTypes map[string]token.Value
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
		DefinedTypes:     map[string]token.Value{},
	}
	p.InitLogger()
	p.Shift()
	return p
}

// Make this return an error
func (p *Parser) New(tokens []token.Token) *Parser {
	// err := InitLogger()
	// if err != nil {
	// 	return nil, err
	// }

	pp := &Parser{
		// source:           append(append([]token.Token{token.TokenMap["{"]}, tokens...), token.TokenMap["}"]),
		source:           tokens,
		length:           len(tokens),
		meta:             NewMeta(),
		IgnoreWhiteSpace: true,
		ProcessedTokens:  NewStack(),
		DefinedTypes:     p.DefinedTypes,
	}
	pp.InitLogger()
	pp.Shift()
	return pp
}

func (p *Parser) Parse() (token.Value, error) {
	syntacticTokens, err := p.Syntactic()
	if err != nil {
		return token.Value{}, errors.Wrap(err, "p.Syntactic()")
	}

	pNew := New(syntacticTokens)
	semanticToken, err := pNew.Semantic()
	if err != nil {
		return token.Value{}, errors.Wrap(err, "pNew.Semantic()")
	}

	return semanticToken, nil
}
