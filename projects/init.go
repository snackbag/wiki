package projects

import (
	"fmt"
	"strings"

	"github.com/snackbag/compass/v2"
)

func Init(server *compass.Server, id string, def string) error {
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
		return compass.Text(fmt.Sprintf("page %s of %s", page, id))
	})

	return nil
}
