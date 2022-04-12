package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// bytes.Equal(x[:], y[:])
const chunkSize = 4 // 4k

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func getSHA256File(filePath string) ([]byte, error) {
	fileData, err := readFile(filePath)
	hash := sha256.Sum256(fileData)
	// Slice referencing the storage of hash array which has
	// a formal length (32)
	return hash[:], err
}

func findDuplicateHash(slice_of_slices [][]byte) map[string][]int {
	duplicate_position := make(map[string][]int)
	for index, item := range slice_of_slices {
		duplicate_position[string(item)] = append(duplicate_position[string(item)], index)
	}
	return duplicate_position
}

func fileCompareByChunk(file1, file2 string) bool {
	// Check files sizes
	file1Stat, _ := os.Stat(file1)
	file2Stat, _ := os.Stat(file2)
	if file1Stat.Size() != file2Stat.Size() {
		return false
	}

	f1, _ := os.Open(file1)
	defer f1.Close()
	f2, _ := os.Open(file2)
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize*1024)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize*1024)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func fileCompareByBuffer(file1, file2 string) bool {
	// Check files sizes
	file1Stat, _ := os.Stat(file1)
	file2Stat, _ := os.Stat(file2)
	if file1Stat.Size() != file2Stat.Size() {
		return false
	}

	file1Descriptor, _ := os.Open(file1)
	defer file1Descriptor.Close()

	file2Descriptor, _ := os.Open(file2)
	defer file2Descriptor.Close()

	reader1 := bufio.NewReader(file1Descriptor)
	buf1 := make([]byte, 16)

	reader2 := bufio.NewReader(file2Descriptor)
	buf2 := make([]byte, 16)

	res := true

	for {
		n, err1 := reader1.Read(buf1)
		p, err2 := reader2.Read(buf2)

		if (err1 != nil) || (err2 != nil) {

			if err1 != io.EOF {
				log.Fatal(err1)
				res = false
			}

			if err2 != io.EOF {
				log.Fatal(err2)
				res = false
			}

			break
		}
		res = res && (string(buf1[0:n]) == string(buf2[0:p]))
		if !res {
			return res
		}
	}
	return res
}

func downloadFile(URL, fileName string) {
	start := time.Now()
	response, _ := http.Get(URL)
	defer response.Body.Close()
	elapsed := time.Since(start)

	log.Printf("Getting image %s took: %s", URL, elapsed)

	if response.StatusCode != 200 {
		log.Fatalf("Received non 200 response code: %d", response.StatusCode)
	}

	file, _ := os.Create(fileName)
	defer file.Close()

	// Write the body to file
	io.Copy(file, response.Body)
}

func printResultBools(t1 bool, t2 bool, t3 bool) {
	if t1 {
		fmt.Println("image_1.jpg and image_2.jpg are the same")
	}
	if t2 {
		fmt.Println("image_1.jpg and image_3.jpg are the same")
	}
	if t3 {
		fmt.Println("image_2.jpg and image_3.jpg are the same")
	}
}

func handleUrlFlag() {
	var urls string
	urlsSlice := make([]string, 0)
	flag.StringVar(&urls, "urls", "", "Comma separated list of image URLs")
	flag.Parse()
	if urls != "" {
		urlsSlice = strings.Split(urls, ",")
	}
	for i, url := range urlsSlice {
		downloadFile(url, fmt.Sprintf("chat_mignon_%d.jpg", i))
	}
}

func compareBySha256Complex() [][]byte {
	imageHashes := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		fileName := fmt.Sprintf("image_%d.jpg", i+1)
		var err error
		imageHashes[i], err = getSHA256File(fmt.Sprint(fileName))
		if err != nil {
			log.Fatalf("Error while getting SHA256 of file %s", fileName)
		}
	}
	return imageHashes
}

func compareBySha256Simple() {
	content, err := ioutil.ReadFile("image_1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	content1, err1 := ioutil.ReadFile("image_2.jpg")
	if err1 != nil {
		log.Fatal(err1)
	}
	content2, err2 := ioutil.ReadFile("image_3.jpg")
	if err2 != nil {
		log.Fatal(err2)
	}
	sum := sha256.Sum256(content)
	sum1 := sha256.Sum256(content1)
	sum2 := sha256.Sum256(content2)

	equal1 := sum == sum1
	equal2 := sum == sum2
	equal3 := sum1 == sum2

	fmt.Println("image 1 et 2: ", equal1)
	fmt.Println("image 1 et 3: ", equal2)
	fmt.Println("image 2 et 3: ", equal3)
}

func printResult256(duplicateImageHashes map[string][]int) {
	for _, v := range duplicateImageHashes {
		if len(v) > 1 {
			fmt.Printf("Images of index %v are the same\n", v)
		} else {
			fmt.Printf("Image 'image_%d.jpg' is uniq\n", v[0]+1)
		}
	}
}

func setupLogger() error {
	log.SetOutput(os.Stderr)
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("logger configured")
	var err error
	return err
}

func main() {
	// Logger
	setupLogger()

	// Flags
	// $ go run main.go -urls "https://cdn.radiofrance.fr/s3/cruiser-production/2020/08/f76e923a-0072-4c0d-9837-7bf89d315133/1136_les_chats.webp,https://i-mom.unimedias.fr/2022/03/14/chat.jpg?auto=format%2Ccompress&crop=faces&cs=tinysrgb&fit=crop&h=453&w=806"
	handleUrlFlag()

	// Compare by sha256
	log.Println("Compare by sha256 Complex")
	imageHashes := compareBySha256Complex()
	duplicateImageHashes := findDuplicateHash(imageHashes)
	printResult256(duplicateImageHashes)

	log.Println("Compare by sha256 Simple")
	compareBySha256Simple()

	// Compare by chunks
	log.Println("Compare by chunks")
	t1 := fileCompareByChunk("image_1.jpg", "image_2.jpg")
	t2 := fileCompareByChunk("image_1.jpg", "image_3.jpg")
	t3 := fileCompareByChunk("image_3.jpg", "image_2.jpg")
	printResultBools(t1, t2, t3)

	// Compare by buffer
	log.Println("Compare by buffer")
	t1 = fileCompareByBuffer("image_1.jpg", "image_2.jpg")
	t2 = fileCompareByBuffer("image_1.jpg", "image_3.jpg")
	t3 = fileCompareByBuffer("image_3.jpg", "image_2.jpg")
	printResultBools(t1, t2, t3)
}
