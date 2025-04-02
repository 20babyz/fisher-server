package handlers

import (
	"net/http"
)

type ExampleHandler struct{}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Example API"))

	if err != nil {
		return
	}
}
