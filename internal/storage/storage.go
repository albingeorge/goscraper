package storage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
)

const BASE = "./output"

const PATH_SEPARATOR = "/"

func Store(save reader.Save, levelData *datasink.LevelData) {
	pathToSave, err := reader.ResolveValue(save.Path, levelData)
	if err != nil {
		fmt.Println("Storage path resolve failure")
	}

	pathToSave = strings.Join([]string{BASE, pathToSave}, PATH_SEPARATOR)

	err = os.MkdirAll(pathToSave, os.ModePerm)

	if err != nil {
		log.Println("error creating directory: ", pathToSave)
	}

	if save.Type == reader.STORAGE_FILE {
		fileName, err := reader.ResolveValue(save.Name, levelData)
		if err != nil {
			fmt.Println("File name resolve failure")
		}

		pathToSave += fileName

		url, err := reader.ResolveValue(save.Content, levelData)
		if err != nil {
			fmt.Println("Download url resolve failure")
		}

		err = download(url, pathToSave)
		if err != nil {
			fmt.Println("download failure: ", err)
		}
	}
}

func download(url string, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non-200 response code")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
