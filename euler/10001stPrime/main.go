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
		var tmp bool
		tmp, err := isPrime(i)
		if err != nil {
			panic(err)
		}
		if tmp {
			c++
		}
	}
	fmt.Println(i)
}

func isPrime(n int) (bool, error) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return n >= 2, nil
}
