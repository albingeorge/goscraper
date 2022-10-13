package httphandler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func GetDocument(url string) (*goquery.Document, error) {
	log.Println("url: ", url)
	res, err := http.Get(url)

	if err != nil {
		log.Printf("Empty response")
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("invalid http status: " + strconv.Itoa(res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
