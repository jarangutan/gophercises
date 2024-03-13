package link

import (
	"strings"

	"golang.org/x/net/html"

	"testing"
)

func TestFindSimpleLink(t *testing.T) {
	input := `<html>
  <body>
    <h1>Hello!</h1>
    <a href="/other-page">A link to another page</a>
  </body>
</html>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}

	var links []Link
	FindA(doc, &links)

	println(links[0].Href)
	println(links[0].Text)

	if links[0].Href != "/other-page" {
		t.Error("Nope")
	}
}
