package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getFirstFibonacciDigit(10))
}

func getFirstFibonacciDigit(digit int64) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	var lim big.Int
	lim.Exp(big.NewInt(10), big.NewInt(digit-1), nil)
	for a.Cmp(&lim) == -1 {
		a.Add(a, b)
		a, b = b, a

	}
	return a
}
