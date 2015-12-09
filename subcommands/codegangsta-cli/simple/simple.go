// https://github.com/codegangsta/cli
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func boom(c *cli.Context) {
	println("boom! I say!")
}

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = boom

	app.Run(os.Args)
}
