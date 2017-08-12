package main

import (
	"os"
)

const Name string = "friction"
const Version string = "v0.1.0"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
