package main

import (
	"encoding/json"
	"github.com/snackbag/wiki/global"
	"io"
	"os"
	"path"
)

type ProjectData struct {
	Id                  string   `json:"id"`
	Display             string   `json:"display"`
	Source              string   `json:"source"`
	Issues              string   `json:"issues"`
	Main                string   `json:"main"`
	DownloadInfo        string   `json:"download_info"`
	DownloadPrecompiled string   `json:"download_precompiled"`
	PageStructure       []string `json:"page_structure"`

	CachedPageStructure []string
	CachedPageContents  map[string][]byte
}

func LoadProjects(pagesDir string, handler *global.Handler) {
	Projects = make(map[string]*ProjectData)

	file, err := os.Open(path.Join(pagesDir, "struct.json"))
	if err != nil {
		handler.DoFatalError("[LoadProjects] Failed to open 'struct.json'. " + err.Error())
		return
	}

	defer file.Close()
	byteVal, _ := io.ReadAll(file)

	err = json.Unmarshal(byteVal, &Projects)
	if err != nil {
		handler.DoFatalError("[LoadProjects] Failed to unmarshal 'struct.json'. " + err.Error())
		return
	}

	for _, proj := range Projects {
		proj.CachedPageContents = make(map[string][]byte)
	}
}
