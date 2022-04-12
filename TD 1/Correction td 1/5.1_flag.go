package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0"

func main() {
	version := flag.Bool("version", false, "Prints program version")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	fmt.Println("Hello world")
}
