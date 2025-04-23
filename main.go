package main

import (
	"github.com/snackbag/compass/compass"
	"github.com/snackbag/wiki/global"
)

func main() {
	server := compass.NewServer()
	handler := global.NewHandler()

	config := AssembleConfiguration(handler)

	server.Port = config.Port
	server.StaticDirectory = config.StaticDir
	server.TemplatesDirectory = config.TemplatesDir
	server.ComponentsDirectory = config.ComponentsDir

	global.SetHandlerServerResponse(handler, &server)

	server.Start()
}
