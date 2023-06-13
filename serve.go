package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/urfave/cli"
)

func ServeCommand(_conf Config) cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serves the application",
		Action: func(c *cli.Context) error {
			template := template.Must(template.New("tpl").ParseGlob(_conf.Template.Dir + "*.html"))

			fs := http.FileServer(http.Dir("./assets"))
			http.Handle("/assets/", http.StripPrefix("/assets/", fs))

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				template.ExecuteTemplate(w, "base.html", nil)
			})

			fmt.Println("Listening on " + _conf.Server.Host + ":" + _conf.Server.Port + ":")
			err := http.ListenAndServe(_conf.Server.Host+":"+_conf.Server.Port, nil)
			if err != nil {
				fmt.Println("Template file assets/config.yml not found")
				return nil
			}

			return nil
		},
	}
}
