package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	logger *log.Logger
	http   *http.Server
}

func newServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.IndexHandler)
	r.Post("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		logger: logger,
		http:   httpServer,
	}
}

func Main() {
	logger := log.New(os.Stdout, "[server]", log.LstdFlags)

	server := newServer(logger)

	logger.Println("Server is starting on :8080")

	if err := server.http.ListenAndServe(); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
