package parse_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/parse"
	"github.com/scottshotgg/express/token"
)

const (
	testRoot   = "../test/"
	testOutput = testRoot + "output/"

	testLex       = testOutput + "lex/"
	testSemantic  = testOutput + "sem/"
	testSyntactic = testOutput + "syn/"
	testCpp       = testOutput + "cpp/"
	testBin       = testOutput + "bin/"

	testPrograms = testRoot + "programs/"

// 	pathOfFile, filename string
// 	lexer                *lex.Lexer
// 	errexpress                 error
// 	lexTokens            []token.Token
// 	// semanticTokens []token.Token
)

var (
	semanticBlock token.Value

	semanticBlockMap = map[string]token.Value{}
)

func init() {
	os.Setenv(parse.ExpressDebug, "true")
}

func TestSemantic(t *testing.T) {
	fmt.Println("TestSemantic")

	TestSyntactic(t)

	var err error
	semanticBlock, err = parse.New(tokens).Semantic()
	if err != nil {
		fmt.Println("semanticErr", err)
		t.Fail()
		return
	}

	fmt.Println(semanticBlock)

	// file, err := os.Open("../test/output/semantic/declarations.expr.sem.json")
	// if err != nil {
	// 	fmt.Println("openErr", err)
	// 	t.Fail()
	// 	return
	// }

	// semanticTokensFromFile, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("semanticTokensFromFileErr", err)
	// 	t.Fail()
	// 	return
	// }

	semanticBlockJSON, err := json.MarshalIndent(semanticBlock, "", "\t")
	if err != nil {
		fmt.Println("semanticTokensFromFileErr", err)
		t.Fail()
		os.Exit(9)
		return
	}

	// if string(semanticTokensFromFile) != string(semanticBlockJSON) {
	// 	fmt.Println("semanticBlock not the same as test tokens")
	// 	fmt.Println("string(semanticBlockJSON)", string(semanticBlockJSON))
	// 	fmt.Println("string(semanticTokensFromFile)", string(semanticTokensFromFile))
	// 	t.Fail()
	// 	return
	// }

	fmt.Println("semanticBlock", string(semanticBlockJSON))
	fmt.Println()
}

func compileExpressProgram(filename string) error {
	fmt.Println("file:", filename)
	pathOfFile, err := filepath.Abs(testPrograms + filename)
	if err != nil {
		return errors.Wrap(err, "filepath.Abs(testPrograms + filename)")
	}

	var lexTokens []token.Token
	lexTokens, err = lexFile(pathOfFile, filename)
	if err != nil {
		return errors.Wrap(err, "lexFile(pathOfFile, filename)")
	}

	pathOfFile, err = filepath.Abs(testPrograms + filename)
	if err != nil {
		return errors.Wrap(err, "filepath.Abs(testPrograms + filename)")
	}

	syntacticTokens, err := syntacticParseFile(filename, lexTokens)
	if err != nil {
		return errors.Wrap(err, "syntacticParseFile(filename, lexTokens)")
	}

	semanticTokens, err := semanticParseFile(filename, syntacticTokens)
	if err != nil {
		return errors.Wrap(err, "semanticParseFile(filename, syntacticTokens)")
	}

	err = cppTranspile(filename, semanticTokens)
	if err != nil {
		return errors.Wrap(err, "cppTranspile(filename, semanticTokens)")
	}

	_, err = exec.Command("clang-format", "-i", testCpp+filename+".cpp").CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "exec.Command(\"clang-format\" ...")
	}

	_, err = exec.Command("clang++", "-std=gnu++2a", testCpp+filename+".cpp", "-o", testBin+filename+".exe").CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "exec.Command(\"clang++\" ...")
	}

	return nil
}

var singleFile string = "var.expr"

func TestRunSingle(t *testing.T) {
	TestSingle(t)

	filepath := testBin + singleFile + ".exe"

	// Run the code
	output, err := exec.Command(filepath).CombinedOutput()
	fmt.Println("Output:", string(output))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSingle(t *testing.T) {
	var err error
	parse.LibBase, err = filepath.Abs("../lib/")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	err = compileExpressProgram(singleFile)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

// TODO: this needs to print out a summary of what passed, what stages failed, etc
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

	var err error
	parse.LibBase, err = filepath.Abs("../lib/")
	if err != nil {
		t.Error(err)
	}

	files, err := ioutil.ReadDir(testPrograms)
	if err != nil {
		t.Error(err)
	}
	if len(files) == 0 {
		fmt.Println("No files to test in:", testPrograms)
		return
	}

	err = os.RemoveAll(testBin)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	// FIXME: w/e just use these permissions for now
	err = os.Mkdir(testBin, 0777)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	// wg := &sync.WaitGroup{}
	for _, file := range files {
		filename := file.Name()
		if !file.IsDir() && filename[0] != '.' {
			// FIXME: for some reason the go funcs are fucking it up rn,
			// probably a global or something
			// wg.Add(1)
			// go func(file os.FileInfo) {
			// 	defer wg.Done()
			err = compileExpressProgram(filename)
			if err != nil {
				t.Error(err)
			}

			// }(file)
		}
	}

	// wg.Wait()

	fmt.Println("Finished!")
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

func semanticParseFile(filename string, syntacticTokens []token.Token) (token.Value, error) {
	// Make a new parser and semantically parse the file
	semanticTokens, err := parse.New(syntacticTokens).Semantic()
	if err != nil {
		// TODO:
		return token.Value{}, err
	}

	semanticBlockMap[filename] = semanticTokens

	semanticTokensJSON, err := json.MarshalIndent(semanticTokens, "", "\t")
	if err != nil {
		// TODO:
		return token.Value{}, err
	}

	err = writeTokensJSONToFile(semanticTokensJSON, testSemantic+filename+".sem.json")
	if err != nil {
		// TODO:
		return token.Value{}, err
	}

	return semanticTokens, nil
}

func cppTranspile(filename string, semanticBlock token.Value) error {
	// Make a new parser and semantically parse the file
	cpp, err := parse.New([]token.Token{}).Transpile(semanticBlock)
	if err != nil {
		// TODO:
		return err
	}

	fmt.Println("cpp transpile: \n------------------\n" +
		cpp +
		"\n------------------\n")

	f, err := os.Create(testCpp + filename + ".cpp")
	if err != nil {
		fmt.Println("got an err creating file", err)
		return err
	}
	n, err := f.WriteString(cpp)
	if err != nil {
		// TODO:
		return err
	}
	if n != len(cpp) {
		// TODO:
	}
	err = f.Close()
	if err != nil {
		// TODO:
	}

	return nil
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
