package custom

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
)

// DelayedImplementation is implementation of Custom parser
// which simulates a real-world http request by introducing delays
type DelayedImplementation struct {
	content datasink.Object
}

type DelayedDownloaderImplementation struct{}

func (m *DelayedImplementation) Run(doc *goquery.Document, fnName string) error {
	if fnName == "chapter_parser" {
		m.content = chapterParser(doc)
		time.Sleep(50 * time.Millisecond)
	} else if fnName == "page_parser" {
		m.content = pageParser(doc)
		time.Sleep(100 * time.Millisecond)
	} else {
		log.Println("Parser missing for input: ", fnName)
	}

	return nil
}

func (m *DelayedImplementation) Sort(s reader.Sort) {}

func (m *DelayedImplementation) GetContent() datasink.Object {
	return m.content
}

func chapterParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}
	for i := 0; i < 2000; i++ {
		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": fmt.Sprintf("Chapter %v", i+1),
			"url":  "",
		})
	}

	return res
}

func pageParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}
	for i := 0; i < 50; i++ {
		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": fmt.Sprintf("Page %v", i+1),
			"url":  "",
		})
	}

	return res
}

func (m DelayedDownloaderImplementation) Download(url string, filename string) error {
	time.Sleep(100 * time.Millisecond)
	return nil
}
