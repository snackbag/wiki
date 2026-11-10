package render

import (
	"fmt"
	"os"
	"path/filepath"
)

var cacheNotFound string

func NotFound() string {
	if cacheNotFound == "" {
		path := filepath.Join("templates", "404.html")
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Sprintf("failed to get 404 page: %s", err)
		}

		cacheNotFound = string(data)
	}

	return cacheNotFound
}
