package main

import (
	"fmt"
	"strings"

	"github.com/jarangutan/gophercises/sitemap"
	"github.com/jarangutan/gophercises/sitemap/link"
)

func main() {
	page, err := sitemap.BuildSitemap("https://www.calhoun.io/")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(strings.NewReader(page))

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
