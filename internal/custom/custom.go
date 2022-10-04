package custom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/extend"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Custom interface {
	Run(doc *goquery.Document, fnName string) error

	// Sorts the content data for the current object
	Sort(reader.Sort)

	// Fetches content data for each object.
	// Each datum can be a key value map with value of type interface{}
	GetContent() []map[string]interface{}
}

func Call(doc *goquery.Document, objectData reader.ObjectData) ([]map[string]interface{}, error) {
	obj, err := getCustomObject(objectData.Parser)

	if err != nil {
		return nil, err
	}

	err = obj.Run(doc, objectData.Parser.Value)

	obj.Sort(objectData.Sort)

	if err != nil {
		return nil, err
	}

	res := obj.GetContent()

	// Print result
	fmt.Println("Parsed result")
	marshalledText, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(marshalledText))

	return res, nil
}

func getCustomObject(variable reader.Parser) (Custom, error) {
	if variable.Struct == "mangapill" {
		return &extend.Mangapill{}, nil
	}

	return nil, errors.New("No custom implementation found for " + variable.Struct)
}
