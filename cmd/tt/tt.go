package main

import (
	"flag"
)

func configFlag() {
	flag.String("new", "", "initializes a new toolchain")
}

func main() {
	flag.Parse()
}
