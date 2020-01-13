package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dhanusaputra/go-exercises/link"
)

func main() {
	URLFlag := flag.String("url", "https://gophercises.com", "the url that you  want to build")
	flag.Parse()

	fmt.Println(*URLFlag)
	resp, err := http.Get(*URLFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}
}

/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links w/ a diff domain
	5. find all pages (BFS)
	6. print out XML
*/
