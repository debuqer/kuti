package dispatch

import (
	"testing"
)

func init() {
	GET(&Route{
		Name:    "root",
		Pattern: "/",
	})

	GET(&Route{
		Name:    "blog-index",
		Pattern: "/blog",
	})

	GET(&Route{
		Name:    "blog-posts",
		Pattern: "/blog/posts",
	})

	GET(&Route{
		Name:    "blog-post",
		Pattern: "/blog/{pid}/posts",
	})

}

func TestSlashRoute(t *testing.T) {
	r, _ := Parse("/")

	if r.Route.Name != "root" {
		t.Errorf("%q expected, but %q given", r.Route.Name, "root")
	}
}

func TestEmptyRoute(t *testing.T) {
	r, _ := Parse("")

	if r.Route.Name != "root" {
		t.Errorf("%q expected, but %q given", r.Route.Name, "root")
	}
}

func TestBlogRoute(t *testing.T) {
	r, _ := Parse("/blog")

	if r.Route.Name != "blog-index" {
		t.Errorf("%q expected, but %q given", r.Route.Name, "root")
	}
}

func TestBlogPostsRoute(t *testing.T) {
	r, _ := Parse("/blog/posts")

	if r.Route.Name != "blog-posts" {
		t.Errorf("%q expected, but %q given", r.Route.Name, "root")
	}
}

func TestBlogPostRoute(t *testing.T) {
	r, _ := Parse("/blog/12/posts")

	if r.Route.Name != "blog-post" {
		t.Errorf("%q expected, but %q given", r.Route.Name, "root")
	}
}
