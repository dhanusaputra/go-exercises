// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

package main

import (
	"fmt"
	"math"
)

func main() {
	var abundants []float64
	for i := 1; i <= 28123; i++ {
		if isAbundantNumber(float64(i)) {
			abundants = append(abundants, float64(i))
		}
	}
	abundantSums := getSumBetweenSlices(abundants, abundants)
	var sum float64
	for i := 1; i <= 28123; i++ {
		if !sliceContains(abundantSums, float64(i)) {
			sum += float64(i)
		}
	}
	fmt.Printf("%.f\n", sum)
}

func getDivisors(f float64) []float64 {
	var res []float64
	for i := 1; i < int(f); i++ {
		if math.Mod(f, float64(i)) == 0 {
			res = append(res, float64(i))
		}
	}
	return res
}

func isAbundantNumber(f float64) bool {
	divisors := getDivisors(f)
	var sum float64
	for _, v := range divisors {
		sum += v
	}
	return sum > f
}

func getSumBetweenSlices(s1 []float64, s2 []float64) []float64 {
	var res []float64
	for _, i := range s1 {
		for _, j := range s2 {
			res = append(res, i+j)
		}
	}
	return res
}

func sliceContains(s []float64, f float64) bool {
	for _, v := range s {
		if v == f {
			return true
		}
	}
	return false
}
