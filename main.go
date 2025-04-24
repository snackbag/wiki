package main

import (
	"fmt"
	"github.com/snackbag/compass/compass"
	"github.com/snackbag/wiki/global"
)

var Server *compass.Server
var Projects map[string]ProjectData
var PagesDir string
var Handler *global.Handler

func main() {
	server := compass.NewServer()
	handler := global.NewHandler()
	Server = &server
	Handler = handler

	config := AssembleConfiguration(handler)
	if config == nil {
		fmt.Println("Config is nil, aborting startup.")
		return
	}

	server.Port = config.Port
	server.StaticDirectory = config.StaticDir
	server.TemplatesDirectory = config.TemplatesDir
	server.ComponentsDirectory = config.ComponentsDir
	PagesDir = config.PagesDir

	global.SetHandlerServerResponse(handler, &server)
	LoadProjects(config.PagesDir, handler)

	AddHomePages()
	AddWikiPages()

	server.Start()
}
