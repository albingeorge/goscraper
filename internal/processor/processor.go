package processor

import (
	"strconv"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
	"github.com/albingeorge/goscraper/internal/storage"
	"go.uber.org/zap"
)

func Process(configLevels []reader.Level, dsLevelData *datasink.LevelData, log *zap.SugaredLogger) {
	for _, level := range configLevels {
		log.Debugf("Processing level: %v", level.Label)

		// Fetch source content
		// We have not processed current object content here, hence passing nil
		sourceUrl, err := reader.ResolveValue(level.Source, dsLevelData, log)

		if err != nil {
			log.Errorf("Error resolving value: %w", err)
			continue
		}

		doc, err := httphandler.GetDocument(log, sourceUrl)
		if err != nil {
			log.Errorf("Error fetching document: %w", err)
			continue
		}

		// Parse source content
		for objName, obj := range level.Objects {
			log.Debugf("Processing object: %v", objName)

			// Contains the data fetched from each object by parsing the document
			var objectData datasink.Object

			if obj.Parser.Selector == reader.CUSTOM {
				objectData, err = custom.Call(doc, obj)

				if err != nil {
					log.Errorf("Error in custom call: %w", err)
					continue
				}
			}

			objContentChan := make(chan string, 2)

			for i, objectDataContent := range objectData.Content {
				contentName := level.Label + "-" + strconv.Itoa(i+1)
				log.Debugf("Goroutine starting: %v", contentName)
				go func(dsLevelData *datasink.LevelData, obj reader.ObjectData, objectDataContent *datasink.ObjectContent, contentName string) {
					childLevelData := datasink.LevelData{
						ParentData:           dsLevelData,
						CurrentObjectContent: objectDataContent,
					}

					// Don't process child entries of the current object
					// if the data is already stored and SkipIfExists is set to true
					// todo: refactor this, so that it checks if the file/directory exists before attempting to store it first
					if !(storage.Store(obj.Save, &childLevelData, log) && obj.Save.SkipIfExists) {
						// Call child process
						Process(obj.Levels, &childLevelData, log)
					}

					log.Debugf("Goroutine completed: %v", contentName)

					objContentChan <- "Done processing content"
				}(dsLevelData, obj, objectDataContent, contentName)

			}

			log.Debugf("Waiting for contents to complete in object: %v", objName)
			for i := 0; i < len(objectData.Content); i++ {
				<-objContentChan
			}
		}
	}
}
