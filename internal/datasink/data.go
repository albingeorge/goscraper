package datasink

type LevelData struct {
	ParentData *LevelData
	Objects    map[string]Object
}

type Object struct {
	Variables interface{}
}
