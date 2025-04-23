package global

import (
	"fmt"
	"github.com/snackbag/compass/compass"
	"strings"
)

type Handler struct {
	Errors []string
}

func (handler *Handler) DoFatalError(err string) {
	handler.Errors = append(handler.Errors, err)

	fmt.Println("\033[1;31mEncountered fatal error:\033[0;31m", err, "\033[0m")
}

func NewHandler() *Handler {
	return &Handler{Errors: make([]string, 0)}
}

func SetHandlerServerResponse(handler *Handler, server *compass.Server) {
	server.SetBeforeRequestHandler(func(request compass.Request) *compass.Response {
		if len(handler.Errors) > 0 {
			ctx := compass.NewTemplateContext(server)

			builder := strings.Builder{}
			for _, value := range handler.Errors {
				builder.WriteString("<li>")
				builder.WriteString(value)
				builder.WriteString("</li>\n")
			}

			ctx.SetVariable("errors", builder.String())

			resp := compass.Fill("fatal_error.html", ctx, server)
			return &resp
		}

		return nil
	})
}
