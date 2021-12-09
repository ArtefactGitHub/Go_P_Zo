package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("hello world")

	confContent, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		panic(err)
	}
	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}
	fmt.Printf("config: %v\n", conf)
}

type Config struct {
	Environment string `yaml:"environment"`
	Key         string `yaml:"key"`
}
