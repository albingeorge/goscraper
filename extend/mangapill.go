package extend

import (
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Chapter struct {
	Name string
	Url  string
}

type Mangapill struct {
	chapters []Chapter
}

func (m *Mangapill) Run(doc *goquery.Document, fnName string) error {
	if fnName == "chapter_parser" {
		m.chapters = chapterParser(doc)
	}
	return nil
}

func (m *Mangapill) Sort(s reader.Sort) {
	// sort.Sl
	// sort.Slice()
	sort.Slice(m.chapters, func(i, j int) bool {
		iName, _ := strconv.ParseFloat(strings.ReplaceAll(m.chapters[i].Name, "Chapter ", ""), 1)
		jName, _ := strconv.ParseFloat(strings.ReplaceAll(m.chapters[j].Name, "Chapter ", ""), 1)
		return iName < jName
	})
}

func (m *Mangapill) GetContent() interface{} {
	return m.chapters
}

func chapterParser(doc *goquery.Document) []Chapter {
	res := []Chapter{}
	doc.Find("#chapters a").Each(func(i int, s *goquery.Selection) {
		attrVal, _ := s.Attr("href")
		res = append(res, Chapter{
			Name: s.Text(),
			Url:  attrVal,
		})
	})

	return res
}
