package util

import "strings"

// Titleify removes the .md end of a filename and replaces all underscores with spaces.
func Titleify(filename string) string {
	filename = strings.TrimSuffix(filename, ".md")
	filename = strings.ReplaceAll(filename, "_", " ")
	return filename
}

// Linkify removes the .md end of a filename, makes all letters lowercase, and replaces
// all underscores with dashes
func Linkify(filename string) string {
	filename = strings.TrimSuffix(filename, ".md")
	filename = strings.ToLower(filename)
	filename = strings.ReplaceAll(filename, "_", "-")
	return filename
}
