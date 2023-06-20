package httphandler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

func GetDocument(log *zap.SugaredLogger, url string) (*goquery.Document, error) {
	log.Debugf("url: %v", url)

	res, err := http.Get(url)

	if err != nil {
		log.Debugf("empty response")
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("invalid http status on url: %v", url)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
