package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	n := flag.Float64("n", 2, "number")
	e := flag.Float64("e", 10, "exponent")
	flag.Parse()
	fmt.Println(getPowerDigitSum(*n, *e))
}

func getPowerDigitSum(n float64, e float64) int {
	var res int
	for _, i := range strings.Split(fmt.Sprintf("%.f", math.Pow(n, e)), "") {
		tmp, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		res += tmp
	}
	return res
}
