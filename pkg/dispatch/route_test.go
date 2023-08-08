package dispatch

import (
	"testing"
)

func TestRoute(t *testing.T) {
	seg := Segment{
		IsParameter: false,
		Name:        "/",
	}

	r := Route{
		Pattern: "blog/category/tech/apple-new-features",
	}
	r.addToTree(0, &seg)

	t.Errorf("1")
}
