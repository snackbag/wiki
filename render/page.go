package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/snackbag/wiki/util"
)

var cachedPageBase string

func Page(id string, page string, allPages []string) (string, error) {
	// html
	if cachedPageBase == "" {
		path := util.PathTemplate("page.html")
		data, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("failed to get page base: %w", err)
		}

		cachedPageBase = string(data)
	}

	// content
	md, err := getPageContentHTML(id, page)
	if err != nil {
		return "", fmt.Errorf("failed to convert md: %w", err)
	}
	content := strings.ReplaceAll(cachedPageBase, "<@Content/>", md)
	content = Sidebar(allPages, page, content)

	return content, nil
}

func md2Html(md []byte) string {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return string(markdown.ToHTML(md, nil, renderer))
}

func getPageContentHTML(id string, page string) (string, error) {
	path := util.PathWikidl(id, ".docs", page)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return md2Html(data), nil
}
