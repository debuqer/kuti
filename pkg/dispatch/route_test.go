package dispatch

import "testing"

func TestLinkWhenContainsFileAddress(t *testing.T) {
	l := Link{
		Pattern: "/s1/{ files in blog\\posts }/post",
	}

	l.Extend()

	if l.Parameters["0"].Name != "blog\\posts" {
		t.Errorf("%q is not equal %q", l.Parameters["0"].Name, "blog\\posts")
	}
}

func TestLinkMultipleParameters(t *testing.T) {
	l := Link{
		Pattern: "/s1/{ files in blog\\posts }/post/{ files in blog\\videos }",
	}

	l.Extend()

	if l.Parameters["0"].Name != "blog\\posts" {
		t.Errorf("%q is not equal %q", l.Parameters["0"].Name, "blog\\posts")
	}
	if l.Parameters["1"].Name != "blog\\videos" {
		t.Errorf("%q is not equal %q", l.Parameters["1"].Name, "blog\\videos")
	}
}

func TestSimpleLink(t *testing.T) {
	l := Link{
		Pattern: "/s1/s2/s3/post",
	}

	l.Extend()

	if l.HasParameter {
		t.Errorf("Not parameter should be assigned")
	}
}

func TestEmptyLink(t *testing.T) {
	l := Link{
		Pattern: "/",
	}

	l.Extend()

	if l.HasParameter {
		t.Errorf("Not parameter should be assigned")
	}
}
