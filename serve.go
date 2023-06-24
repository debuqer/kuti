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

func ServeCommand(_conf Config) cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serves the application",
		Action: func(c *cli.Context) error {
			port := ""
			if _conf.Server.Port != "" {
				port = (":" + _conf.Server.Port)
			} else {
				port = ""
			}
			_conf.Server.Url = "http://" + _conf.Server.Host + port + "/"

			blog := Blog{}
			blog.fetch(_conf, _conf.Source.Dir)

			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			template := template.Must(template.New("tpl").ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			for pattern, page := range _conf.Routes {
				fullPath := pattern
				if page.Parameter != "" {
					fullPath += ":" + page.Parameter
				}

				router.GET(fullPath, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
					segments := strings.Split(fullPath, "/")
					lastP := segments[0 : len(segments)-1]
					exceptParameter := strings.Join(lastP, "/") + "/"

					page := _conf.Routes[exceptParameter]

					template.ParseFiles(path.Join(_conf.Template.Dir, page.Template))
					if page.Parameter != "" && p.ByName(page.Parameter) != "" {
						blog.CurrentPost = blog.find(_conf, path.Join(_conf.Source.Dir, page.Dir, p.ByName(page.Parameter)))
					}

					template.ExecuteTemplate(w, page.Template, blog)
				})
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
