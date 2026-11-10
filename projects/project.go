package projects

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/snackbag/wiki/render"
)

type Project struct {
	Description string   `json:"description"`
	Pages       []string `json:"pages"`

	links map[string]string
}

func Read(id string) (Project, error) {
	path := filepath.Join("wikidl", id, ".docs", "wiki.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return Project{}, err
	}

	var project Project
	err = json.Unmarshal(data, &project)
	if err != nil {
		return Project{}, err
	}

	project.links = make(map[string]string)
	for _, page := range project.Pages {
		project.links[render.Linkify(page)] = page
	}

	return project, nil
}
