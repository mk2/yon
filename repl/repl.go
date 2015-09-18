package repl

import "github.com/mk2/yon/interp"

type Repl struct {
	interp *interp.interpreter
}

func New() *Repl {

	return &Repl{
		interp: interp.New(),
	}
}
