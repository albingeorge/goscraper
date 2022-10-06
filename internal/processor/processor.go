package processor

import (
	"fmt"
	"log"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
	"github.com/albingeorge/goscraper/internal/storage"
)

func Process(configLevels []reader.Level, dsLevelData *datasink.LevelData) {
	for _, level := range configLevels {
		fmt.Println("Processing level: ", level.Label)

		// Fetch source content
		// We have not processed current object content here, hence passing nil
		sourceUrl, err := reader.ResolveValue(level.Source, dsLevelData)

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
			}

			for _, objeactData := range result[objName].Content {
				childLevelData := datasink.LevelData{
					ParentData:           dsLevelData,
					CurrentObjectContent: objeactData,
				}

				// Don't process child entries of the current object
				// if the data is already stored and SkipIfExists is set to true
				if storage.Store(obj.Save, &childLevelData) && obj.Save.SkipIfExists {
					continue
				}

				// Call child process
				Process(obj.Levels, &childLevelData)
			}
		}
	}
}
