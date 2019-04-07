package client

import (
	"bufio"
	"os"

	"github.com/mk2/yon/repl/kit"
)

type client struct {
	r *bufio.Reader
	s kit.ReplServer
}

func New(s kit.ReplServer) kit.ReplClient {

	return &client{
		r: bufio.NewReader(os.Stdin),
		s: s,
	}
}

func (c *client) Read() (string, error) {

	if s, err := c.r.ReadString('\n'); err != nil {
		return "", err
	} else {
		return s, nil
	}
}

func (c *client) ShowHelp(s string) string {

	return ""
}

func (c *client) Eval(s string) string {

	c.s.Send(s)
	return c.s.Receive(0)
}

func (c *client) EvalFile(f string) (string, error) {

	return "", nil
}
