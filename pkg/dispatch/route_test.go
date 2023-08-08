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
		Pattern: "blog/category/it/a",
	}
	r.addToTree(0, &seg)

	r = Route{
		Pattern: "blog/category/food/apple-new-features",
	}
	r.addToTree(0, &seg)

	r = Route{
		Pattern: "blog/category/tech/apple-new-features/1",
	}
	r.addToTree(0, &seg)
	r = Route{
		Pattern: "blog/category/tech/apple-new-features/2",
	}
	r.addToTree(0, &seg)

	url := "blog/category/tech/apple-new-features/2"
	Parse(url, &seg)

	t.Errorf("1")
}
