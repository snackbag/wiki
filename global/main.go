package global

import "fmt"

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
