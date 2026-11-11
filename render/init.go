package render

import (
	"fmt"

	"github.com/snackbag/compass/v2"
	"github.com/snackbag/wiki/projects"
)

func Init(server *compass.Server, id string, project projects.Project) {
	server.AddRoute(fmt.Sprintf("/%s", id), func(request compass.Request) compass.Response {
		return compass.PermaRedirect(fmt.Sprintf("/%s/home", id))
	})

	server.AddRoute(fmt.Sprintf("/%s/<page>", id), func(request compass.Request) compass.Response {
		page, _ := request.GetRouteParam("page")

		rendered, err := Page(id, project.Links[page], project.Pages)
		if err != nil {
			server.Logger.Errorf("failed to render %s/%s: %v", id, page, err)
			return compass.TextWithCode("Failed to render page", 500)
		}

		return compass.HTML(rendered)
	})
}
