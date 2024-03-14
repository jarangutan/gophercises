package main

import (
	"bytes"
	"fmt"

	"github.com/jarangutan/gophercises/sitemap"
	"github.com/jarangutan/gophercises/sitemap/link"
)

func main() {
	page, err := sitemap.BuildSitemap("https://www.calhoun.io/")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(bytes.NewReader(page))

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
