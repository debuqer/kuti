package main

import (
	"os"

	"github.com/urfave/cli"
)

var _conf Config
var blog Blog

func main() {
	app := cli.NewApp()
	app.Name = "Kuti"
	app.Usage = "Build your fast and secure hacker blog"

	err := _conf.load()
	if err != nil {
		panic(err)
	}
	blog = Blog{}
	blog.fetch(_conf.Source.Dir)
	app.Commands = []cli.Command{
		InitCommand(),
		ServeCommand(),
		BuildCommand(),
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
