package main

import (
	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/processor"
	"github.com/albingeorge/goscraper/internal/reader"
	"github.com/albingeorge/goscraper/pkg/log"
)

func main() {
	log := log.InitializeLogger()
	objects := reader.Read(log)

	processor.Process(objects.Levels, &datasink.LevelData{}, log)
}
