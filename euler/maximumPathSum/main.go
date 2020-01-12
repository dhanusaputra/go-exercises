package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var triangle string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func main() {
	slice, err := parseStringToFloatSlice(triangle)
	if err != nil {
		panic(err)
	}
	res, err := sumMaximumPath(slice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}

func parseStringToFloatSlice(s string) ([][]float64, error) {
	var res [][]float64
	ss := strings.Split(s, "\n")
	for _, sss := range ss {
		var tmp []float64
		for _, ssss := range strings.Fields(sss) {
			fssss, err := strconv.ParseFloat(ssss, 64)
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, fssss)
		}
		res = append(res, tmp)
	}
	return res, nil
}

func sumMaximumPath(n [][]float64) (float64, error) {
	for i := len(n) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			n[i][j] += math.Max(n[i+1][j], n[i+1][j+1])
		}
	}
	return n[0][0], nil
}
