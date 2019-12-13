package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("input", 13195, "prime number input")
	flag.Parse()
	var res int
	i := 2
	for i <= *n {
		if isPrime(i) && *n%i == 0 {
			*n /= i
			res = i
		}
		i++
	}
	fmt.Println(res)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
