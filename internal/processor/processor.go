package processor

import (
	"log"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
	"github.com/albingeorge/goscraper/internal/storage"
)

func Process(objects []reader.Level, levelData *datasink.LevelData) {
	for _, level := range objects {
		// fmt.Printf("Level label: %v\n", level.Label)

		// todo: Resolve source if starts with %

		// Fetch source content
		doc, err := httphandler.GetDocument(level.Source)
		if err != nil {
			log.Println(err)
		}

		result := map[string][]datasink.Object{}

		// Parse source content
		for objName, obj := range level.Objects {
			if obj.Parser.Selector == reader.CUSTOM {
				result[objName], err = custom.Call(doc, obj)

				if err != nil {
					log.Println(err)
				}
			}
		}

		// Set child data for processing down the line
		levelData.Objects = result

		// Handle data storage
		storage.Store(level.Save, result)

		// If child levels exist, call Process on the same
		// To be passed to child level processes
		// childData := datasink.LevelData{
		// 	ParentData: levelData,
		// }
	}
}
