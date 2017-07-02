package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli"
)

var root = os.Getenv("GCM_ROOT")

var commands = []cli.Command{
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

func cmdRoot(c *cli.Context) {
	fmt.Println(root)
}

func cmdNew(c *cli.Context) error {
	var name string
	if c.Args().Present() {
		name = c.Args().First()
	} else {
		fmt.Print("Name: ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			return errors.New("canceled")
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}
	}

	dir := filepath.Join(root, name)
	file := path(dir)

	if exists(file) {
		return fmt.Errorf("%s is already exists.", name)
	}
	if err := os.Mkdir(dir, 0777); err != nil {
		return err
	}
	if _, err := os.Create(file); err != nil {
		return err
	}
	return runEditor(file)
}

func cmdEdit(c *cli.Context) error {
	var name string
	if c.Args().Present() {
		name = c.Args().First()
	} else {
		return fmt.Errorf("")
	}

	dir := filepath.Join(root, name)

	if exists(dir) {
		return fmt.Errorf("%s is not exists.", dir)
	}
	return runEditor(path(dir))
}

func cmdList(c *cli.Context) {
	fmt.Println("cmdList")
}

func cmdRm(c *cli.Context) {
	name := c.Args().First()
	if name == "" {
		cli.ShowCommandHelp(c, "rm")
		os.Exit(1)
	}

	dir := fmt.Sprintf("%s/%s", root, name)
	_, err := os.Stat(dir)
	if err == nil {
		if err := os.RemoveAll(dir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func cmdImport(c *cli.Context) {
	fmt.Println("cmdNew")
}

func cmdShow(c *cli.Context) {
	fmt.Println("cmdShow")
}

func cmdBuild(c *cli.Context) {
	fmt.Println("cmdBuild")
}

func cmdRun(c *cli.Context) {
	fmt.Println("cmdRun")
}

func runEditor(file string) error {
	cmd := exec.Command("vim", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = root

	return cmd.Run()
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func path(dirs ...string) string {
	dir := filepath.Join(dirs...)
	return filepath.Join(dir, "Dockerfile")
}
