package sitemap

import (
	"io"
	"net/http"
)

func BuildSitemap(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	page := string(body[:])

	return page, nil
}
