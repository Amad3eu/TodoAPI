package main

import (
	"github.com/Amad3eu/api-pstg-go/configs"
	"github.com/Amad3eu/api-pstg-go/handlers"
	"net/http"
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
	r.get("/{id}", handlers.Get)

	http.ListenAndServe(fm.Sprintf("%s", configs.GetServerPort()), r)
}
