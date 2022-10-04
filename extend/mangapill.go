package extend

import (
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
	}
	return nil
}

func (m *Mangapill) Sort(s reader.Sort) {
	// Sort based on chapter number instead of alphabetically
	sort.Slice(m.chapters.Content, func(i, j int) bool {
		iContent := *(m.chapters.Content[i])
		jContent := *(m.chapters.Content[j])
		// &(m.chapters.Content[i]).
		iName, _ := strconv.ParseFloat(strings.ReplaceAll(iContent[s.By].(string), "Chapter ", ""), 32)
		jName, _ := strconv.ParseFloat(strings.ReplaceAll(jContent[s.By].(string), "Chapter ", ""), 32)
		return iName < jName
	})
}

func (m *Mangapill) GetContent() datasink.Object {
	return m.chapters
}

func (m *Mangapill) chapterParser(doc *goquery.Document) datasink.Object {
	res := datasink.Object{
		Content: []*datasink.ObjectContent{},
	}

	count := 1

	// Temp code to fetch only last 10 chapters
	doc.Find("#chapters a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attrVal, _ := s.Attr("href")

		res.Content = append(res.Content, &datasink.ObjectContent{
			"name": s.Text(),
			"url":  attrVal,
		})
		count++

		return count <= 4
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
