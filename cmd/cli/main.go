package main

import (
	"os"

	"github.com/debuqer/kuti/internal/commands"
	"github.com/debuqer/kuti/internal/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kuti"
	app.Usage = "Build your fast and secure hacker blog"

	err := config.Cfg.Load()
	if err != nil {
		panic(err)
	}

	app.Commands = []cli.Command{
		commands.Init(),
		commands.Serve(),
		commands.Build(),
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}