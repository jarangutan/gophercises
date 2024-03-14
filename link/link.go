package link

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func grabText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += grabText(c)
	}

	r := strings.Join(strings.Fields(ret), " ")
	return r
}

func grabAnchor(n *html.Node) Link {
	var h string
	for _, a := range n.Attr {
		if a.Key == "href" {
			h = a.Val
		}
	}
	t := grabText(n)

	return Link{
		href: h,
		text: t,
	}
}

func FindAnchors(n *html.Node, links []Link) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		// t := strings.TrimSuffix(n.FirstChild.Data, "\n")
		links = append(links, grabAnchor(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = FindAnchors(c, links)
	}
	return links
}
