package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Parameter string
	Dir       string
	Template  string
}

type Config struct {
	Name   string
	Author struct {
		Name        string
		Username    string
		Description string
		Profile     string
	}
	Server struct {
		Host string
		Port string
		Url  string
		Ext  string
	}
	Template struct {
		Dir string
	}
	Source struct {
		Dir string
		Ext string
	}
	Routes map[string]Route
}

func (c *Config) load() (err error) {

	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return
	}

	return nil
}
