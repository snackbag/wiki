package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/snackbag/compass/v2"
	"github.com/snackbag/wiki/projects"
	"github.com/snackbag/wiki/render"
)

var server *compass.Server
var Config struct {
	Projects       map[string]string `json:"projects"`
	LoadedProjects map[string]projects.Project
}

func main() {
	server = compass.NewServer(compass.NewStandardConfiguration())

	// load config
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(fmt.Errorf("reading config.json: %w", err))
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Fatal(fmt.Errorf("parsing config.json: %w", err))
	}

	// register projects
	for id, def := range Config.Projects {
		loadedProject, err := projects.Init(server, id, def)
		if err != nil {
			server.Logger.Error(fmt.Sprintf("failed to set up project %s: %v", id, err))
			server.AddRoute(fmt.Sprintf("/%s", id), func(request compass.Request) compass.Response {
				return compass.TextWithCode("An error occurred while setting up this wiki page.", 500)
			})
		} else {
			Config.LoadedProjects[id] = loadedProject
		}

		render.Init(server, id, loadedProject)
	}

	// legacy url fallback
	server.AddRoute("/w/<project>", func(request compass.Request) compass.Response {
		project, _ := request.GetRouteParam("project")
		return compass.PermaRedirect(fmt.Sprintf("/%s", project))
	})

	server.AddRoute("/w/<project>/<page>", func(request compass.Request) compass.Response {
		project, _ := request.GetRouteParam("project")
		page, _ := request.GetRouteParam("page")
		return compass.PermaRedirect(fmt.Sprintf("/%s/%s", project, page))
	})

	// 404 page
	server.NotFoundHandler = func(request compass.Request) compass.Response {
		return compass.HTMLWithCode(render.NotFound(), 404)
	}

	server.MustRun()
}
