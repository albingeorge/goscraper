package processor

import (
	"fmt"
	"log"

	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/httphandler"
	"github.com/albingeorge/goscraper/internal/reader"
)

func Process(objects []reader.Level, parentData datasink.LevelData) {
	for _, level := range objects {
		fmt.Printf("Level label: %v\n", level.Label)

		// Resolve source

		// Fetch source content
		doc, err := httphandler.GetDocument(level.Source)
		if err != nil {
			log.Println(err)
		}

		// Parse source content

		// result := []datasink.Object{}
		result := map[string]datasink.Object{}

		for objName, obj := range level.Objects {
			o := datasink.Object{
				Variables: map[string]interface{}{},
			}

			if obj.Data.Selector == reader.CUSTOM {
				o.Variables, err = custom.Call(doc, obj.Data, obj.Sort)

				if err != nil {
					log.Println(err)
				}

			}
			result[objName] = o
		}

		// For each object
		// Pre-process object
		// Parse variables

		// Apply sort on objects if required

		// For each object
		// Process object(i.e, apply save)

		// If child exists, Process(child)
	}
}
