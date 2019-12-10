package main

import (
	"flag"
	"fmt"
)

func main() {
	maxValue := flag.Int("max", 4000000, "maximum even fibonacci value")
	flag.Parse()
	var result, currentValue, previousValue, previous2Value int
	previousValue = 1
	for result < *maxValue-currentValue {
		currentValue = previousValue + previous2Value
		result += currentValue
		fmt.Printf("%d Total: %d\n", currentValue, result)
		previous2Value = previousValue
		previousValue = currentValue
	}
}
