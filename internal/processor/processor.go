package processor

import (
	"log"
	"strconv"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
	"github.com/albingeorge/goscraper/internal/storage"
)

func Process(configLevels []reader.Level, dsLevelData *datasink.LevelData) {
	for _, level := range configLevels {
		log.Println("Processing level: ", level.Label)

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

		// Parse source content
		for objName, obj := range level.Objects {
			log.Println("Processing object: ", objName)

			// Contains the data fetched from each object by parsing the document
			var objectData datasink.Object

			if obj.Parser.Selector == reader.CUSTOM {
				objectData, err = custom.Call(doc, obj)

				if err != nil {
					log.Println(err)
					continue
				}
			}

			objContentChan := make(chan string, 2)

			for i, objectDataContent := range objectData.Content {
				contentName := level.Label + "-" + strconv.Itoa(i+1)
				log.Printf("Goroutine starting: %v\n", contentName)
				go func(dsLevelData *datasink.LevelData, obj reader.ObjectData, objectDataContent *datasink.ObjectContent, contentName string) {
					childLevelData := datasink.LevelData{
						ParentData:           dsLevelData,
						CurrentObjectContent: objectDataContent,
					}

					// Don't process child entries of the current object
					// if the data is already stored and SkipIfExists is set to true
					// todo: refactor this, so that it checks if the file/directory exists before attempting to store it first
					if !(storage.Store(obj.Save, &childLevelData) && obj.Save.SkipIfExists) {
						// Call child process
						Process(obj.Levels, &childLevelData)
					}

					log.Printf("Goroutine completed: %v\n", contentName)

					objContentChan <- "Done processing content"
				}(dsLevelData, obj, objectDataContent, contentName)

			}
			log.Println("Waiting for contents to complete in object: ", objName)
			for i := 0; i < len(objectData.Content); i++ {
				<-objContentChan
			}
		}
	}
}
