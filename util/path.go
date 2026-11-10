package util

import "path/filepath"

func PathTemplate(file ...string) string {
	parts := append([]string{"assets", "templates"}, file...)
	return filepath.Join(parts...)
}

func PathWikidl(id string, file ...string) string {
	parts := append([]string{"wikidl", id}, file...)
	return filepath.Join(parts...)
}
