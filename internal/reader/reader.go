package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/albingeorge/goscraper/internal/datasink"
)

type Levels struct {
	Levels []Level `json:"levels"`
}

// Level represents a URL
// It contains multiple objects
type Level struct {
	Source  Resolve               `json:"source"`
	Label   string                `json:"label"`
	Objects map[string]ObjectData `json:"objects"`
}

// Object represents various sections which can be extracted from a page(Level)
// For example, from a manga page(say, https://mangapill.com/manga/2/one-piece)
// we can get multiple objects:
// 1. List of chapters
// 2. Manga details(like author name, manga cover, etc)
// Hence, we create a map of ObjectData under each Level
type ObjectData struct {
	Parser Parser  `json:"parser"`
	Sort   Sort    `json:"sort"`
	Save   Save    `json:"save"`
	Levels []Level `json:"levels"`
}

// Parser to fetch object data from the source
type Parser struct {
	Selector string
	Struct   string
	Value    string
}

// Sort objects if required
type Sort struct {
	By    string
	Order string
}

// Determine how to save an object content if required.
// For example, for each chapter, we might need to create a directory
// and for each page of a chapter, we will have to download the file content
// from the image URL.
type Save struct {
	Type    string  `json:"type"`
	Name    Resolve `json:"name"`
	Path    string  `json:"path"`
	Content string  `json:"content"`
}

type Resolve struct {
	Type    string
	Content string
}

const CUSTOM = "custom"

const RESOLVE_TYPE = "resolve"

// Handles read of in the input config file
// Reads from input/input.json
func Read() Levels {
	file, err := os.Open("input/input.json")
	if err != nil {
		panic("error reading input file")
	}

	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic("error reading input file")
	}

	objects := new(Levels)
	json.Unmarshal(fileContents, objects)

	// fmt.Printf("%#v\n", objects)

	return *objects
}

// Resolve a value
func ResolveValue(resolve Resolve, data *datasink.LevelData) (string, error) {
	result := resolve.Content

	if resolve.Type == RESOLVE_TYPE {
		r := regexp.MustCompile(`%([^%]*)%`)
		result = r.ReplaceAllStringFunc(resolve.Content, func(find string) string {
			input := find[1 : len(find)-1]
			fmt.Println(input)
			return input
		})
	}

	return result, nil
}
