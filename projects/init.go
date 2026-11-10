package projects

import (
	"fmt"
	"strings"

	"github.com/snackbag/compass/v2"
	"github.com/snackbag/wiki/render"
)

func Init(server *compass.Server, id string, def string) error {
	server.Logger.Info(fmt.Sprintf("Setting up project %s", id))
	if strings.HasPrefix(def, "git@") || strings.HasPrefix(def, "https://") {
		err := Download(id, def)
		if err != nil {
			return fmt.Errorf("downloading %s: %w", id, err)
		}
	} else {
		err := CopyLocal(id, def)
		if err != nil {
			return fmt.Errorf("copying %s: %w", id, err)
		}
	}

	server.AddRoute(fmt.Sprintf("/%s", id), func(request compass.Request) compass.Response {
		return compass.PermaRedirect(fmt.Sprintf("/%s/home", id))
	})

	server.AddRoute(fmt.Sprintf("/%s/<page>", id), func(request compass.Request) compass.Response {
		page, _ := request.GetRouteParam("page")

		rendered, err := render.Render(id, page, []string{"hi.md", "*hiow", "woah_hI.md"})
		if err != nil {
			server.Logger.Error(fmt.Sprintf("failed to render %s/%s: %v", id, page, err))
			return compass.TextWithCode("Failed to render page", 500)
		}

		return compass.HTML(rendered)
	})

	return nil
}
