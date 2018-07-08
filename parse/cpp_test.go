package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/ExpressRedo/parse"
)

// TODO: FIXME: this needs to be reworked to take the syntactic tokens
func TestTranspile(t *testing.T) {
	fmt.Println("TestTranspile")

	TestSemantic(t)

	statements, err := parse.New(tokens).Transpile()
	if err != nil {
		fmt.Println("semanticErr", err)
		t.Fail()
		return
	}

	fmt.Println(statements)
	fmt.Println()
}
