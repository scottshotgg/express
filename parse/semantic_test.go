package parse_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/scottshotgg/ExpressRedo/lex"
	"github.com/scottshotgg/ExpressRedo/parse"
	"github.com/scottshotgg/ExpressRedo/token"
)

const (
	testRoot   = "../test/"
	testOutput = testRoot + "output/"

	testLex       = testOutput + "lex/"
	testSemantic  = testOutput + "semantic/"
	testSyntactic = testOutput + "syntactic/"
	testCpp       = testOutput + "cpp/"
	testBin       = testOutput + "bin/"

	testPrograms = testRoot + "programs/"

// 	pathOfFile, filename string
// 	lexer                *lex.Lexer
// 	err                  error
// 	lexTokens            []token.Token
// 	// semanticTokens []token.Token
)

func TestSemantic(t *testing.T) {
	fmt.Println("TestSemantic")

	TestSyntactic(t)

	values, err := parse.New(tokens).Semantic()
	if err != nil {
		fmt.Println("semanticErr", err)
		t.Fail()
		return
	}

	token.PrintValues(values, "\t")
	fmt.Println()
}

func TestAll(t *testing.T) {
	// ls ../test/programs directory
	// for each file
	//	- lex the contents
	//	- syntactically parse the contents
	//	- semantically parse the contents
	//	transpile to cpp
	//	generate binary
	//	run binary
	//	capture output
	// At each step, compare result with output dir

	files, err := ioutil.ReadDir(testPrograms)
	if err != nil {
		fmt.Println("ReadDirErr", err)
		t.Fail()
	}

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			pathOfFile, err := filepath.Abs(testPrograms + filename)
			if err != nil {
				fmt.Println("AbsErr", err)
				// TODO: make this more individual later
				t.Fail()
				return
			}
			fmt.Println(pathOfFile)

			var lexTokens []token.Token
			lexTokens, err = lexFile(pathOfFile, filename)
			if err != nil {
				fmt.Println("lexFileErr", err)
				t.Fail()
				continue
			}

			pathOfFile, err = filepath.Abs(testPrograms + filename)
			if err != nil {
				fmt.Println("AbsErr", err)
				// TODO: make this more individual later
				t.Fail()
				return
			}
			fmt.Println(pathOfFile)

			syntacticTokens, err := syntacticParseFile(filename, lexTokens)
			if err != nil {
				fmt.Println("syntacticParseFileErr", err)
				t.Fail()
				continue
			}

			_, err = semanticParseFile(filename, syntacticTokens)
			if err != nil {
				fmt.Println("lexFileErr", err)
				t.Fail()
				continue
			}
		}
	}
}

func lexFile(pathOfFile, filename string) ([]token.Token, error) {
	// Make a new lexer
	lexer, err := lex.NewFromFile(pathOfFile)
	if err != nil {
		fmt.Println("NewFromFileErr", err)
		// t.Fail()
		return []token.Token{}, err
	}

	// Lex the file
	lexTokens, err := lexer.Lex()
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	lexTokensJSON, err := json.MarshalIndent(lexTokens, "", "\t")
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	err = writeTokensJSONToFile(lexTokensJSON, testLex+filename+".lex.json")
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	return lexTokens, nil
}

func syntacticParseFile(filename string, lexTokens []token.Token) ([]token.Token, error) {
	// Make a new parser and syntactically parse the file
	syntacticTokens, err := parse.New(lexTokens).Syntactic()
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	syntacticTokensJSON, err := json.MarshalIndent(syntacticTokens, "", "\t")
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	err = writeTokensJSONToFile(syntacticTokensJSON, testSyntactic+filename+".syn.json")
	if err != nil {
		// TODO:
		return []token.Token{}, err
	}

	return syntacticTokens, nil
}

func semanticParseFile(filename string, syntacticTokens []token.Token) ([]token.Value, error) {
	// Make a new parser and semantically parse the file
	semanticTokens, err := parse.New(syntacticTokens).Semantic()
	if err != nil {
		// TODO:
		return []token.Value{}, err
	}

	semanticTokensJSON, err := json.MarshalIndent(semanticTokens, "", "\t")
	if err != nil {
		// TODO:
		return []token.Value{}, err
	}

	err = writeTokensJSONToFile(semanticTokensJSON, testSemantic+filename+".sem.json")
	if err != nil {
		// TODO:
		return []token.Value{}, err
	}

	return semanticTokens, nil
}

func writeTokensJSONToFile(tokensJSON []byte, pathOfFile string) error {
	lexFile, err := os.Create(pathOfFile)
	if err != nil {
		// TODO:
		return err
	}

	n, err := lexFile.Write(tokensJSON)
	if err != nil {
		// TODO:
		return err
	}
	if n != len(tokensJSON) {
		// TODO:
		// need to rewrite the lexFile
		return errors.New("Not all bytes were written")
	}

	err = lexFile.Close()
	if err != nil {
		// TODO:
		return err
	}

	return nil
}
