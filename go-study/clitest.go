package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "clitest"
	app.Usage = "test github.com/codegangsta/cli"
	app.CommandNotFound = DefaultAppComplete
	app.Action = func(c *cli.Context) {
		println("xxxxxxxxxxxxxxxxx")
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "n",
					Usage: "show number",
				},
			},
			Action: func(c *cli.Context) {
				if c.Bool("n") {
					println("aaaaaaaaa")
				}
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}

func DefaultAppComplete(c *cli.Context, str string) {
	println(str)
	println("aaaaaaaa")
}
