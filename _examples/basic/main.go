package main

import (
	"log"

	download "github.com/shiningrush/go-download"
)

func main() {

	// no options specified so will default to 10 concurrent download by default

	f, err := download.Open("https://storage.googleapis.com/golang/go1.8.1.src.tar.gz", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// f implements io.Reader, write file somewhere or do some other sort of work with it
}
