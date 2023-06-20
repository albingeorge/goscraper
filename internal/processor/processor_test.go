package processor

import (
	"encoding/json"
	"testing"

	"github.com/albingeorge/goscraper/internal/datasink"
	"github.com/albingeorge/goscraper/internal/reader"
	"go.uber.org/zap"
)

func BenchmarkProcess(b *testing.B) {
	// log.SetOutput(io.Discard)
	levelStr := `{
		"source": {
			"type": "default",
			"content": "https://mangapill.com/manga/2/one-piece"
		},
		"label": "chapter-list",
		"objects": {
			"chapters": {
				"parser": {
					"selector": "custom",
					"struct": "delayed",
					"value": "chapter_parser"
				},
				"sort": {
					"by": "name",
					"order": "asc"
				},
				"count": 2,
				"save": {
					"type": "directory",
					"path": {
						"type": "resolve",
						"content": "OnePiece/%current.name%"
					},
					"skipIfExists": false
				},
				"levels": [
					{
						"source": {
							"type": "resolve",
							"content": "https://mangapill.com%parent.url%"
						},
						"label": "page-list",
						"objects": {
							"pages": {
								"parser": {
									"selector": "custom",
									"struct": "delayed",
									"value": "page_parser"
								},
								"sort": {
									"by": "page_number",
									"order": "asc"
								},
								"count": 2,
								"save": {
									"type": "file",
									"name": {
										"type": "resolve",
										"content": "%current.name%.jpg"
									},
									"path": {
										"type": "resolve",
										"content": "OnePiece/%parent.name%/"
									},
									"content": {
										"type": "resolve",
										"content": "%current.src%"
									},
									"downloader": "delayed",
									"skipIfExists": true
								}
							}
						}
					}
				]
			}
		}
	}`
	levelVal := reader.Level{}
	json.Unmarshal([]byte(levelStr), &levelVal)
	levels := []reader.Level{
		levelVal,
	}
	dsLevelData := datasink.LevelData{}
	for i := 0; i < b.N; i++ {
		Process(levels, &dsLevelData, getSugaredLogger())
	}

}

func getSugaredLogger() *zap.SugaredLogger {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build()
	return logger.Sugar()
}
