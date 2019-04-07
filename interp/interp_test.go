package interp_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mk2/yon/interp"
)

func TestInterpEval(t *testing.T) {

	interp := interp.New()
	interp.EvalAndWait(bytes.NewBufferString("1 2 `test` 2 dup true {1 2 3} rot"))
	interp.PrintStack()
	fmt.Println("stdout: ", interp.StdoutString())
	interp.EvalAndWait(bytes.NewBufferString(".s"))
}
