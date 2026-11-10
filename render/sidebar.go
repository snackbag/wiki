package render

import "strings"

// Sidebar updates <@Sidebar/> in the buffer HTML file, and returns an updated buffer
func Sidebar(allPages []string, buf string) string {
	var sidebar strings.Builder

	for _, page := range allPages {
		if strings.HasPrefix(page, "*") {
			sidebar.WriteString("<p>")
			sidebar.WriteString(page)
			sidebar.WriteString("</p>")
		} else {
			sidebar.WriteString(`<a href="`)
			sidebar.WriteString(linkify(page))
			sidebar.WriteString(`" class="silent"><div><p>`)
			sidebar.WriteString(titleify(page))
			sidebar.WriteString("</p></div></a>")
		}

		sidebar.WriteString("\n")
	}

	return strings.ReplaceAll(buf, "<@Sidebar/>", sidebar.String())
}

// titleify removes the .md end of a filename and replaces all underscores with spaces.
func titleify(filename string) string {
	filename = strings.TrimSuffix(filename, ".md")
	filename = strings.ReplaceAll(filename, "_", " ")
	return filename
}

// linkify removes the .md end of a filename, makes all letters lowercase, and replaces
// all underscores with dashes
func linkify(filename string) string {
	filename = strings.TrimSuffix(filename, ".md")
	filename = strings.ToLower(filename)
	filename = strings.ReplaceAll(filename, "_", "-")
	return filename
}
