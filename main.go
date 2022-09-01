package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheusgb/crud-library/configs"
	"github.com/matheusgb/crud-library/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
