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
	Source     string              `json:"source"`
	Label      string              `json:"label"`
	NameFormat string              `json:"name_format"`
	SaveAs     string              `json:"save_as"`
	Variables  map[string]Variable `json:"variables"`
	Sort       Sort                `json:"sort"`
	Levels     []Level             `json:"levels"`
}

type Variable struct {
	Selector string
	Value    string
}

type Sort struct {
	By    string
	Order string
}

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

	return *objects
}
