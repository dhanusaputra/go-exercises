package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	max := flag.Int("max", 12, "a+b+c")
	flag.Parse()
	for i := 1; i < *max; i++ {
		for j := i + 1; j < *max; j++ {
			for k := j + 1; k < *max; k++ {
				if i+j+k != *max {
					continue
				}
				if isPythagoreanTriplet(float64(i), float64(j), float64(k)) {
					fmt.Printf("%d*%d*%d=%d\n", i, j, k, i*j*k)
				}
			}
		}
	}
}

func isPythagoreanTriplet(a float64, b float64, c float64) bool {
	if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
		return a < b && b < c
	}
	return false
}
