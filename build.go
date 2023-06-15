package main

import (
	"log"
	"os"
	"path"
	"text/template"

	cp "github.com/otiai10/copy"
	"github.com/urfave/cli"
)

func BuildCommand(_conf Config) cli.Command {
	return cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the application",
		Action: func(c *cli.Context) error {
			blog := Blog{}
			blog.fetch(_conf)

			cp.Copy("assets", "builds/assets")

			template := template.Must(template.New("tpl").ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			for pattern, page := range _conf.Routes {
				fullPath := pattern
				if page.Parameter != "" {
					fullPath += ":" + page.Parameter
				}

				if page.Parameter == "" {
					template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))

					f, err := os.OpenFile(path.Join("builds/", pattern, page.Template), os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					template.ExecuteTemplate(f, page.Template, blog)
				}
			}

			return nil
		},
	}
}
