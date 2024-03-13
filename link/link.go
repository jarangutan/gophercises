package link

import (
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func FindA(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		t := strings.TrimSuffix(n.FirstChild.Data, "\n")

		var h string
		for _, a := range n.Attr {
			if a.Key == "href" {
				h = a.Val
			}
		}

		*links = append(*links, Link{
			Href: h,
			Text: t,
		})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FindA(c, links)
	}
}
