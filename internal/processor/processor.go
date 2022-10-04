package processor

import (
	"log"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
)

func Process(objects []reader.Level, levelData *datasink.LevelData) {
	for _, level := range objects {
		// Fetch source content
		sourceUrl, err := reader.ResolveValue(level.Source, levelData)
		if err != nil {
			log.Println(err)
			continue
		}

		doc, err := httphandler.GetDocument(sourceUrl)
		if err != nil {
			log.Println(err)
			continue
		}

		result := map[string]datasink.Object{}

		// Parse source content
		for objName, obj := range level.Objects {
			if obj.Parser.Selector == reader.CUSTOM {
				result[objName], err = custom.Call(doc, obj)

				if err != nil {
					log.Println(err)
					continue
				}

				// for _, objeactData := range result[objName].Content {

				// }
			}
		}

		// Set child data for processing down the line
		levelData.Objects = result

		// Handle data storage
		// storage.Store(level.Save, result)

		// If child levels exist, call Process on the same
		// To be passed to child level processes
		// childData := datasink.LevelData{
		// 	ParentData: levelData,
		// }
	}
}
