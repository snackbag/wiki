package main

import (
	"errors"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/snackbag/compass/compass"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

func AddWikiPages() {
	Server.AddRoute("/w/<project>", func(request compass.Request) compass.Response {
		project := request.GetParam("project")

		p, ok := Projects[project]
		if !ok {
			ctx := compass.NewTemplateContext(Server)
			ctx.SetVariable("message", "Project '"+project+"' not found")
			return compass.FillWithCode("404.html", ctx, Server, 404)
		}

		return GeneratePage(p, p.Main)
	})

	Server.AddRoute("/w/<project>/<page>", func(request compass.Request) compass.Response {
		project := request.GetParam("project")
		page := request.GetParam("page")

		p, ok := Projects[project]
		if !ok {
			ctx := compass.NewTemplateContext(Server)
			ctx.SetVariable("message", "Project '"+project+"' not found")
			return compass.FillWithCode("404.html", ctx, Server, 404)
		}

		return GeneratePage(p, page)
	})
}

func GeneratePage(project ProjectData, page string) compass.Response {
	matches, err := regexp.MatchString("^[a-zA-Z0-9-_]*$", page)
	if err != nil {
		Handler.DoFatalError("[Renderer/GeneratePage] Failed to match page with regex pattern. " + err.Error())
		return compass.Text("you're not supposed to see this, reload the page")
	}

	if !matches {
		ctx := compass.NewTemplateContext(Server)
		ctx.SetVariable("message", "The "+project.Display+" wiki does not have a page named '"+page+"'")
		return compass.FillWithCode("404.html", ctx, Server, 404)
	}

	p := path.Join(PagesDir, project.Id, page+".md")
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		ctx := compass.NewTemplateContext(Server)
		ctx.SetVariable("message", "The "+project.Display+" wiki does not have a page named '"+page+"'")
		return compass.FillWithCode("404.html", ctx, Server, 404)
	}

	file, err := os.Open(p)
	if err != nil {
		Handler.DoFatalError("[Renderer/GeneratePage] Failed to open file at '" + p + "'. " + err.Error())
		return compass.Text("you're not supposed to see this, reload the page")
	}

	defer file.Close()

	md, err := io.ReadAll(file)
	if err != nil {
		Handler.DoFatalError("[Renderer/GeneratePage] Failed to read file at '" + p + "'. " + err.Error())
		return compass.Text("you're not supposed to see this, reload the page")
	}

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	generated := markdown.ToHTML(md, nil, renderer)
	pages := GetPages(project)

	ctx := compass.NewTemplateContext(Server)
	ctx.SetVariable("content", ApplyStylingToContent(string(generated)))
	ctx.SetVariable("cp", project.Id)
	ctx.SetVariable("pages", BeautifyPages(pages, page, project))

	return compass.Fill("page.html", ctx, Server)
}

func GetPages(project ProjectData) []string {
	p := path.Join(PagesDir, project.Id)
	entries, err := os.ReadDir(p)
	if err != nil {
		Handler.DoFatalError("[Renderer/GetPages] Failed to ReadDir of '" + p + "'")
		return make([]string, 0)
	}

	pages := make([]string, 0)
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			pages = append(pages, strings.TrimSuffix(entry.Name(), ".md"))
		}
	}

	return pages
}

func BeautifyPages(pages []string, currentPage string, project ProjectData) string {
	builder := strings.Builder{}
	for _, page := range pages {
		builder.WriteString(`<a href="/w/`)
		builder.WriteString(project.Id)
		builder.WriteString("/")
		builder.WriteString(page)
		builder.WriteString(`" class="silent`)
		if currentPage == page {
			builder.WriteString(" active")
		}
		builder.WriteString(`"><div><p>`)
		builder.WriteString(strings.Title(strings.ReplaceAll(strings.ToLower(page), "-", " ")))
		builder.WriteString("</p></div></a>\n")
	}

	return builder.String()
}

func ApplyStylingToContent(content string) string {
	content = strings.ReplaceAll(content, "BEGIN NOTE", "<div class=\"note\">")
	content = strings.ReplaceAll(content, "END NOTE", "</div>")

	return content
}
