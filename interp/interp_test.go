package interp_test

import (
	"fmt"
	"testing"

	"bytes"

	"github.com/mk2/yon/interp"
)

func TestInterpEval(t *testing.T) {

	interp := interp.New()
	interp.EvalAndWait(bytes.NewBufferString("1 2 `test` 2 dup true"))
	interp.PrintStack()
	interp.PrintVocab()
	fmt.Println("stdout: ", interp.StdoutString())
}
