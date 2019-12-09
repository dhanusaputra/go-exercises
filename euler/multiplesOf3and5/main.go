package main

import (
	"flag"
	"fmt"
)

func main() {
	number := flag.Int("n", 1000, "the inputted number")
	flag.Parse()

	fmt.Printf("The number: %d\n", *number)

	result := 0
	for i := 1; i < *number; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Printf("%d\n", i)
			result += i
		}
	}
	fmt.Printf("The result: %d\n", result)
}
