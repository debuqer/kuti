package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func InitCommand(_conf Config) cli.Command {
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
				fmt.Println("Could not init the project, assets/config.yml.example not found")
				return nil
			}
			os.WriteFile("config.yml", configTpl, 0644)
			fmt.Println("Config.yml added to project")

			fmt.Println("Initialized a new KUTI project: ", c.Args().First())
			return nil
		},
	}
}
