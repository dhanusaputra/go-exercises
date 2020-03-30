package main

import (
	"io"
	"os"

	"github.com/dhanusaputra/go-exercises/transform/primitive"
)

func main() {
	f, err := os.Open("tmp/test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	out, err := primitive.Transform(f, 50)
	if err != nil {
		panic(err)
	}
	os.Remove("out.png")
	outFile, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	io.Copy(outFile, out)
}
