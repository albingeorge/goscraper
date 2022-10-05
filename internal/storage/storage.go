package storage

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
)

const BASE = "./output"

const PATH_SEPARATOR = "/"

func Store(save reader.Save, levelData *datasink.LevelData) {
	pathToSave, err := reader.ResolveValue(save.Name, *levelData.CurrentObjectContent, levelData)
	if err != nil {
		fmt.Println("Storage path resolve failure")
	}

	pathToSave = strings.Join([]string{BASE, pathToSave}, PATH_SEPARATOR)

	if save.Type == reader.STORAGE_DIRECTORY {
		err := os.MkdirAll(pathToSave, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
