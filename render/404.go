package render

import (
	"fmt"
	"os"

	"github.com/snackbag/wiki/util"
)

var cacheNotFound string

func NotFound() string {
	if cacheNotFound == "" {
		path := util.PathTemplate("404.html")
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Sprintf("failed to get 404 page: %s", err)
		}

		cacheNotFound = string(data)
	}

	return cacheNotFound
}
