package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v3"
)

func main() {
	app := cli.NewApp()
	app.Name = "Kuti"
	app.Usage = "Build your fast and secure hacker blog"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println("Could not found config.yml")
		return
	}
	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		fmt.Println("Bad configuration config.yml")
		return
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
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serves the application",
			Action: func(c *cli.Context) error {
				http.HandleFunc("/", getRoot)

				fmt.Println("Listening on http://localhost:3333: ")
				err := http.ListenAndServe(":3333", nil)
				if err != nil {
					fmt.Println("Template file assets/config.yml not found")
					return nil
				}

				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}
