package config

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Config{}

	_ = c.Load("example_config.yaml")
	if c.Name != "My Blog" {
		t.Errorf("expected %v, got %v", "My Blog", c.Name)
	}

	if c.Description != "A simple blog about my thoughts and experiences" {
		t.Errorf("expected %v, got %v", "A simple blog about my thoughts and experiences", c.Description)
	}

	if c.Routes["post"].Url != "/post/:slug" {
		t.Errorf("expected %v, got %v", "/post/:slug", c.Routes["post"].Url)
	}

	fmt.Println(c.Routes["post"])
	if c.Routes["post"].Lookup != "post" {
		t.Errorf("expected %v, got %v", "post", c.Routes["post"].Lookup)
	}
}
