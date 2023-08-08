package dispatch

import (
	"testing"
)

var routes *Node

func init() {
	routes = &Node{
		IsParameter: false,
		Name:        "/",
	}

	r := Route{
		Name:    "root",
		Pattern: "/",
	}
	r.addToTree(0, routes)

	r = Route{
		Name:    "blog-index",
		Pattern: "/blog",
	}
	r.addToTree(0, routes)

	r = Route{
		Name:    "blog-posts",
		Pattern: "/blog/posts",
	}
	r.addToTree(0, routes)

	r = Route{
		Name:    "blog-post",
		Pattern: "/blog/{pid}/posts",
	}
	r.addToTree(1, routes)

}

func TestSlashRoute(t *testing.T) {
	node := Parse("/", routes)
	if node.Route.Name != "root" {
		t.Errorf("detected %q, %q was expected", node.Route.Name, "root")
	}
}

func TestBlogRoute(t *testing.T) {
	node := Parse("/blog", routes)
	if node.Route.Name != "blog" {
		t.Errorf("detected %q, %q was expected", node.Route.Name, "blog")
	}
}

func TestNestedRoute(t *testing.T) {
	node := Parse("/blog/posts", routes)
	if node.Route.Name != "blog-posts" {
		t.Errorf("detected %q, %q was expected", node.Route.Name, "blog-posts")
	}
}

func TestParameter(t *testing.T) {
	node := Parse("/blog/12/posts", routes)
	if node.Route.Name != "blog-post" {
		t.Errorf("detected %q, %q was expected", node.Route.Name, "blog-post")
	}
}
