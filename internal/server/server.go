package server

import (
	"net/http"
	"time"
)

func NewServer(handler http.Handler, port string) *http.Server {
	return &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
