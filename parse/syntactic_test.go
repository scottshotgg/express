package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/express-rearch/token"
)

var (
	tokens []token.Token
)

func TestSyntactic(t *testing.T) {
	fmt.Println("TestSyntactic")
	TestNewFromFile(t)

	var err error
	tokens, err = p.Syntactic()
	if err != nil {
		fmt.Println("syntacticErr", err)
		t.Fail()
		return
	}

	token.PrintTokens(tokens, "\t")
	fmt.Println()
}
