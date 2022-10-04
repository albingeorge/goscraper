package main

import (
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/processor"
	"github.com/albingeorge/goscraper/internal/reader"
)

func main() {
	objects := reader.Read()
	processor.Process(objects.Levels, &datasink.LevelData{})
}
