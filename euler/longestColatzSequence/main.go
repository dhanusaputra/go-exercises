package main

import (
	"flag"
	"fmt"
)

func main() {
	max := flag.Int("max", 13, "maximum sequence")
	flag.Parse()
	var result, tmp []int
	for i := 2; i < *max; i++ {
		tmp = getColatz(i)
		if len(tmp) > len(result) {
			result = tmp
		}
	}
	fmt.Println(result[0], len(result))
}

func getColatz(n int) []int {
	var r []int
	r = append(r, n)
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		r = append(r, n)
	}
	return r
}
