package main

import (
	"github.com/Amad3eu/api-pstg-go/configs"
	"github.com/Amad3eu/api-pstg-go/handlers"
	"github.com/go-chi/chi/v5"
	http "net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fm.Sprintf("%s", configs.GetServerPort()), r)
}
