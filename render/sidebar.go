package render

import (
	"strings"

	"github.com/snackbag/wiki/util"
)

// Sidebar updates <@Sidebar/> in the buffer HTML file, and returns an updated buffer
func Sidebar(allPages []string, active string, buf string) string {
	var sidebar strings.Builder

	for _, page := range allPages {
		if strings.HasPrefix(page, "*") {
			sidebar.WriteString("<p>")
			sidebar.WriteString(strings.TrimPrefix(page, "*"))
			sidebar.WriteString("</p>")
		} else {
			sidebar.WriteString(`<a href="`)
			sidebar.WriteString(util.Linkify(page))
			sidebar.WriteString(`" class="silent`)

			if page == active {
				sidebar.WriteString(" active")
			}

			sidebar.WriteString(`"><div><p>`)
			sidebar.WriteString(util.Titleify(page))
			sidebar.WriteString("</p></div></a>")
		}

		sidebar.WriteString("\n")
	}

	return strings.ReplaceAll(buf, "<@Sidebar/>", sidebar.String())
}
