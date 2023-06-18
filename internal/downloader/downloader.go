package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Downloader interface {
	Download(url string, filename string) error
}

type DefaultDownloader struct{}

func (d DefaultDownloader) Download(url string, filename string) error {
	log.Printf("Downloading %v to %v", url, filename)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("received non-200 response code; url: %v", url)
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
