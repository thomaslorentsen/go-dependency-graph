package main

import "flag"
import (
	"github.com/artyom/autoflags"
	"fmt"
)

var config = struct {
	Param string `flag:"param,a parameter"`
}{}

func main() {
	autoflags.Define(&config)
	flag.Parse()

	fmt.Printf("%s\n", config.Param)
}