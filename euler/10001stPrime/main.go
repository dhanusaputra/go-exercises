package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.Int("x", 10, "Prime number")
	flag.Parse()
	c := 0
	i := 1
	for c < *input {
		i++
		if isPrime(i) {
			c++
		}
	}
	fmt.Println(i)
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}
