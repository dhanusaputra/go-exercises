package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	s := flag.String("start", "1901-01-01", "start of calculation")
	e := flag.String("end", "2000-12-31", "end of calculation")
	flag.Parse()
	start, err := time.Parse(time.RFC3339[0:10], *s)
	if err != nil {
		panic(err)
	}
	end, err := time.Parse(time.RFC3339[0:10], *e)
	if err != nil {
		panic(err)
	}
	res, err := countSundayInFirstMonth(start, end)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}

func countSundayInFirstMonth(start time.Time, end time.Time) (int, error) {
	var res int
	for i := start; i.Before(end); i = i.AddDate(0, 0, 1) {
		if i.Weekday() == time.Sunday && i.Day() == 1 {
			res++
		}
	}
	return res, nil
}
