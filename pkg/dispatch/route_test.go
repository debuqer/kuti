package dispatch

import (
	"math/rand"
	"testing"
	"time"
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

func BenchmarkRoutesDefine(b *testing.B) {

	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())

		// Getting random character
		c := rand.Int31()
		GET(&Route{
			Name:    "blog-index",
			Pattern: "/" + string(c),
		})
	}
}

func BenchmarkRoutesLoopup(b *testing.B) {
	var c int32
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())

		// Getting random character
		c = rand.Int31()

		GET(&Route{
			Name:    "blog-index",
			Pattern: "/" + string(c),
		})
	}

	Parse("/blog/" + string(c))
}
