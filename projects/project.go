package projects

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/snackbag/wiki/util"
)

type Project struct {
	Description string   `json:"description"`
	Pages       []string `json:"pages"`

	Links map[string]string
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

	project.Links = make(map[string]string)
	for _, page := range project.Pages {
		project.Links[util.Linkify(page)] = page
	}

	return project, nil
}
