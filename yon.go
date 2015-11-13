package yon

import (
	"github.com/Scalingo/codegangsta-cli"
	"github.com/mk2/yon/repl"
)

func New() *cli.App {

	app := cli.NewApp()
	app.Name = "yon the concatenative interpreting language"
	app.Usage = "write code, and run"
	app.Commands = []cli.Command{
		repl.NewCommand(),
	}

	return app
}
