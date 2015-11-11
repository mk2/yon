package server

import (
	"bytes"

	"github.com/mk2/yon/interp"
)
import interpkit "github.com/mk2/yon/interp/kit"
import "github.com/mk2/yon/repl/kit"

type server struct {
	interp     interpkit.Interpreter
	buf        *bytes.Buffer
	clientTxCh kit.TxCh // TODO
	clientRxCh kit.RxCh // TODO
}

func New() kit.ReplServer { // TODO

	return &server{
		interp: interp.New(),
		buf:    new(bytes.Buffer),
	}
}

func (s *server) Start() error {

	return nil
}

func (s *server) Send(input string) error {

	return nil
}

func (s *server) Receive(timeoutSeconds int) (string, error) {

	return "", nil
}
