package main

import (
	"fmt"
	"net/http"
	"path"
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
			blog.fetch(_conf)

			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			template := template.Must(template.New("tpl").ParseFiles(path.Join(_conf.Template.Dir, "base.html")))

			router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				template.ParseFiles(path.Join(_conf.Template.Dir, "index.html"))

				template.ExecuteTemplate(w, "index.html", blog)
			})

			router.GET("/article/:article", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				template.ParseFiles(path.Join(_conf.Template.Dir, "article.html"))

				blog.CurrentPost = blog.find(_conf, p.ByName("article"))

				template.ExecuteTemplate(w, "article.html", blog)
			})

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
