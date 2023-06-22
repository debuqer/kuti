package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
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
			if _conf.Server.Url == "" {
				_conf.Server.Url, _ = os.Getwd()
				_conf.Server.Url += "/builds/"
			}

			blog := Blog{}
			blog.fetch(_conf, _conf.Source.Dir)

			cp.Copy("assets", "builds/assets")

			template := template.Must(template.New("tpl").ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			for pattern, page := range _conf.Routes {
				if page.Parameter != "" {
					entries, _ := os.ReadDir(path.Join(_conf.Source.Dir, page.Dir))
					for _, e := range entries {
						if !e.IsDir() {
							param := e.Name()
							outputDir := strings.Replace(e.Name(), "."+_conf.Source.Ext, "", 1)
							outputFile := "index.html"

							template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))

							curDir := ""
							fmt.Println(strings.Split(path.Join("builds/", pattern, outputDir), "/"))
							for _, f := range strings.Split(path.Join("builds/", pattern, outputDir), "/") {
								curDir = path.Join(curDir, f)
								os.Mkdir(curDir, os.ModePerm)
							}
							f, err := os.Create(path.Join("builds/", pattern, outputDir, outputFile))
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							blog.CurrentPost = blog.find(_conf, path.Join(_conf.Source.Dir, page.Dir, param))
							template.ExecuteTemplate(f, page.Template, blog)
						}
					}
				} else {
					template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))

					f, err := os.Create(path.Join("builds/", pattern, page.Template))
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
