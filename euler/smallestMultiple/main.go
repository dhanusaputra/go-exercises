package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.Int("i", 10, "Number of smallest multiple")
	flag.Parse()
	result := *input
	for {
		if b, err := isEvenlyDivisible(result, *input); b && err == nil {
			break
		}
		result++
	}
	fmt.Println(result)
}

func isEvenlyDivisible(input int, div int) (bool, error) {
	if input < 1 {
		return false, fmt.Errorf("Error number is not positive")
	}
	for i := div; i > 0; i-- {
		if input%i != 0 {
			return false, nil
		}
	}
	return true, nil
}
