package commands

import (
	"errors"
	"os"

	"github.com/debuqer/kuti/pkg/cl"
	"github.com/urfave/cli"
)

func Init() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a new KUTI project",
		Action: func(c *cli.Context) error {
			if c.Args().First() == "" {
				cl.Write("Please provide a valid project name")
				return nil
			}

			configTpl, err := os.ReadFile("assets/config.yml.example")
			if err != nil {
				cl.WriteErr(errors.New("could not initialize the project, assets/config.yml.example not found"))
			}

			os.WriteFile("config.yml", configTpl, 0644)
			cl.Write("Config.yml added to project")

			cl.Write("Initialized a new KUTI project: ", c.Args().First())

			return nil
		},
	}
}
