package main

import (
	"errors"
	"github.com/gomarkdown/markdown"
	"github.com/snackbag/compass/compass"
	"io"
	"os"
	"path"
	"regexp"
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

	html := markdown.ToHTML(md, nil, nil)

	ctx := compass.NewTemplateContext(Server)
	ctx.SetVariable("content", string(html))
	ctx.SetVariable("cp", project.Id)

	return compass.Fill("page.html", ctx, Server)
}
