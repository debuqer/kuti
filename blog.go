package main

import (
	"bytes"
	"os"
	"path"
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
	Tags          []string
}

type Blog struct {
	Config      Config
	CurrentPost Post
	Posts       []Post
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
			post := b.find(_conf, e.Name())

			posts = append(posts, post)
		}
	}

	b.Config = _conf
	b.Posts = posts

	return nil
}

func (b *Blog) find(_conf Config, title string) Post {
	ext := ""
	if !strings.HasSuffix(title, "."+_conf.Source.Ext) {
		ext += "." + _conf.Source.Ext
	}

	fileAddr := path.Join(_conf.Source.Dir + "/" + title + ext)

	fs, _ := os.OpenFile(fileAddr, os.O_RDONLY, 0644)
	date, _ := fs.Stat()

	buf, _ := os.ReadFile(fileAddr)

	content := bytes.NewBuffer(mdToHTML(buf)).String()

	post := Post{
		title,
		content,
		date.ModTime().Format("January 02, 2006 15:04"),
		len(strings.Split(content, " ")) / 250,
		make([]string, 0),
	}

	return post
}
