package cms

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/debuqer/kuti/internal/config"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Blog struct {
	Config      config.Config
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

func (b *Blog) Fetch(dir string) error {
	posts := make([]Post, 0)

	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if e.IsDir() {
			b.Fetch(path.Join(dir, e.Name()))
		} else {
			post := b.Find(path.Join(dir, e.Name()))

			fmt.Println(post.Title + " added to " + dir)
			posts = append(posts, post)
		}
	}

	b.Config = config.Cfg
	if b.Posts == nil {
		b.Posts = make(map[string][]Post)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return strings.Compare(posts[i].Date, posts[j].Date) == 1
	})

	b.Posts[strings.Replace(dir, config.Cfg.Tech.ContentDir, "", 1)] = posts
	return nil
}

func (b *Blog) Find(addr string) Post {
	ext := ""
	if !strings.HasSuffix(addr, ".md") {
		ext += ".md"
	}

	fileAddr := path.Join(addr + ext)

	fs, _ := os.OpenFile(fileAddr, os.O_RDONLY, 0644)
	date, _ := fs.Stat()

	buf, _ := os.ReadFile(fileAddr)

	content := bytes.NewBuffer(mdToHTML(buf)).String()

	fileName := strings.Split(fs.Name(), "/")[len(strings.Split(fs.Name(), "/"))-1]
	fileName = strings.Replace(fileName, ".md", "", 1)
	post := Post{
		fileName,
		content,
		config.Cfg.Tech.Url + path.Join("article", fileName),
		date.ModTime().Format("January 02, 2006 15:04"),
		len(strings.Split(content, " ")) / 250,
		make([]string, 0),
	}

	return post
}

func (blog *Blog) GetPostList(addr string) []Post {
	return blog.Posts[addr]
}

func (blog *Blog) GetRootPosts() []Post {
	return blog.GetPostList("")
}

func (blog *Blog) GetPost(addr string) Post {
	return blog.Find(path.Join(config.Cfg.Tech.ContentDir, addr))
}

func (blog *Blog) GetPostContent(addr string) string {
	return blog.GetPost(addr).Content
}

func (blog *Blog) GetCurrentPost() Post {
	return blog.CurrentPost
}

func (blog *Blog) RenderIndex(wr io.Writer, tpl *template.Template, page config.Route) {
	tpl.ParseFiles(path.Join(config.Cfg.Tech.ThemeDir, page.Template))

	tpl.ExecuteTemplate(wr, page.Template, blog)
}

func (blog *Blog) RenderPost(wr io.Writer, tpl *template.Template, page config.Route, param string) {
	tpl.ParseFiles(path.Join(config.Cfg.Tech.ThemeDir, page.Template))
	blog.CurrentPost = blog.Find(path.Join(config.Cfg.Tech.ContentDir, page.Template, param))

	tpl.ExecuteTemplate(wr, page.Template, blog)
}
