package server

import "github.com/mk2/yon/interp"
import interpkit "github.com/mk2/yon/interp/kit"
import "github.com/mk2/yon/repl/kit"

type server struct {
	interp interpkit.Interpreter
}

func New() kit.ReplServer {

	return &server{interp: interp.New()}
}
