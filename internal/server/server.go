package server

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Server() {
	r := chi.NewRouter()

	r.Get("/", handlers.IndexHandler)
	r.Post("/upload", handlers.UploadHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка сервера: %s", err.Error())
		return
	}
}
