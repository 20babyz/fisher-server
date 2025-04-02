package main

import (
	"crawler/internal/config"
	"crawler/internal/handlers"
	"crawler/internal/middleware"
	"crawler/internal/server"
	"crawler/pkg/db"

	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Println("Starting server...")

	if err := db.InitDB(config.DBUser, config.DBPassword, config.DBHost, config.DBName); err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/examples", middleware.Logging(handlers.NewExampleHandler()))

	srv := server.NewServer(mux, config.Port)

	fmt.Println("Server running on port", config.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
