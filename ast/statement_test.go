package ast_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/express/ast"
)

func TestNewStatementNode(t *testing.T) {
	s := ast.NewStatementNode(&ast.Location{
		Start: &ast.Pos{
			Line:   0,
			Column: 0,
		},
		End: &ast.Pos{
			Line:   0,
			Column: 0,
		},
	})

	fmt.Println("s", s)
}
