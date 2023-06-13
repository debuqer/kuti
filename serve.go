package main

import (
	"fmt"
	"net/http"
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
			blog := Blog{}
			blog.fetch(_conf)

			template, err := template.New("tpl").ParseGlob(_conf.Template.Dir + "*.html")
			if err != nil {
				fmt.Println(err)
			}

			router := httprouter.New()
			router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

			router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
				template.ExecuteTemplate(w, "index.html", blog)
			})

			fmt.Println("Listening on " + _conf.Server.Host + ":" + _conf.Server.Port + ": ")
			err = http.ListenAndServe(_conf.Server.Host+":"+_conf.Server.Port, router)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}

			return nil
		},
	}
}
