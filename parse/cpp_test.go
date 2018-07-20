package parse_test

import (
	"fmt"
	"testing"
)

// TODO: FIXME: this needs to be reworked to take the syntactic tokens
func TestTranspile(t *testing.T) {
	fmt.Println("TestTranspile")

	TestSemantic(t)

	statements, err := p.Transpile(semanticBlock)
	if err != nil {
		fmt.Println("semanticErr", err)
		t.Fail()
		return
	}

	fmt.Println("statements", statements)
	fmt.Println()
}
