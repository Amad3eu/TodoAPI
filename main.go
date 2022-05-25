package main

import (
	"fmt"
	"net/http"

	"github.com/aprendagolang/api-pqsql/configs"
	"github.com/go-chi/chi/v5"
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
