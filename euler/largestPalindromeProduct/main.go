package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPalindrome(i int) bool {
	if strconv.Itoa(i) == reverseInt(i) {
		return true
	}
	return false
}

func reverseInt(i int) string {
	s := strconv.Itoa(i)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrome(i * j) {
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
				os.Exit(0)
			}
		}
	}
}
