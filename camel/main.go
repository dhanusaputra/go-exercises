package main

import "fmt"

import "strings"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Printf("Input is: %s\n", input)
	answer := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
		// fmt.Println(ch)
	}
	fmt.Println(answer)
}
