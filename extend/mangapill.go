package extend

import (
	"fmt"
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

func (m *Mangapill) Run(doc *goquery.Document, fnName string) error {
	if fnName == "chapter_parser" {
		m.chapters = m.chapterParser(doc)
	} else if fnName == "page_parser" {
		m.chapters = m.pageParser(doc)
	} else {
		fmt.Println("Parser missing for input: ", fnName)
	}

	return nil
}

func (m *Mangapill) Sort(s reader.Sort) {
	if s.By == "name" {
		// Sort based on chapter number instead of alphabetically
		sort.Slice(m.chapters.Content, func(i, j int) bool {
			iContent := *(m.chapters.Content[i])
			jContent := *(m.chapters.Content[j])
			// &(m.chapters.Content[i]).
			iName, _ := strconv.ParseFloat(strings.ReplaceAll(iContent[s.By].(string), "Chapter ", ""), 32)
			jName, _ := strconv.ParseFloat(strings.ReplaceAll(jContent[s.By].(string), "Chapter ", ""), 32)
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

	// Temp code to fetch only last 2 chapters
	doc.Find("#chapters a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attrVal, _ := s.Attr("href")

		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": s.Text(),
			"url":  attrVal,
		})

		return i != 1
	})

	// doc.Find("#chapters a").Each(func(i int, s *goquery.Selection) {
	// 	attrVal, _ := s.Attr("href")
	// 	res.Content = append(res.Content, &datasink.ObjectContent{
	// 		"name": s.Text(),
	// 		"url":  attrVal,
	// 	})
	// })

	return res
}

func (m *Mangapill) pageParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}

	doc.Find("chapter-page img").EachWithBreak(func(i int, s *goquery.Selection) bool {

		imgSrc, _ := s.Attr("data-src")

		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": i + 1,
			"src":  imgSrc,
		})

		// Remove later
		return i != 1
	})

	return res
}
