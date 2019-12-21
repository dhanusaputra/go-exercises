package main

import (
	"flag"
	"fmt"
)

func main() {
	max := flag.Int("max", 10, "maximum amount of prime")
	flag.Parse()
	var res int
	for i := 2; i < *max; i++ {
		if isPrime(i) {
			res += i
		}
	}
	fmt.Println(res)
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num >= 2
}
