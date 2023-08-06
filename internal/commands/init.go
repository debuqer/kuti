package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Init() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a new KUTI project",
		Action: func(c *cli.Context) error {
			if c.Args().First() == "" {
				fmt.Println("Please provide a valid project name")
				return nil
			}

			configTpl, err := os.ReadFile("assets/config.yml.example")
			if err != nil {
				panic("Could not initialize the project, assets/config.yml.example not found")
			}

			os.WriteFile("config.yml", configTpl, 0644)
			fmt.Println("Config.yml added to project")

			fmt.Println("Initialized a new KUTI project: ", c.Args().First())

			return nil
		},
	}
}
