package main

import (
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	cp "github.com/otiai10/copy"
	"github.com/urfave/cli"
)

func BuildCommand() cli.Command {
	return cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the application",
		Action: func(c *cli.Context) error {
			blog := Blog{}
			blog.fetch(_conf.Source.Dir)

			os.RemoveAll("builds")
			os.Mkdir("builds", os.ModePerm)
			cp.Copy("assets", "builds/assets")

			template := template.Must(template.New("tpl").Funcs(template.FuncMap{
				"asset": BuildUrl,
				"url":   BuildQualifiedUrl,
			}).ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			for pattern, page := range _conf.Routes {
				if page.Type == "post" {
					entries, _ := os.ReadDir(path.Join(_conf.Source.Dir, page.Dir))
					for _, e := range entries {
						if !e.IsDir() {
							exploredFile, filename := fileInfo(e)
							pattern = strings.Replace(pattern, ":filename", filename, -1)

							buildNestedDirectories(path.Join("builds/", pattern))

							f, err := os.Create(path.Join("builds/", pattern, _conf.Server.Ext))
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							blog.renderPost(f, template, page, exploredFile)
						}
					}
				} else {
					buildNestedDirectories(path.Join("builds/", pattern))

					f, err := os.Create(path.Join("builds/", pattern, page.Template))
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					blog.renderIndex(f, template, page)
				}
			}

			return nil
		},
	}
}

func BuildUrl(url string) string {
	if _conf.Server.Url == "" {
		_conf.Server.Url, _ = os.Getwd()
		_conf.Server.Url += "/builds/"
	}

	return _conf.Server.Url + url
}

func BuildQualifiedUrl(url string) string {
	if _conf.Server.Ext != "" {
		return BuildUrl(url) + "/" + _conf.Server.Ext
	}

	return BuildUrl(url)
}

func buildNestedDirectories(addr string) {
	curDir := ""

	for _, f := range strings.Split(addr, "/") {
		curDir = path.Join(curDir, f)
		os.Mkdir(curDir, os.ModePerm)
	}
}

func fileInfo(e fs.DirEntry) (exploredFile string, parameter string) {
	exploredFile = e.Name()
	parameter = strings.Replace(e.Name(), "."+_conf.Source.Ext, "", 1)

	return
}
