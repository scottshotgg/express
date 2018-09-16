package parse_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/scottshotgg/express/parse"
	"go.uber.org/zap"
)

func init() {
	os.Setenv(parse.ExpressDebug, "true")
	InitLogger()
}

func InitLoggerAST() error {
	// // FIXME: for now just check 'true' for now
	// if os.Getenv("EXPR_DEBUG") == "true" {
	// 	zapConfig.Development = true
	// }

	// var err error
	// logger, err = zapConfig.Build()
	// if err != nil {
	// 	return err
	// }

	// // Use a sugared logger; slower but has print/f/ln which makes it more versatile and readable
	// // sugar = logger.Sugar ()

	logger, _ = zap.NewProduction()

	return nil
}

var singleFileAST string = "declare_int_ref.expr"

func TestRunSingleAST(t *testing.T) {
	var err error
	parse.LibBase, err = filepath.Abs("../lib/")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = compileExpressProgramAST(singleFileAST)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.FailNow()
	}

	// // Run the code
	// output, err := exec.Command(testBin + singleFileAST + ".exe").CombinedOutput()
	// fmt.Println("Output:\n" + string(output))
	// if err != nil {
	// 	t.Error(err)
	// 	t.FailNow()
	// }
}

func compileExpressProgramAST(filename string) error {
	fmt.Println("file:", filename)
	pathOfFile, err := filepath.Abs(testPrograms + filename)
	if err != nil {
		return err
	}

	lexTokens, err := lexFile(pathOfFile, filename)
	if err != nil {
		return err
	}

	fmt.Println("lexTokens:", lexTokens)

	// syntacticTokens, err := syntacticParseFile(filename, lexTokens)
	// if err != nil {
	// 	return err
	// }

	// semanticTokens, err := semanticParseFile(filename, syntacticTokens)
	// if err != nil {
	// 	return err
	// }

	// err = cppTranspile(filename, semanticTokens)
	// if err != nil {
	// 	return err
	// }

	// var output []byte
	// output, err = exec.Command("clang-format", "-i", testCpp+filename+".cpp").CombinedOutput()
	// fmt.Println("clang-format output:", string(output))
	// if err != nil {
	// 	return err
	// }

	// output, err = exec.Command("clang++", "-std=gnu++2a", testCpp+filename+".cpp", "-o", testBin+filename+".exe").CombinedOutput()
	// fmt.Println("clang++ output:", string(output))
	// if err != nil {
	// 	return err
	// }

	return nil
}
