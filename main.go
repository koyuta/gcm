package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gcm"
	app.Usage = ""
	app.Version = ""
	app.Commands = commands

	return app
}

func msg(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return 1
	}
	return 0
}

func main() {
	app := newApp()
	err := app.Run(os.Args)

	os.Exit(msg(err))
}
