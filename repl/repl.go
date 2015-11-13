package repl

import (
	"fmt"
	"log"

	"github.com/Scalingo/codegangsta-cli"
	"github.com/mk2/yon/repl/client"
	"github.com/mk2/yon/repl/kit"
	"github.com/mk2/yon/repl/server"
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

			var s string

			for {
				fmt.Printf("(scratch) ")
				if _, err := fmt.Scanln(&s); err != nil {
					continue
				}
				fmt.Printf("=> %s\n", repl.GetClient().Eval(s))
				s = ""
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
