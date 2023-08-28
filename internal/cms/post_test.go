package cms

import (
	"testing"
)

func TestCreatePostFromFile(t *testing.T) {
	post, _ := MakePostFromFile("contents/test.md")

	if post.GetTitle() != "test" {
		t.Errorf("expected %v, got %v", "test", post.GetTitle())
	}
	if post.FileName != "test.md" {
		t.Errorf("expected %v, got %v", "test", "test.md")
	}
	if post.Content != "Hello\n\nthis is test" {
		t.Errorf("expected %v, got %v", "test", "Hello\n\nthis is test")
	}
}

func TestCreatePostFromFileWhenNotExists(t *testing.T) {
	p, err := MakePostFromFile("contents/test_not_exists.md")

	if err == nil {
		t.Errorf("No error when file not exists")
	}
	if p.FileName != "" {
		t.Errorf("filename filled while file not exists")
	}
}
