package config

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Config{}

	err := c.Load("example_config.yaml")
	if c.Name != "My Blog" {
		fmt.Println(err)
		t.Errorf("expected %v, got %v", "My Blog", c.Name)
	}

	if c.Description != "A simple blog about my thoughts and experiences" {
		fmt.Println(err)
		t.Errorf("expected %v, got %v", "A simple blog about my thoughts and experiences", c.Description)
	}
}
