package main

import (
	"fmt"
	"github.com/team-toolchain/tt/internal/args"
	"github.com/team-toolchain/tt/internal/logo"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fmt.Println(logo.Show())
	args.HandleArgs()
	return 0
}
