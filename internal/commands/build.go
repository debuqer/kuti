package commands

import (
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"text/template"

	"github.com/debuqer/kuti/internal/cms"
	"github.com/debuqer/kuti/internal/config"
	cp "github.com/otiai10/copy"
	"github.com/pahanini/go-sitemap-generator"
	"github.com/urfave/cli"
)

func Build() cli.Command {
	return cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "build [project-root] [build-address]",
		Action: func(c *cli.Context) error {
			var wg sync.WaitGroup

			os.RemoveAll("builds")
			os.Mkdir("builds", os.ModePerm)
			cp.Copy("assets", "builds/assets")

			g := sitemap.New(sitemap.Options{
				Filename: "sitemap",
				Dir:      "builds/",
				BaseURL:  ServeUrl(""),
			})
			g.Open()

			blog := cms.Blog{}
			blog.Fetch(config.Cfg.Source.Dir)

			template := template.Must(template.New("tpl").Funcs(template.FuncMap{
				"asset":       BuildUrl,
				"url":         BuildQualifiedUrl,
				"post":        blog.GetPost,
				"contentof":   blog.GetPostContent,
				"currentPost": blog.CurrentPost,
				"postsof":     blog.GetPostList,
				"posts":       blog.GetRootPosts,
			}).ParseFiles(path.Join(config.Cfg.Template.Dir, "base.html")))

			for pattern, page := range config.Cfg.Routes {
				if page.Type == "post" {
					entries, _ := os.ReadDir(path.Join(config.Cfg.Source.Dir, page.Dir))
					for _, e := range entries {
						if !e.IsDir() {
							exploredFile, filename := fileInfo(e)
							dest := strings.Replace(pattern, ":filename", filename, -1)

							buildNestedDirectories(path.Join("builds/", dest))

							f, err := os.Create(path.Join("builds/", dest, config.Cfg.Server.Ext))
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							blog.RenderPost(f, template, page, exploredFile)
							go g.Add(sitemap.URL{Loc: ServeQualifiedUrl(dest), Priority: `0.5`})
						}
					}
				} else {
					dest := path.Join(pattern)
					buildNestedDirectories(path.Join("builds/", dest))

					f, err := os.Create(path.Join("builds/", dest, config.Cfg.Server.Ext))
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					blog.RenderIndex(f, template, page)
					go g.Add(sitemap.URL{Loc: ServeQualifiedUrl(dest), Priority: `0.5`})
				}
			}

			wg.Wait()

			e := g.Close()
			if e != nil {
				panic(e)

			}

			return nil
		},
	}
}

func BuildUrl(url string) string {
	if config.Cfg.Server.Url == "" {
		config.Cfg.Server.Url, _ = os.Getwd()
		config.Cfg.Server.Url += "/builds/"
	}

	return config.Cfg.Server.Url + url
}

func BuildQualifiedUrl(url string) string {
	if config.Cfg.Server.Ext != "" {
		return BuildUrl(url) + "/" + config.Cfg.Server.Ext
	}

	return BuildUrl(url)
}

func buildNestedDirectories(addr string) {
	curDir := ""

	for _, f := range strings.Split(addr, "/") {
		curDir = path.Join(curDir, f)
		if _, err := os.Stat(curDir); os.IsNotExist(err) {
			os.Mkdir(curDir, os.ModePerm)
		}
	}
}

func fileInfo(e fs.DirEntry) (exploredFile string, parameter string) {
	exploredFile = e.Name()
	parameter = strings.Replace(e.Name(), "."+config.Cfg.Source.Ext, "", 1)

	return
}
