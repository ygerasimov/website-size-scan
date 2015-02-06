package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var count int
var workers int
var baseUrl string
var inputFile string
var outputFile string

func init() {
	flag.IntVar(&count, "count", 100, "Number of URLs in the input file")
	flag.IntVar(&workers, "workers", 5, "Number of workers. Be careful with this number. Big number will put your site down.")
	flag.StringVar(&baseUrl, "baseUrl", "http://www.example.com", "Will be used to adjust relative paths.")
	flag.StringVar(&inputFile, "inputFile", "./links.csv", "Will be used as source file for links.")
	flag.StringVar(&outputFile, "outputFile", "./output.csv", "Will be used as output for results.")

	flag.Parse()
}

func main() {

	input := make(chan string)
	output := make(chan string)

	// Read the file and feed links to input channel.
	go func(count int, in chan string) {
		file, err := os.Open(inputFile)
		checkError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			in <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			checkError(err)
		}
	}(count, input)

	// Run workers to process each of urls from input channel.
	for j := 1; j <= workers; j++ {
		go func(k int, in chan string, out chan string) {
			for {
				url := <-in
				fmt.Println("Worker", k, "start scan:", url)

				doc, err := goquery.NewDocument(url)
				if err != nil {
					log.Fatal(err)
				}

				totalSize := getUrlSize(url, baseUrl)

				// Extract images and check their size.
				// fmt.Printf("Document %s\n", url)
				doc.Find("body img").Each(func(i int, s *goquery.Selection) {
					src, ok := s.Attr("src")
					// fmt.Printf("Image %s\n", src)
					if ok {
						totalSize = totalSize + getUrlSize(src, baseUrl)
					}
				})

				// Extract css size.
				doc.Find("link").Each(func(i int, s *goquery.Selection) {
					rel, _ := s.Attr("rel")
					if rel == "stylesheet" {
						href, ok := s.Attr("href")
						// fmt.Printf("CSS %s\n", href)
						if ok {
							totalSize = totalSize + getUrlSize(href, baseUrl)
						}
					}
				})

				// Extract javascript size.
				doc.Find("script").Each(func(i int, s *goquery.Selection) {
					scriptType, _ := s.Attr("type")
					if scriptType == "text/javascript" {
						src, ok := s.Attr("src")
						// fmt.Printf("JS %s\n", src)
						if ok {
							totalSize = totalSize + getUrlSize(src, baseUrl)
						}
					}
				})

				var totalSizeFloat float64
				totalSizeFloat = float64(totalSize) / (1024 * 1024)

				out <- fmt.Sprintf("%s, %s", url, strconv.FormatFloat(totalSizeFloat, 'f', 2, 64))
			}
		}(j, input, output)
	}

	f, err := os.Create(outputFile)
	checkError(err)
	defer f.Close()

	for i := 1; i <= count; i++ {
		result := <-output
		n, err := f.WriteString(result + "\n")
		if err != nil {
			fmt.Println(n, err)
		}
		fmt.Println(i, result)
	}

	f.Sync()
}

func getUrlSize(url, baseUrl string) int64 {
	if url == "" {
		return 0
	}

	// If url starts with // (like google font)
	if url[0:2] == "//" {
		url = "http:" + url
	}

	// If url is relative.
	if !strings.Contains(url, "://") {
		url = baseUrl + url
	}

	response, err := http.Head(url)
	if err != nil {
		log.Println("Error while downloading", url, ":", err)
		return 0
	}
	// Verify if the response was ok
	if response.StatusCode != http.StatusOK {
		log.Println("Server return non-200 status: ", url, response.Status)
		return 0
	}

	length, _ := strconv.Atoi(response.Header.Get("Content-Length"))
	sourceSize := int64(length)

	return sourceSize
}

func checkError(e error) {
	if e != nil {
		fmt.Printf("%s", e)
		panic(e)
	}
}
