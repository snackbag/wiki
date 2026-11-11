package projects

import (
	"fmt"
	"strings"

	"github.com/snackbag/compass/v2"
)

func Init(server *compass.Server, id string, def string) (Project, error) {
	server.Logger.Infof("Setting up project %s", id)
	if strings.HasPrefix(def, "git@") || strings.HasPrefix(def, "https://") {
		err := Download(id, def)
		if err != nil {
			return Project{}, fmt.Errorf("downloading %s: %w", id, err)
		}
	} else {
		err := CopyLocal(id, def)
		server.Logger.Warnf("%s may be incomplete. copying %s: %v", id, id, err)
	}

	project, err := Read(id)
	if err != nil {
		return Project{}, fmt.Errorf("reading wiki.json: %w", err)
	}

	return project, nil
}
