package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kuti"
	app.Usage = "Build your fast and secure hacker blog"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize a new kuti project",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("Project name is invalid")
					return nil
				}

				configTpl, err := os.ReadFile("assets/config.yml")
				if err != nil {
					fmt.Println("Template file assets/config.yml not found")
					return nil
				}
				os.WriteFile("config.yml", configTpl, 0644)

				fmt.Println("Initialized a new kuti project: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
