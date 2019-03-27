package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const url = "https://google.com"

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	urls := os.Args[1:]

	for _, url := range urls {
		data, err := getURL(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully fetched %d bytes for %s\n", len(data), url)
	}
}

func usage() {
	prog := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s URL1 [URL2 ...]\n", prog)
}

func getURL(url string) ([]byte, error) {
	fmt.Printf("GET %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
