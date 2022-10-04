package reader

import (
	"encoding/json"
	"io"
	"os"
)

type Levels struct {
	Levels []Level `json:"levels"`
}

type Level struct {
	Source  string                `json:"source"`
	Label   string                `json:"label"`
	Save    Save                  `json:"save"`
	Objects map[string]ObjectData `json:"objects"`
	Sort    Sort                  `json:"sort"`
	Levels  []Level               `json:"levels"`
}

type Save struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	Content string `json:"content"`
}

type ObjectData struct {
	Parser Parser `json:"parser"`
	Sort   Sort   `json:"sort"`
}

type Parser struct {
	Selector string
	Struct   string
	Value    string
}

type Sort struct {
	By    string
	Order string
}

const CUSTOM = "custom"

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

	// fmt.Printf("%+v\n", objects)

	return *objects
}

// func Resolve(name string, data datasink.LevelData) string
