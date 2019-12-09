package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhanusaputra/go-exercises/adventure"
)

func main() {
	filename := flag.String("file", "/home/dhanu/go/src/github.com/dhanusaputra/go-exercises/adventure/gopher.json", "the JSON file with CYOA")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := adventure.JSONStory(f)

	fmt.Printf("%+v\n", story)
}
