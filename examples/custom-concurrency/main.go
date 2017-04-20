package main

import (
	"log"

	download "github.com/joeybloggs/go-download"
)

func main() {

	options := &download.Options{
		Concurrency: func(size int64) int64 {
			// break it up into 1MB chunked downloads
			return size / 1000000
		},
	}

	f, err := download.Open("https://storage.googleapis.com/golang/go1.8.1.src.tar.gz", options)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// f implements io.Reader, write file somewhere or do some other sort of work with it
}