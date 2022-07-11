package main

import (
	"github.com/team-toolchain/tt/internal/args"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args.HandleArgs()
	return 0
}
