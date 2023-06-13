package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/urfave/cli"
)

type Post struct {
	Title         string
	Content       string
	Date          string
	EstimatedTime int
}

type Blog struct {
	Config Config
	Posts  []Post
}

func (b *Blog) fetch(_conf Config) error {
	post := Post{
		"Hello",
		"Hoy",
		"12/02/1402",
		1,
	}

	posts := make([]Post, 0)

	posts = append(posts, post)

	b.Config = _conf
	b.Posts = posts

	return nil
}

func ServeCommand(_conf Config) cli.Command {
	return cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serves the application",
		Action: func(c *cli.Context) error {
			blog := Blog{}
			blog.fetch(_conf)

			template := template.Must(template.New("tpl").ParseGlob(_conf.Template.Dir + "*.html"))

			fs := http.FileServer(http.Dir("./assets"))
			http.Handle("/assets/", http.StripPrefix("/assets/", fs))

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				template.ExecuteTemplate(w, "index.html", blog)
			})

			fmt.Println("Listening on " + _conf.Server.Host + ":" + _conf.Server.Port + ": ")
			err := http.ListenAndServe(_conf.Server.Host+":"+_conf.Server.Port, nil)
			if err != nil {
				fmt.Println("Template file assets/config.yml not found")
				return nil
			}

			return nil
		},
	}
}
