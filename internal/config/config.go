package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Cfg Config

type Author struct {
	Name         string
	Bio          string
	ProfileImage string
}

type Tech struct {
	Url        string
	ThemeDir   string
	ContentDir string
}

type Route struct {
	Url      string
	Template string
}

type Config struct {
	Name        string
	Description string
	Routes      []Route
	Author      Author

	Tech Tech
}

func (c *Config) Load(fname string) (err error) {

	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return
	}

	return nil
}
