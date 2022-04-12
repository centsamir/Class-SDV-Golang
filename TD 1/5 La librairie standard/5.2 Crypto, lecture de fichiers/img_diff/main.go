package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func hashingFile(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}

func main() {
	//fmt.Println(readFile("image_1.jpg"))
	//fmt.Println(hashingFile("image_1.jpg"))
	//images := os.Args[1:]
	//if len(images) <= 0 {
	//	images = []string{"./images/*"}
	//}
	//fmt.Println(images)

	log.SetOutput(os.Stderr)
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate)
	log.Println("logger configured")
	var err error
	fmt.Println(err)

}
