package yon

import "github.com/Scalingo/codegangsta-cli"

func New() *cli.App {

	app := cli.NewApp()
	app.Name = "yon the concatenative interpreting language"
	app.Usage = "write code, and run"

	return app
}
