package args

import (
	flag "github.com/spf13/pflag"
)

func HandleArgs() {
	configFlag()
	flag.Parse()
}

func configFlag() {
	flag.String("new", ".", "initializes a new toolchain")
}
