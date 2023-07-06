package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title         string
	Content       string
	Link          string
	Date          string
	EstimatedTime int
	Tags          []string
}

type Blog struct {
	Config      Config
	CurrentPost Post
	Posts       map[string][]Post
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

func (b *Blog) fetch(dir string) error {
	posts := make([]Post, 0)

	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if e.IsDir() {
			b.fetch(path.Join(dir, e.Name()))
		} else {
			post := b.find(path.Join(dir, e.Name()))

			fmt.Println(post.Title + " added to " + dir)
			posts = append(posts, post)
		}
	}

	b.Config = _conf
	if b.Posts == nil {
		b.Posts = make(map[string][]Post)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return strings.Compare(posts[i].Date, posts[j].Date) == 1
	})

	b.Posts[strings.Replace(dir, _conf.Source.Dir, "", 1)] = posts
	return nil
}

func (b *Blog) find(addr string) Post {
	ext := ""
	if !strings.HasSuffix(addr, "."+_conf.Source.Ext) {
		ext += "." + _conf.Source.Ext
	}

	fileAddr := path.Join(addr + ext)

	fs, _ := os.OpenFile(fileAddr, os.O_RDONLY, 0644)
	date, _ := fs.Stat()

	buf, _ := os.ReadFile(fileAddr)

	content := bytes.NewBuffer(mdToHTML(buf)).String()

	fileName := strings.Split(fs.Name(), "/")[len(strings.Split(fs.Name(), "/"))-1]
	fileName = strings.Replace(fileName, "."+_conf.Source.Ext, "", 1)
	post := Post{
		fileName,
		content,
		_conf.Server.Url + path.Join("article", fileName),
		date.ModTime().Format("January 02, 2006 15:04"),
		len(strings.Split(content, " ")) / 250,
		make([]string, 0),
	}

	return post
}

func GetList(addr string) []Post {
	return blog.Posts[addr]
}

func GetPost(addr string) Post {
	return blog.find(path.Join(_conf.Source.Dir, addr))
}

func GetPostContent(addr string) string {
	return GetPost(addr).Content
}

func (blog *Blog) renderIndex(wr io.Writer, tpl *template.Template, page Route) {
	tpl.ParseFiles(path.Join(_conf.Template.Dir, page.Template))

	tpl.ExecuteTemplate(wr, page.Template, blog)
}

func (blog *Blog) renderPost(wr io.Writer, tpl *template.Template, page Route, param string) {
	tpl.ParseFiles(path.Join(_conf.Template.Dir, page.Template))
	blog.CurrentPost = blog.find(path.Join(_conf.Source.Dir, page.Dir, param))

	tpl.ExecuteTemplate(wr, page.Template, blog)
}
