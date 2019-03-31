package repl

import (
	"fmt"
	"log"

	"github.com/mk2/yon/repl/client"
	"github.com/mk2/yon/repl/kit"
	"github.com/mk2/yon/repl/server"
	"github.com/urfave/cli"
)

type repl struct {
	c kit.ReplClient
	s kit.ReplServer
}

func NewCommand() cli.Command {

	return cli.Command{
		Name:    "repl",
		Aliases: []string{"r"},
		Usage:   "start yon repl",
		Action: func(c *cli.Context) {

			log.Println("starting repl...")
			repl := New()
			log.Println("repl started!!")

			for {
				fmt.Printf("(user) ")
				if s, err := repl.GetClient().Read(); err != nil {
					continue
				} else {
					fmt.Printf("=> %s\n", repl.GetClient().Eval(s))
				}
			}
		},
	}
}

func New() kit.Repl {

	s := server.New()
	c := client.New(s)

	return &repl{
		s: s,
		c: c,
	}
}

func (r *repl) GetClient() kit.ReplClient {

	return r.c
}

func (r *repl) GetPrimaryServer() kit.ReplServer {

	return r.s
}

func (r *repl) GetServers() []kit.ReplServer {

	return nil
}
