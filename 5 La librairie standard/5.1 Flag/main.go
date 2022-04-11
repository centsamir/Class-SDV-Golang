package main

import (
	"flag"
	"fmt"
)

func main() {
	const VERSION = "1.0"

	var nFlag = flag.Bool("version", false, "Show version")
	flag.Parse()

	if *nFlag == true {
		fmt.Println(VERSION)
	}
}
