// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

// What is the total of all the name scores in the file?

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	lines, err := parseTxt("euler/namesScore/p022_names.txt")
	if err != nil {
		panic(err)
	}
	splitedLines := strings.Split(strings.Trim(strings.Join(lines, "\",\""), "\""), "\",\"")
	sort.Strings(splitedLines)
	var res int
	for i, v := range splitedLines {
		res += i * getStringWorth(v)
	}
	fmt.Println(res)
}

func parseTxt(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func getStringWorth(s string) int {
	var res int
	for _, v := range s {
		res += int(v) - 64
	}
	return res
}
