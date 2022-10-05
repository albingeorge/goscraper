package datasink

import (
	"errors"
	"strings"
)

// Determines how we store data internally
type LevelData struct {
	ParentData *LevelData

	CurrentObjectContent *ObjectContent

	// Data fetched for each object
	// We do processing of child levels based on these objects
	// Key is the object name and value is the list of data fetched for each object
	// For example:
	// Key: "chapter"
	// Value: List of chapter data(which in and of itself will have map[string]interface{})
	Objects map[string]Object
}

type Object struct {
	Content []*ObjectContent
}

type ObjectContent map[string]interface{}

// Find the value for a dotted notation input string from the datasink object
func FindValue(input string, levelData *LevelData) (interface{}, error) {
	sections := strings.Split(input, ".")

	if len(sections) != 2 {
		return "", errors.New("invalid format for input: " + input)
	}

	obj := *levelData.CurrentObjectContent

	// Resolve which level the data is to be fetched from
	if sections[0] == "parent" {
		if levelData.ParentData == nil {
			return "", errors.New("no parent content available for this level")
		}

		if levelData.ParentData.CurrentObjectContent != nil {
			obj = *levelData.ParentData.CurrentObjectContent
		}
	}

	if obj == nil {
		return "", errors.New("unable to find value ObjectContent in " + sections[0])
	}

	if val, ok := obj[sections[1]]; ok {
		return val, nil
	}

	return "", errors.New("unable to find value in ObjectContent in " + sections[0])
}
