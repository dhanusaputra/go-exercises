// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.

package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	n := flag.Int("n", 10000, "sum of all the amicable numbers under ...")
	flag.Parse()

	divisors := make(map[float64][]float64)
	for i := 1; i < *n; i++ {
		divisors[float64(i)] = getDivisors(float64(i))
	}
	amicables := make(map[float64]float64)
	for k, v := range divisors {
		sum := getSumOfArray(v)
		if val, ok := divisors[sum]; ok && k == getSumOfArray(val) && k != sum {
			amicables[k] = sum
		}
	}
	var sumAmicables float64
	for _, v := range amicables {
		sumAmicables += v
	}
	fmt.Println(sumAmicables)
}

func getDivisors(n float64) []float64 {
	var res []float64
	for i := 1; i < int(n); i++ {
		if math.Mod(n, float64(i)) == 0 {
			res = append(res, float64(i))
		}
	}
	return res
}

func getSumOfArray(s []float64) float64 {
	var res float64
	for _, v := range s {
		res += v
	}
	return res
}
