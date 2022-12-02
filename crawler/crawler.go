package crawler

import (
	"log"
	"net/http"
)

func GetUrl(url string) *http.Response {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	return res
}
