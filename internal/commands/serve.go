package commands

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/debuqer/kuti/internal/cms"
	"github.com/debuqer/kuti/internal/config"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
)

func Serve() cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serves the application",
		Action: func(c *cli.Context) error {
			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			blog := cms.Blog{}
			blog.Fetch(config.Cfg.Source.Dir)

			template := template.Must(template.New("tpl").Funcs(template.FuncMap{
				"asset":       ServeUrl,
				"url":         ServeQualifiedUrl,
				"post":        blog.GetPost,
				"contentof":   blog.GetPostContent,
				"currentPost": blog.CurrentPost,
				"postsof":     blog.GetPostList,
				"posts":       blog.GetRootPosts,
			}).ParseFiles(path.Join(config.Cfg.Template.Dir, "base.html")))

			toBasePattern := make(map[string]string, 0)

			for pattern, page := range config.Cfg.Routes {
				fullPath := pattern
				if config.Cfg.Server.Ext != "" {
					if fullPath == "/" {
						fullPath += config.Cfg.Server.Ext
					} else {
						fullPath += "/" + config.Cfg.Server.Ext
					}
				}
				toBasePattern[fullPath] = pattern

				if page.Type == "post" {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						page := config.Cfg.Routes[toBasePattern[fullPath]]

						template.ParseFiles(path.Join(config.Cfg.Template.Dir, page.Template))
						blog.RenderPost(w, template, page, p.ByName("filename"))
					})
				} else {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						page := config.Cfg.Routes[toBasePattern[fullPath]]

						blog.RenderIndex(w, template, page)
					})
				}
			}

			fmt.Println("Listening on " + config.Cfg.Server.Host + ":" + config.Cfg.Server.Port + ": ")
			err := http.ListenAndServe(config.Cfg.Server.Host+":"+config.Cfg.Server.Port, router)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}

			return nil
		},
	}
}

func ServeUrl(url string) string {
	port := ""
	if config.Cfg.Server.Port != "" {
		port = (":" + config.Cfg.Server.Port)
	} else {
		port = ""
	}
	return "http://" + config.Cfg.Server.Host + port + "/" + url
}

func ServeQualifiedUrl(url string) string {
	if config.Cfg.Server.Ext != "" {
		return ServeUrl(url) + "/" + config.Cfg.Server.Ext
	}
	return ServeUrl(url)
}
