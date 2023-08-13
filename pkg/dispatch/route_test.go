package dispatch

import (
	"fmt"
	"testing"
)

func init() {
	NewRoute(&Route{
		Name:    "root",
		Pattern: "/",
	})

	NewRoute(&Route{
		Name:    "blog-index",
		Pattern: "/blog",
	})

	NewRoute(&Route{
		Name:    "blog-post",
		Pattern: "/blog/{pid}/posts",
	})

	NewRoute(&Route{
		Name:    "blog-posts",
		Pattern: "/blog/posts",
	})

}

func TestSlashRoute(t *testing.T) {
	r, err := Parse("/blog/{pid}/posts")
	fmt.Println(r)
	fmt.Println(err)

	t.Errorf("detected %q, %q was expected", "1", "root")
}
