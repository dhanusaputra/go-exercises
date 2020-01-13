package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	i := flag.Int64("input", 10, "factorial number")
	flag.Parse()
	factorial, err := getFactorial(*i)
	if err != nil {
		panic(err)
	}
	sum, err := sumOfCharDigits(factorial.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)
}

func getFactorial(i int64) (big.Int, error) {
	var b big.Int
	b.MulRange(1, i)
	return b, nil
}

func sumOfCharDigits(s string) (int, error) {
	var res int
	for _, ss := range s {
		tmp, err := strconv.Atoi(string(ss))
		if err != nil {
			panic(err)
		}
		res += tmp
	}
	return res, nil
}
