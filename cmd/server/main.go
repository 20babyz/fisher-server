package main

import (
	"api/internal/config"
	"api/internal/handlers"
	"api/internal/middleware"
	"api/internal/server"
	"api/pkg/db"

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
