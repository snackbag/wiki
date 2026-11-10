package projects

import (
	"encoding/json"
	"os"

	"github.com/snackbag/wiki/util"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pages       []string `json:"pages"`

	Links map[string]string
}

func Read(id string) (Project, error) {
	path := util.PathWikidl(id, ".docs", "wiki.json")
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
