package processor

import (
	"fmt"

	"github.com/albingeorge/goscraper/internal/reader"
)

func Process(objects []reader.Level) {
	for _, level := range objects {
		fmt.Printf("Level label: %v\n", level.Label)

		// Resolve source

		// Fetch source content

		// Parse source content

		// Resolve variables

		// Apply sort on variables

		// For each variable

		// Process data

		// If child exists, Process(child)
	}
}
