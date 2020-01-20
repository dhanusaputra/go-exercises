// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

// 012   021   102   120   201   210

// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package main

import (
	"fmt"
)

func main() {
	permutations := getPermutations("0123456789")
	fmt.Printf("%v\n", len(filterLexicographicPermutations(permutations)))
}

func join(r []rune, c rune) []string {
	var res []string
	for i := 0; i <= len(r); i++ {
		res = append(res, string(r[:i])+string(c)+string(r[i:]))
	}
	return res
}

func getPermutations(s string) []string {
	var n func(s []rune, p []string) []string
	n = func(s []rune, p []string) []string {
		if len(s) == 0 {
			return p
		}
		res := []string{}
		for _, e := range p {
			res = append(res, join([]rune(e), s[0])...)
		}
		return n(s[1:], res)
	}
	res := []rune(s)
	return n(res[1:], []string{string(res[0])})
}

func filterLexicographicPermutations(s []string) []string {
	var res []string
	for _, i := range s {
		isLex := true
		tmp := []rune(i)
		for j := 0; j < 7; j++ {
			if j != 0 && tmp[j] < tmp[j-1] {
				isLex = false
				break
			}
		}
		if isLex {
			res = append(res, i[0:7])
		}
	}
	return res
}
