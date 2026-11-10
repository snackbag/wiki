package render

import (
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func Render(id string, page string, allPages []string) (string, error) {
	path := filepath.Join("wikidl", id, ".docs", page)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	buf := md2Html(data)
	buf = Sidebar(allPages, page, buf)

	return buf, nil
}

func md2Html(md []byte) string {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return string(markdown.ToHTML(md, nil, renderer))
}
