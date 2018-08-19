package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/express/parse"
	"github.com/scottshotgg/express/token"
)

// TODO: FIXME: this needs to be reworked to take the syntactic tokens
func TestTranspile(t *testing.T) {
	// var semanticBlock token.Value
	fmt.Println("TestTranspile")

	TestSemantic(t)
	// // FIXME: need to read from the file for the block
	// file, err := os.Open("../test/output/semantic/loop.expr.sem.json")
	// if err != nil {
	// 	// TODO:
	// 	t.Fail()
	// 	return
	// }
	// filedata, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	// TODO:
	// 	t.Fail()
	// 	return
	// }

	// fmt.Println("filedata", string(filedata))
	// err = json.Unmarshal(filedata, &semanticBlock)
	// if err != nil {
	// 	// TODO:
	// 	t.Fail()
	// 	return
	// }

	cpp, err := parse.New([]token.Token{}).Transpile(semanticBlock)
	if err != nil {
		fmt.Println("transpileErr", err)
		t.Fail()
		return
	}

	fmt.Println("cpp transpile: \n------------------\n" +
		cpp +
		"\n------------------\n")
}
