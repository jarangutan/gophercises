package link

import (
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"testing"
)

func createChildNodes(content string, ctx *html.Node) []*html.Node {
	n, err := html.ParseFragment(strings.NewReader(content), ctx)
	if err != nil {
		panic(err)
	}

	return n
}

func createAnchorNode(href string, content string) *html.Node {
	attrs := make([]html.Attribute, 1)
	attrs[0] = html.Attribute{
		Namespace: "",
		Key:       "href",
		Val:       href,
	}

	anchor := &html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Type:        html.ElementNode,
		DataAtom:    atom.A,
		Data:        "a",
		Namespace:   "",
		Attr:        attrs,
	}

	childNodes := createChildNodes(content, anchor)

	for _, n := range childNodes {
		anchor.AppendChild(n)
	}

	return anchor
}

func Test_grabAnchor_Simple(t *testing.T) {
	expected := Link{"/other-page", "A link to a page!"}
	anchorNode := createAnchorNode(expected.href, expected.text)

	result := grabAnchor(anchorNode)
	if result != expected {
		t.Errorf("Result of \"%+v\" did not match expected \"%+v\"", result, expected)
	}
}

func Test_FindAnchors_Complex(t *testing.T) {
	expected := Link{"/other-page", "A link to a page!"}
	content := `<div>A link <span>to a <!--weird--><em>page</em></span></div>!`

	node := createAnchorNode(expected.href, content)

	result := FindAnchors(node, make([]Link, 0))
	if result[0] != expected {
		t.Errorf("Result of \"%+v\" did not match expected \"%+v\"", result, expected)
	}
}

func Test_FindAnchors_SampleFile(t *testing.T) {
	f, err := os.Open("ex4.html")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	result := FindAnchors(node, make([]Link, 0))
	expected := Link{"/dog-cat", "dog cat"}

	if result[0] != expected {
		t.Errorf("Result of \"%+v\" did not match expected \"%+v\"", result, expected)
	}
}
