package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	//"io/ioutil"
)

func checkError( e error){
	if e != nil {
		fmt.Printf("%s", e)
		panic(e)
	}
}

func main() {
	input := make(chan string) // no buffer
	output := make(chan string)
	count := 5000 
	workers := 20
	path := "/home/ygerasimov/go/";

	// Read the file and feed links to input channel.
	go func(count int, in chan string) {
		file, err := os.Open("links.txt")
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
				
				cmd := path + "website-size.sh -l 1 " + url;
				cmd_out, err := exec.Command("/bin/sh", "-c", cmd).Output()
				checkError(err)
				
				out <- fmt.Sprintf("%s", cmd_out)			
			}
		}(j, input, output)
	}

	f, err := os.Create(path + "/output.txt")
	checkError(err)
	defer f.Close()
		
	for i := 1; i <= count; i++ {
		result := <-output
		n, err := f.WriteString(result)
		if err != nil {
			fmt.Println(n, err)
		}
		fmt.Println(i, result)
	}
	
	f.Sync()
}
