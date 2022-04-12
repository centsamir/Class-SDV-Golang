package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0"

func main() {
	var version *bool
	*version = false
	flag.BoolVar(version, "version", false, "Prints program version")
	flag.BoolVar(version, "v", false, "Prints program version")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	fmt.Println("Hello world")
}
