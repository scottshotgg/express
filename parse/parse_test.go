package parse_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/parse"
	"github.com/scottshotgg/express/token"
)

var (
	p          *parse.Parser
	testTokens []token.Token
	// testLexemesString = []byte(`[{"type":"type","value":"int"},{"type":"space","value":" "},{"type":"ident","value":"i"},{"type":"space","value":" "},{"type":"assign","value":"="},{"type":"space","value":" "},{"type":"lit","value":"int","true":5},{"type":"end"}]`)
	testTokensString = []byte(`[
		{"Type":"TYPE","Value":{"Type":"int","String":"int"}},
		{"Type":"WS","Value":{"Type":"space","String":" "}},
		{"Type":"IDENT","Value":{"String":"i"}},
		{"Type":"WS","Value":{"Type":"space","String":" "}},
		{"Type":"ASSIGN","Value":{"Type":"assign","String":"="}},
		{"Type":"WS","Value":{"Type":"space","String":" "}},
		{"Type":"LITERAL","Value":{"Type":"int","True":5}},
		{"Type":"WS","Value":{"Type":"space","String":" "}}
	]`)
)

func TestNew(t *testing.T) {
	fmt.Println("TestNew")

	err := json.Unmarshal(testTokensString, &testTokens)
	if err != nil {
		fmt.Println("jsonErr", err)
		t.Fail()
		return
	}

	p = parse.New(testTokens)
	if p.Length() != len(testTokens) {
		t.Fail()
	}
}

// Figure out someway where we can put the test in here
func TestNewFromFile(t *testing.T) {
	fmt.Println("TestNewFromFile")

	f, err := os.Open("../test/programs/arithmetic.expr")
	// f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("openErr", err)
		os.Exit(9)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("openErr", err)
		os.Exit(9)
	}

	lexTokens, err := lex.New(string(data)).Lex()
	if err != nil {
		fmt.Println("lexErr", err)
		os.Exit(9)
	}
	testTokens = lexTokens

	p = parse.New(lexTokens)
	if p.Length() != len(lexTokens) {
		t.Fail()
		return
	}
}

func TestShift(t *testing.T) {
	fmt.Println("TestShift")
	TestNew(t)

	p.Shift()
	if reflect.DeepEqual(p.NextToken, testTokens[1]) {
		t.Fail()
	}
}

func TestUnshift(t *testing.T) {
	fmt.Println("TestShift")
	TestNew(t)

	p.Shift()
	p.Shift()
	if reflect.DeepEqual(p.NextToken, testTokens[2]) {
		t.Fail()
	}

	p.Unshift()
	if reflect.DeepEqual(p.NextToken, testTokens[1]) {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	fmt.Println("TestParse")
	TestNew(t)

	value, err := p.Parse()
	if err != nil {
		// TODO:
	}
	fmt.Println("value", value)
}

func TestSaveState(t *testing.T) {
	fmt.Println("TestSaveState")
	TestNew(t)
	p.Shift()
	p.Shift()

	p.SaveState()
	fmt.Println("p", p.States)
}
