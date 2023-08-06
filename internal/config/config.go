package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Cfg Config

type Route struct {
	Type      string
	Parameter string
	Dir       string
	Template  string
	Paginate  int
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

func (c *Config) Load() (err error) {

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
