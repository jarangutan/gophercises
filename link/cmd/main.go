package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jarangutan/gophercises/link"
	"golang.org/x/net/html"
)

func main() {
	filename := flag.String("file", "ex.html", "html file to parse for a refs")
	flag.Parse()

	fmt.Printf("Looking for a refs in file %s\n", *filename)
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	links := link.FindAnchors(doc, make([]link.Link, 0))
	fmt.Printf("links %+v\n", links)
}
