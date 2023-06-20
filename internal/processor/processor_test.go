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
		"label": "chapter",
		"objects": {
			"chapter": {
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
						"label": "page",
						"objects": {
							"page": {
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

	log, _ := zap.NewDevelopment()
	dsLevelData := datasink.LevelData{}
	for i := 0; i < b.N; i++ {
		Process(levels, &dsLevelData, log.Sugar())
	}

}
