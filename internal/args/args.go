package args

import (
	"flag"
)

func HandleArgs() {
	configFlag()
	flag.Parse()
}

func configFlag() {
	flag.String("new", "", "initializes a new toolchain")
}
