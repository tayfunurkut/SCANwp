package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func UrlFormat(url string) string {

	if !strings.Contains(url, "http") {
		cl := http.Client{}
		_, err := cl.Head(fmt.Sprintf("https://%s", url))
		if err != nil {
			return fmt.Sprintf("http://%s", url)
		}
		return fmt.Sprintf("https://%s", url)

	}
	return url
}

func SendHTTPGetRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
