package datasink

type LevelData struct {
	ParentData *LevelData

	// Data fetched for each object
	// We do processing of child levels based on these objects
	// Key is the object name and value is the list of data fetched for each object
	// For example:
	// Key: "chapter"
	// Value: List of chapter data(which in and of itself will have map[string]interface{})
	Objects map[string][]Object
}

type Object map[string]interface{}
