package cms

import (
	"testing"

	"github.com/debuqer/kuti/internal/config"
	"github.com/debuqer/kuti/pkg/dispatch"
)

var b Blog

func init() {
	c := config.Config{}
	c.Routes = append(c.Routes, config.Route{
		Url: "/list",
	})
	c.Routes = append(c.Routes, config.Route{
		Url: "/blog/:pid",
	})
	c.Routes = append(c.Routes, config.Route{
		Url: "/blog/:pid/comments",
	})

	dispatch.GET(&dispatch.Route{
		Name:    "root",
		Pattern: "/",
	})

	for _, r := range c.Routes {
		dispatch.GET(&dispatch.Route{
			Name:    r.Url,
			Pattern: r.Url,
		})
	}
}

func TestUrlCanParse(t *testing.T) {
	r := b.Query("/")
	if r.Route.Name != "root" {
		t.Errorf("excepted %q, got %q", "root", r.Route.Name)
	}
}
