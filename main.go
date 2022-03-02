package main

import (
    "os"
	"github.com/s4pv/learnBC/cli"

)

func main() {
    defer os.Exit(0)

    cmd := cli.CommandLine{}

    cmd.Run()
}