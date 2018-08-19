package libc_test

import (
	"testing"

	"github.com/scottshotgg/express/libc"
)

func TestParseLibCFunctions(t *testing.T) {
	libc.ParseLibCFunctions()
}

func TestRipLibC(t *testing.T) {
	libc.RipLibC()
}
