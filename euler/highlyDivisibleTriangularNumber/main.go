package main

import (
	"flag"
	"fmt"
)

func main() {
	max := flag.Int("max", 5, "number of maximum divisors")
	flag.Parse()
	i := 1
	for {
		if len(getDivisors(getTringularNumber(i))) > *max {
			break
		}
		i++
	}
	fmt.Println(getTringularNumber(i))
}

func getTringularNumber(n int) int {
	return n * (n + 1) / 2
}

func getDivisors(in int) []int {
	var res []int
	for i := 1; i <= in; i++ {
		if in%i == 0 {
			res = append(res, i)
		}
	}
	return res
}
