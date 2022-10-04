package datasink

import (
	"errors"
	"strings"
)

// Determines how we store data internally
type LevelData struct {
	ParentData *LevelData

	// Data fetched for each object
	// We do processing of child levels based on these objects
	// Key is the object name and value is the list of data fetched for each object
	// For example:
	// Key: "chapter"
	// Value: List of chapter data(which in and of itself will have map[string]interface{})
	Objects map[string]Object
}

type Object struct {
	ParentObject *Object
	Content      []*ObjectContent
}

type ObjectContent map[string]interface{}

// Find the value for a dotted notation input string from the datasink object
func FindValue(input string, obj Object) (string, error) {
	sections := strings.Split(input, ".")

	// Resolve which level the data is to be fetched from
	if sections[0] == "parent" {
		obj = *obj.ParentObject
	}

	// objData := []Object{}
	// // Resolve the object
	// if objData, ok := sink.Objects[sections[1]]; ok {

	// }

	return "", errors.New("unable to find value in leveldata")
}
