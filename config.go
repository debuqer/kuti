package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string
	Server struct {
		Host string
		Port string
	}
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

	fmt.Println(c)

	return nil
}
