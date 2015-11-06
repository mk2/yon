package repl

import "github.com/mk2/yon/repl/kit"

type repl struct {
}

func New() kit.Repl {

	return &repl{}
}

func (r *repl) GetClient() kit.ReplClient {

	return nil
}

func (r *repl) GetPrimaryServer() kit.ReplServer {

	return nil
}

func (r *repl) GetServers() []kit.ReplServer {

	return nil
}
