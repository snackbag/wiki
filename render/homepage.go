package render

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/snackbag/wiki/projects"
)

var cacheHomepage string

func Homepage(projects map[string]projects.Project) string {
	if cacheHomepage == "" {
		path := filepath.Join("templates", "index.html")
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Sprintf("failed to get homepage: %s", err)
		}

		cacheHomepage = string(data)
	}

	cards := ""
	for id, project := range projects {
		cards += Card(id, project)
	}
	cacheHomepage = strings.ReplaceAll(cacheHomepage, "%cards%", cards)

	return cacheHomepage
}

func Card(id string, project projects.Project) string {
	b := strings.Builder{}

	// begin
	b.WriteString(`<a class="silent" href="/`)
	b.WriteString(id)
	b.WriteString(`"><div class="card">`)

	// title
	b.WriteString("<h2>")
	b.WriteString(project.Name)
	b.WriteString("</h2>")

	// description
	b.WriteString("<p>")
	b.WriteString(project.Description)
	b.WriteString("</p>")

	// close
	b.WriteString("</div></a>")

	return b.String()
}
