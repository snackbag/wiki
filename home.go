package main

import (
	"github.com/snackbag/compass/compass"
)

func AddHomePages() {
	Server.AddRoute("/", func(request compass.Request) compass.Response {
		ctx := compass.NewTemplateContext(Server)
		return compass.Fill("index.html", ctx, Server)
	})
}
