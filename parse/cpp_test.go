package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/ExpressRedo/token"
)

// TODO: FIXME: this needs to be reworked to take the syntactic tokens
func TestTranspile(t *testing.T) {
	var semanticBlock token.Value
	fmt.Println("TestTranspile")

	TestSemantic(t)

	cpp, err := p.Transpile(semanticBlock)
	if err != nil {
		fmt.Println("transpileErr", err)
		t.Fail()
		return
	}

	fmt.Println("cpp transpile: \n------------------\n" +
		cpp +
		"\n------------------\n")
}
