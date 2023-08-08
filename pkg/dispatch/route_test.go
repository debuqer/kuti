package dispatch

import (
	"fmt"
	"testing"
)

func TestRoute(t *testing.T) {
	seg := Segment{
		IsParameter: false,
		Name:        "/",
	}

	r := Route{
		Pattern: "blog/posts/{post-number}/page",
	}
	r.addToTree(0, &seg)

	r = Route{
		Pattern: "blog/posts/hello",
	}
	r.addToTree(0, &seg)

	url := "blog/posts/1/page"
	fmt.Println(Parse(url, &seg))

	t.Errorf("1")
}
