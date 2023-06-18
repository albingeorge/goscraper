package extend

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Mangapill struct {
	chapters datasink.Object
}

type MangapillDownloader struct{}

func (m *Mangapill) Run(doc *goquery.Document, fnName string) error {
	if fnName == "chapter_parser" {
		m.chapters = m.chapterParser(doc)
	} else if fnName == "page_parser" {
		m.chapters = m.pageParser(doc)
	} else {
		log.Println("Parser missing for input: ", fnName)
	}

	return nil
}

func (m *Mangapill) Sort(s reader.Sort) {
	if s.By == "name" {
		// Sort based on chapter number instead of alphabetically
		sort.Slice(m.chapters.Content, func(i, j int) bool {
			iContent := *(m.chapters.Content[i])
			jContent := *(m.chapters.Content[j])

			iName, _ := strconv.ParseFloat(strings.ReplaceAll(iContent[s.By].(string), "Chapter ", ""), 32)
			jName, _ := strconv.ParseFloat(strings.ReplaceAll(jContent[s.By].(string), "Chapter ", ""), 32)
			if s.Order == "desc" {
				return iName > jName
			}

			return iName < jName
		})
	} else if s.By == "page_number" {
		sort.Slice(m.chapters.Content, func(i, j int) bool {
			iVal := *(m.chapters.Content[i])
			jVal := *(m.chapters.Content[j])
			return iVal["name"].(int) < jVal["name"].(int)
		})
	}
}

func (m *Mangapill) GetContent() datasink.Object {
	return m.chapters
}

func (m *Mangapill) chapterParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}

	doc.Find("#chapters a").Each(func(i int, s *goquery.Selection) {
		attrVal, _ := s.Attr("href")
		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": s.Text(),
			"url":  attrVal,
		})
	})

	return res
}

func (m *Mangapill) pageParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}

	doc.Find("chapter-page img").Each(func(i int, s *goquery.Selection) {
		imgSrc, _ := s.Attr("data-src")
		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": i + 1,
			"src":  imgSrc,
		})
	})

	return res
}

func (m MangapillDownloader) Download(url string, filename string) error {
	log.Printf("Downloading %v to %v", url, filename)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("referer", "https://mangapill.com/")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("received non-200 response code for url %v", url)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
