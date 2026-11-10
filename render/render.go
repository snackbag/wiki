package render

import (
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/snackbag/wiki/util"
)

func Render(id string, page string, allPages []string) (string, error) {
	path := util.PathWikidl(id, ".docs", page)
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
