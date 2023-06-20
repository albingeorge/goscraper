package storage

import (
	"fmt"
	"os"
	"strings"

	"github.com/albingeorge/goscraper/extend"
	"github.com/albingeorge/goscraper/internal/custom"
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/downloader"
	"github.com/albingeorge/goscraper/internal/reader"
	"go.uber.org/zap"
)

const BASE = "./output"

const PATH_SEPARATOR = "/"

// Stores the data fetched for each object content
// Returns bool representing whether the data is already stored
func Store(save reader.Save, levelData *datasink.LevelData, log *zap.SugaredLogger) bool {
	log.With(
		"type", save.Type,
		"path", save.Path,
	).Debugf("Storing data")

	pathToSave, err := reader.ResolveValue(save.Path, levelData, log)
	if err != nil {
		log.Errorf("Storage path resolve failure: %w", err)
	}

	pathToSave = strings.Join([]string{BASE, pathToSave}, PATH_SEPARATOR)

	if save.Type == reader.STORAGE_DIRECTORY {
		if _, err := os.Stat(pathToSave); !os.IsNotExist(err) {
			return true
		}
	}

	err = os.MkdirAll(pathToSave, os.ModePerm)

	if err != nil {
		log.Errorf("error creating directory: %v", pathToSave)
	}

	if save.Type == reader.STORAGE_FILE {
		fileName, err := reader.ResolveValue(save.Name, levelData, log)
		if err != nil {
			log.Error("File name resolve failure")
		}

		pathToSave += fileName

		url, err := reader.ResolveValue(save.Content, levelData, log)
		if err != nil {
			log.Error("Download url resolve failure")
		}

		if _, err := os.Stat(pathToSave); !os.IsNotExist(err) {
			return true
		}

		downloader, err := getDownloader(save)
		if err != nil {
			log.Errorf("downloader not found: %w", err)
		}

		err = downloader.Download(url, pathToSave)
		if err != nil {
			log.Errorf("download failure: %w", err)
		}
	}

	return false
}

func getDownloader(save reader.Save) (downloader.Downloader, error) {
	switch save.Downloader {
	case "mangapill":
		return extend.MangapillDownloader{}, nil
	case "delayed":
		return custom.DelayedDownloaderImplementation{}, nil
	case "":
		return downloader.DefaultDownloader{}, nil
	}
	return nil, fmt.Errorf("invalid downloader: %v", save.Downloader)
}
