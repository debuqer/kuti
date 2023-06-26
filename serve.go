package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
)

func ServeCommand() cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serves the application",
		Action: func(c *cli.Context) error {
			blog := Blog{}
			blog.fetch(_conf.Source.Dir)

			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			template := template.Must(template.New("tpl").Funcs(template.FuncMap{
				"asset": ServeUrl,
				"url":   ServeQualifiedUrl,
			}).ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			for pattern, page := range _conf.Routes {
				fullPath := pattern
				if page.Parameter != "" {
					fullPath += ":" + page.Parameter
				}

				if _conf.Server.Ext != "" {
					if fullPath == "/" {
						fullPath += _conf.Server.Ext
					} else {
						fullPath += "/" + _conf.Server.Ext
					}
				}

				if page.Parameter == "" {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						if _conf.Server.Ext != "" {
							fullPath = strings.Replace(fullPath, "/"+_conf.Server.Ext, "", 1)
							fullPath = strings.Replace(fullPath, _conf.Server.Ext, "", 1)
						}

						segments := strings.Split(fullPath, "/")
						lastP := segments[0 : len(segments)-1]
						exceptParameter := strings.Join(lastP, "/") + "/"

						page := _conf.Routes[exceptParameter]

						blog.renderIndex(w, template, page)
					})
				} else {
					router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
						if _conf.Server.Ext != "" {
							fullPath = strings.Replace(fullPath, "/"+_conf.Server.Ext, "", 1)
							fullPath = strings.Replace(fullPath, _conf.Server.Ext, "", 1)
						}

						segments := strings.Split(fullPath, "/")
						lastP := segments[0 : len(segments)-1]
						exceptParameter := strings.Join(lastP, "/") + "/"

						page := _conf.Routes[exceptParameter]

						template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))
						blog.renderPost(w, template, page, p.ByName(page.Parameter))
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
	_conf.Server.Url = "http://" + _conf.Server.Host + port + "/"

	return _conf.Server.Url + url
}

func ServeQualifiedUrl(url string) string {
	if _conf.Server.Ext != "" {
		return ServeUrl(url) + "/" + _conf.Server.Ext
	}
	return ServeUrl(url)
}
