package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	n := flag.Int("n", 10, "Natural numbers")
	flag.Parse()
	fmt.Printf("%f - %f = %f\n", squareOfSum(*n), sumOfSquare(*n), squareOfSum(*n)-sumOfSquare(*n))
}

func sumOfSquare(input int) float64 {
	var output float64
	for i := 1; i <= input; i++ {
		output += math.Pow(float64(i), 2)
	}
	return output
}

func squareOfSum(input int) float64 {
	var output float64
	for i := 1; i <= input; i++ {
		output += float64(i)
	}
	output = math.Pow(output, 2)
	return output
}
