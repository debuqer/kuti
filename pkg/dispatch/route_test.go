package dispatch

import "testing"

func TestLinkWhenContainsFileAddress(t *testing.T) {
	l := Link{
		Pattern: "/s1/s2/{ md files in blog\\posts }/post",
	}

	f, a := l.Extend()
	if f != "md" {
		t.Errorf("%q is not equal %q", f, "md")
	}

	if a != "blog\\posts" {
		t.Errorf("%q is not equal %q", f, "blog\\posts")
	}
}

func TestSimpleLink(t *testing.T) {
	l := Link{
		Pattern: "/s1/s2/s3/post",
	}

	f, a := l.Extend()
	if f != "" {
		t.Errorf("%q is not equal %q", f, "")
	}

	if a != "" {
		t.Errorf("%q is not equal %q", f, "")
	}
}

func TestEmptyLink(t *testing.T) {
	l := Link{
		Pattern: "/",
	}

	f, a := l.Extend()
	if f != "" {
		t.Errorf("%q is not equal %q", f, "")
	}

	if a != "" {
		t.Errorf("%q is not equal %q", f, "")
	}
}
