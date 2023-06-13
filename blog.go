package main

import (
	"bytes"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func (b *Blog) fetch(_conf Config) error {
	posts := make([]Post, 0)

	entries, _ := os.ReadDir(_conf.Source.Dir)
	for _, e := range entries {
		if !e.IsDir() {
			fs, _ := os.OpenFile(_conf.Source.Dir+"/"+e.Name(), os.O_RDONLY, 0644)
			date, _ := fs.Stat()

			buf, _ := os.ReadFile(_conf.Source.Dir + "/" + e.Name())

			content := bytes.NewBuffer(mdToHTML(buf)).String()

			post := Post{
				e.Name(),
				content,
				date.ModTime().Format("January 02, 2006 15:04"),
				len(strings.Split(content, " ")) / 250,
			}

			posts = append(posts, post)
		}
	}

	b.Config = _conf
	b.Posts = posts

	return nil
}
