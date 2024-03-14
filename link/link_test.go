package link

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"

	"testing"
)

func createNode(content string) *html.Node {
	n, err := html.Parse(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	return n
}

func createAnchorNode(href string, text string) *html.Node {
	textNode := createNode(text)
	attrs := make([]html.Attribute, 1)
	attrs[0] = html.Attribute{
		Namespace: "",
		Key:       "href",
		Val:       href,
	}

	return &html.Node{
		Parent:      nil,
		FirstChild:  textNode,
		LastChild:   textNode,
		PrevSibling: nil,
		NextSibling: nil,
		Type:        html.ElementNode,
		DataAtom:    0,
		Data:        "a",
		Namespace:   "",
		Attr:        attrs,
	}
}

func Test_grabText_JustText(t *testing.T) {
	expected := "some text"
	node := createNode(expected)

	result := grabText(node, "")
	if result != expected {
		t.Errorf("Result of \"%v\" did not match expected \"%v\"", result, expected)
	}
}

func Test_grabText_TextInsideSpan(t *testing.T) {
	expected := "some text inside a span"
	node := createNode(fmt.Sprintf("<span>%v</span>", expected))

	result := grabText(node, "")
	if result != expected {
		t.Errorf("Result of \"%v\" did not match expected \"%v\"", result, expected)
	}
}

func Test_grabText_TextInsideSpanWithComments(t *testing.T) {
	expected := "some text inside a span"
	node := createNode(fmt.Sprintf("<span><!-- Hello I am a comment :) -->%v</span>", expected))

	result := grabText(node, "")
	if result != expected {
		t.Errorf("Result of \"%v\" did not match expected \"%v\"", result, expected)
	}
}

func Test_grabText_TextInAnchorSplitBySpan(t *testing.T) {
	expected := "some text split by things"
	node := createNode(fmt.Sprintf(`<div>
      <a href=\"https://www.youtube.com/watch?v=dQw4w9WgXcQ\">some text split
        <span>by <!--multiple--></span>things
      </a>
    <div>
    `))

	result := grabText(node, "")
	if result != expected {
		t.Errorf("Result of \"%v\" did not match expected \"%v\"", result, expected)
	}
}

func Test_grabAnchor_JustAnchor(t *testing.T) {
	expected := Link{"/other-page", "A link to a page!"}
	anchorNode := createAnchorNode(expected.href, expected.text)

	result := grabAnchor(anchorNode)
	if result != expected {
		t.Errorf("Result of \"%+v\" did not match expected \"%+v\"", result, expected)
	}
}

func TestFindSimpleLink(t *testing.T) {
	expected := Link{"/other-page", "A link to a page!"}

	node := createNode(`
    <html>
      <body>
        <a href="/other-page">A link to a page!</a>
      </body>
    </html>
    `)

	result := FindAnchors(node, make([]Link, 0))
	if result[0] != expected {
		t.Errorf("Result of \"%+v\" did not match expected \"%+v\"", result, expected)
	}
}