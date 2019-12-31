package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	smallNumbers = []string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}
	tens = []string{
		"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}
)

func main() {
	var res int
	for i := 1; i <= 1000; i++ {
		res += len(strings.Replace(numToWords(float64(i)), " ", "", -1))
	}
	fmt.Println(res)
}

func numToWords(n float64) string {
	var isNegative bool
	var res string
	if n < 0 {
		n = math.Abs(n)
		isNegative = true
	}
	if n == 0 {
		return "zero"
	}
	if n < 20 {
		res = smallNumbers[int(n)]
	} else if n < 100 {
		res = tens[int(math.Mod(n/10, 10))] + " " + smallNumbers[int(math.Mod(n, 10))]
	} else if n < 1000 {
		if int(math.Mod(n/10, 10)) >= 2 {
			res = smallNumbers[int(math.Mod(n/100, 10))] + " hundred and " + tens[int(math.Mod(n/10, 10))] + " " + smallNumbers[int(math.Mod(n, 10))]
		} else {
			res = smallNumbers[int(math.Mod(n/100, 10))] + " hundred and " + smallNumbers[int(n)-int(math.Mod(n/100, 10))*100]
		}
	} else if n == 1000 {
		res = "one thousand"
	} else {
		return "Error, number below -1000 or above 1000"
	}
	if isNegative {
		res = "minus " + res
	}
	return res
}
