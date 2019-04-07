package server

import (
	"bytes"

	"github.com/mk2/yon/interp"
	interpkit "github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/repl/kit"
)

type server struct {
	interp interpkit.Interpreter
	buf    *bytes.Buffer
}

func New() kit.ReplServer {

	return &server{
		interp: interp.New(),
		buf:    new(bytes.Buffer),
	}
}

func (s *server) Send(input string) error {

	s.buf.Reset()
	s.buf.WriteString(input)

	return s.interp.EvalAndWait(s.buf)
}

func (s *server) Receive(timeoutSeconds int) string {

	return s.interp.StdoutString()
}
