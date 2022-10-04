package extend

import (
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Mangapill struct {
	chapters []map[string]interface{}
}

func (m *Mangapill) Run(doc *goquery.Document, fnName string) error {
	if fnName == "chapter_parser" {
		m.chapters = m.chapterParser(doc)
	}
	return nil
}

func (m *Mangapill) Sort(s reader.Sort) {
	// Sort based on chapter number instead of alphabetically
	sort.Slice(m.chapters, func(i, j int) bool {
		iName, _ := strconv.ParseFloat(strings.ReplaceAll(m.chapters[i][s.By].(string), "Chapter ", ""), 32)
		jName, _ := strconv.ParseFloat(strings.ReplaceAll(m.chapters[i][s.By].(string), "Chapter ", ""), 32)
		return iName < jName
	})
}

func (m *Mangapill) GetContent() []map[string]interface{} {
	return m.chapters
}

func (m *Mangapill) chapterParser(doc *goquery.Document) []map[string]interface{} {
	res := []map[string]interface{}{}

	count := 1

	// Temp code to fetch only last 10 chapters
	doc.Find("#chapters a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attrVal, _ := s.Attr("href")

		res = append(res, map[string]interface{}{
			"name": s.Text(),
			"url":  attrVal,
		})
		count++

		return count <= 10
	})

	// doc.Find("#chapters a").EachWithBreak(func(i int, s *goquery.Selection) {
	// 	attrVal, _ := s.Attr("href")
	// 	res = append(res, Chapter{
	// 		Name: s.Text(),
	// 		Url:  attrVal,
	// 	})
	// })

	return res
}
