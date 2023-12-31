package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
)

func ServeCommand() cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serves the application",
		Action: func(c *cli.Context) error {
			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			template := template.Must(template.New("tpl").Funcs(template.FuncMap{
				"asset":       ServeUrl,
				"url":         ServeQualifiedUrl,
				"post":        GetPost,
				"contentof":   GetPostContent,
				"currentPost": CurrentPost,
				"postsof":     GetPostList,
				"posts":       GetRootPosts,
			}).ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			toBasePattern := make(map[string]string, 0)

			for pattern, page := range _conf.Routes {
				fullPath := pattern
				if _conf.Server.Ext != "" {
					if fullPath == "/" {
						fullPath += _conf.Server.Ext
					} else {
						fullPath += "/" + _conf.Server.Ext
					}
				}
				toBasePattern[fullPath] = pattern

				if page.Type == "post" {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						page := _conf.Routes[toBasePattern[fullPath]]

						template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))
						blog.renderPost(w, template, page, p.ByName("filename"))
					})
				} else {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						page := _conf.Routes[toBasePattern[fullPath]]

						blog.renderIndex(w, template, page)
					})
				}
			}

			fmt.Println("Listening on " + _conf.Server.Host + ":" + _conf.Server.Port + ": ")
			err := http.ListenAndServe(_conf.Server.Host+":"+_conf.Server.Port, router)
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
	if _conf.Server.Port != "" {
		port = (":" + _conf.Server.Port)
	} else {
		port = ""
	}
	return "http://" + _conf.Server.Host + port + "/" + url
}

func ServeQualifiedUrl(url string) string {
	if _conf.Server.Ext != "" {
		return ServeUrl(url) + "/" + _conf.Server.Ext
	}
	return ServeUrl(url)
}
