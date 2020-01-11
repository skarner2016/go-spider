package common

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var defaultHeader = map[string]string{
	"Host":                      "movie.douban.com",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	"Referer":                   "https://movie.douban.com/top250",
}

func SendRequest(header map[string]string, url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(header) == 0 {
		header = defaultHeader
	}

	// header
	for key, value := range header {
		request.Header.Add(key, value)
	}
	client := &http.Client{}

	return client.Do(request)
}

func GetDoc(header map[string]string, url string) (*goquery.Document, error) {
	response, err := SendRequest(header, url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("status code: " + strconv.Itoa(response.StatusCode))
	}

	return goquery.NewDocumentFromReader(response.Body)
}
