package custom

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/extend"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Custom interface {
	Run(doc *goquery.Document, fnName string) error

	// Sorts the content data for the current object
	Sort(reader.Sort)

	// Fetches content data for each object.
	// Each datum can be a key value map with value of type interface{}
	GetContent() datasink.Object
}

func Call(doc *goquery.Document, objectData reader.ObjectData) (datasink.Object, error) {
	obj, err := getCustomObject(objectData.Parser)

	if err != nil {
		return datasink.Object{}, err
	}

	err = obj.Run(doc, objectData.Parser.Value)

	if err != nil {
		return datasink.Object{}, err
	}

	obj.Sort(objectData.Sort)

	res := obj.GetContent()

	count := len(res.Content)
	if objectData.Count > 0 {
		count = objectData.Count
	}

	res.Content = res.Content[:count]

	return res, nil
}

func getCustomObject(variable reader.Parser) (Custom, error) {
	if variable.Struct == "mangapill" {
		return &extend.Mangapill{}, nil
	}

	if variable.Struct == "delayed" {
		return &DelayedImplementation{}, nil
	}

	return nil, errors.New("No custom implementation found for " + variable.Struct)
}
