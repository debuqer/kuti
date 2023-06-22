package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var _conf Config

func main() {
	app := cli.NewApp()
	app.Name = "Kuti"
	app.Usage = "Build your fast and secure hacker blog"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	err := _conf.load()
	if err != nil {
		log.Fatal(err)
	}

	app.Commands = []cli.Command{
		InitCommand(),
		ServeCommand(),
		BuildCommand(),
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
