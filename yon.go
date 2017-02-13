package yon

import (
	"github.com/mk2/yon/repl"
	"github.com/urfave/cli"
)

// New for instantiate yon instance
func New() *cli.App {

	app := cli.NewApp()
	app.Name = "yon the concatenative interpreting language"
	app.Usage = "write code, and run"
	app.Commands = []cli.Command{
		repl.NewCommand(),
	}

	return app
}
