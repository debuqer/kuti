package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string
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
