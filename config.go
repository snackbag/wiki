package main

import (
	"encoding/json"
	"github.com/snackbag/wiki/global"
	"os"
)

type Configuration struct {
	Port          int    `json:"port"`
	StaticDir     string `json:"static_dir"`
	TemplatesDir  string `json:"templates_dir"`
	ComponentsDir string `json:"components_dir"`
	PagesDir      string `json:"pages_dir"`
}

func AssembleConfiguration(handler *global.Handler) *Configuration {
	f, err := os.OpenFile("config.json", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		handler.DoFatalError("[AssembleConfiguration] Failed top open file 'config.json' (perm 0755): " + err.Error())
		return nil
	}

	defer f.Close()
	defaultConfig := Configuration{
		Port:          3000,
		StaticDir:     "static",
		TemplatesDir:  "templates",
		ComponentsDir: "components",
		PagesDir:      "pages",
	}

	var config Configuration
	fileInfo, err := f.Stat()
	if err != nil {
		handler.DoFatalError("[AssembleConfiguration] Failed to get file info for 'config.json'. " + err.Error())
		return nil
	}

	config = defaultConfig
	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(f)
		if err = decoder.Decode(&config); err != nil {
			handler.DoFatalError("[AssembleConfiguration] Failed to parse JSON on config 'config.json'. " + err.Error())
			return nil
		}

		if config.Port == 0 {
			config.Port = defaultConfig.Port
		}

		if config.StaticDir == "" {
			config.StaticDir = defaultConfig.StaticDir
		}

		if config.TemplatesDir == "" {
			config.TemplatesDir = defaultConfig.TemplatesDir
		}

		if config.ComponentsDir == "" {
			config.ComponentsDir = defaultConfig.ComponentsDir
		}

		if config.PagesDir == "" {
			config.PagesDir = defaultConfig.PagesDir
		}
	}

	updatedJSON, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		handler.DoFatalError("[AssembleConfiguration] Failed while marshalling updated JSON of 'config.json'. " + err.Error())
		return nil
	}

	err = os.WriteFile("config.json", updatedJSON, 0755)
	if err != nil {
		handler.DoFatalError("[AssembleConfiguration] Failed to write updated to 'config.json'. " + err.Error())
		return nil
	}

	return &config
}
