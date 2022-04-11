package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func main() {
	fmt.Println(readFile("image_1.jpg"))
}
