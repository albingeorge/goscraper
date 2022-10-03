package custom

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/albingeorge/goscraper/extend"
	"github.com/albingeorge/goscraper/internal/reader"
)

type Custom interface {
	Run(doc *goquery.Document, fnName string) error
	Sort(reader.Sort)
	GetContent() interface{}
}

func Call(doc *goquery.Document, v reader.Variable, sort reader.Sort) (interface{}, error) {
	obj, err := getCustomObject(v)

	if err != nil {
		return nil, err
	}

	err = obj.Run(doc, v.Value)

	obj.Sort(sort)

	if err != nil {
		return nil, err
	}

	res := obj.GetContent()

	fmt.Println("Result parsed")
	fmt.Println(res)

	return res, nil
}

func getCustomObject(variable reader.Variable) (Custom, error) {
	if variable.Struct == "mangapill" {
		return &extend.Mangapill{}, nil
	}

	return nil, errors.New("No custom implementation found for " + variable.Struct)
}
