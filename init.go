package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func InitCommand() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initialize a new kuti project",
		Action: func(c *cli.Context) error {
			if c.Args().First() == "" {
				fmt.Println("Project name is invalid")
				return nil
			}

			configTpl, err := os.ReadFile("assets/config.yml.example")
			if err != nil {
				fmt.Println("Template file assets/config.yml not found")
				return nil
			}
			os.WriteFile("config.yml", configTpl, 0644)

			fmt.Println("Initialized a new kuti project: ", c.Args().First())
			return nil
		},
	}
}
