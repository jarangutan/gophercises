package link

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func grabText(n *html.Node, s string) string {
	if n.Type == html.TextNode {
		s = s + " " + n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s = grabText(c, s)
	}
	return strings.TrimSpace(s)
}

func grabAnchor(n *html.Node) Link {
	var h string
	for _, a := range n.Attr {
		if a.Key == "href" {
			h = a.Val
		}
	}
	t := grabText(n, "")

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
