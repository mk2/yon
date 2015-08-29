package interp_test

import (
	"testing"

	"bytes"

	"github.com/mk2/yon/interp"
)

func TestInterpEval(t *testing.T) {

	interp := interp.New()
	interp.Eval(bytes.NewBufferString("1 2 `test` 2"))
	interp.Wait()
	interp.PrintStack()
}
