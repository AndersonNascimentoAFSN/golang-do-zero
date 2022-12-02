package crawler

import (
	"net/http"
)

func GetUrl(url string) (*http.Response, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return res, nil
}
