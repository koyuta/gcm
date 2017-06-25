package main

import (
	"os"

	"github.com/urfave/cli"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gcm"
	app.Usage = ""
	app.Version = ""
	app.Commands = []cli.Command{
		{
			Name:    "root",
			Aliases: []string{"r"},
			Usage:   "set a root directory of managed dockerfiles",
			Action:  cmdRoot,
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new dockerfile",
			Action:  cmdNew,
		},
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "import dockerfile",
			Action:  cmdImport,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list managed images",
			Action:  cmdList,
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "edit a dockerfile",
			Action:  cmdEdit,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "show a dockerfile",
			Action:  cmdShow,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build a image",
			Action:  cmdBuild,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run a new container",
			Action:  cmdRun,
		},
		{
			Name:   "rm",
			Usage:  "remove images",
			Action: cmdRm,
		},
	}

	return app
}

func main() {
	app := newApp()
	app.Run(os.Args)
}
