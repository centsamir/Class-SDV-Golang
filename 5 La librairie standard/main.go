package main

import (
	"flag"
	"fmt"
)

func main() {
	const VERSION = "1.0"

	var nFlag = flag.Bool("version", false, "test")
	flag.Parse()

	if *nFlag == true {
		fmt.Println("Version is ", VERSION)
	}
}
