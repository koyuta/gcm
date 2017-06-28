package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

var root = os.Getenv("GCM_ROOT")

func cmdRoot(c *cli.Context) {
	fmt.Println(root)
}

func cmdNew(c *cli.Context) {
	name := c.Args().First()
	if name == "" {
		cli.ShowCommandHelp(c, "new")
		os.Exit(1)
	}
	dir := fmt.Sprintf("%s/%s", root, name)
	_, err := os.Stat(dir)
	if err != nil {
		if err := os.Mkdir(dir, 0777); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	path := fmt.Sprintf("%s/Dockerfile", dir)

	cmd := exec.Command("vim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = root

	err = cmd.Run()
	if err == nil {
		cmd.Wait()
		os.Exit(0)
	}
}

func cmdImport(c *cli.Context) {
	fmt.Println("cmdNew")
}

func cmdList(c *cli.Context) {
	fmt.Println("cmdList")
}

func cmdEdit(c *cli.Context) {
	fmt.Println("cmdEdit")
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
