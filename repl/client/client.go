package client

import "github.com/mk2/yon/repl/kit"

type client struct {
	s kit.ReplServer
}

func New(s kit.ReplServer) kit.ReplClient {

	return &client{
		s: s,
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
