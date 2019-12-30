package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	grid := flag.Int("grid", 2, "number of grid")
	flag.Parse()
	fmt.Printf("%.f\n", countingLaticePaths(*grid))
}

func countingLaticePaths(n int) float64 {
	var nominator, denominator float64
	nominator, denominator = 1, 1
	for i := 1; i <= n*2; i++ {
		nominator *= float64(i)
	}
	for i := 1; i <= n; i++ {
		denominator *= float64(i)
	}
	denominator = math.Pow(denominator, 2)
	return nominator / denominator
}
