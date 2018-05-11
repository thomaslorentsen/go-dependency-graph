package main

import "flag"
import (
	"github.com/artyom/autoflags"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config = struct {
	Param string `flag:"repo,a parameter"`
}{}

type Graph struct {
	Repo string `yaml:"repo"`
}

func main() {
	autoflags.Define(&config)
	flag.Parse()

	g := new(Graph)
	g.Repo = config.Param
	y, _ := yaml.Marshal(&g)
	ioutil.WriteFile("graph.yml", y, 0644)
	fmt.Printf("Written graph.yml:\n---\n%s---\n", y)
}