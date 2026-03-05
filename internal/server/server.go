package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	logger *log.Logger
	http   *http.Server
}

func New(logger *log.Logger) *Server {
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

func (s *Server) Start() error {
	s.logger.Println("Server is starting on :8080")

	if err := s.http.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
