package cms

import (
	"testing"

	"github.com/debuqer/kuti/internal/config"
	"github.com/debuqer/kuti/pkg/dispatch"
)

var b Blog

func init() {
	c := config.Config{}
	c.Routes = make(map[string]config.Route)

	c.Routes["list"] = config.Route{
		Url:      "/list",
		Template: "list.html",
	}
	c.Routes["post"] = config.Route{
		Url:      "/blog/:pid",
		Template: "pid.html",
	}
	c.Routes["comments"] = config.Route{
		Url:      "/blog/:pid/comments",
		Template: "comment.html",
	}

	dispatch.GET(&dispatch.Route{
		Name:    "root",
		Pattern: "/",
	})

	for rName, r := range c.Routes {
		dispatch.GET(&dispatch.Route{
			Name:    rName,
			Pattern: r.Url,
			CallBack: &dispatch.TemplateCallBack{
				Template: r.Template,
			},
		})
	}
}

func TestUrlCanParse(t *testing.T) {
	r := b.Query("/")
	if r.Route.Name != "root" {
		t.Errorf("excepted %q, got %q", "root", r.Route.Name)
	}

	r = b.Query("/list")
	if r.Route.Name != "list" {
		t.Errorf("excepted %q, got %q", "list", r.Route.Name)
	}

	r = b.Query("/blog/12/comments")
	if r.Route.Name != "comments" {
		t.Errorf("excepted %q, got %q", "comments", r.Route.Name)
	}

	r = b.Query("/blog/12")
	if r.Route.Name != "post" {
		t.Errorf("excepted %q, got %q", "post", r.Route.Name)
	}
}
