package main

import (
	"golang_htmx_teml/internal"

	"github.com/go-chi/chi"
)

func main() {
	s := internal.Server{Router: chi.NewRouter()}
	s.InitServer()
	s.InitRoutes()
	s.StartServer()
}
