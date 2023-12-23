package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"golang_htmx_teml/internal/views"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/mavolin/go-htmx"
)

type Server struct {
	Router *chi.Mux
	Logger *httplog.Logger
}

func (s *Server) InitServer() {
	staticDir := http.Dir("internal/assets")

	// Logger
	logger := httplog.NewLogger("Server Logger", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelInfo,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	s.Logger = logger

	s.Router.Use(htmx.NewMiddleware())
	s.Router.Use(httplog.RequestLogger(logger))
	s.Router.Use(middleware.Heartbeat("/ping"))
	// s.Router.Use(middleware.Logger)
	s.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServer(staticDir)))
}

func (s *Server) InitRoutes() {
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		htmxReq := htmx.Request(r)
		if htmxReq != nil && htmxReq.Boosted {
			fmt.Println("Request is boosted by HTMX")
		}
		// component := views.Hello("John")
		// component := views.Header("John")
		// component.Render(r.Context(), w)
		component := views.Page("John")
		component.Render(r.Context(), w)
	})

	s.Router.Post("/clicked", func(w http.ResponseWriter, r *http.Request) {
		htmxReq := htmx.Request(r)
		if htmxReq != nil && htmxReq.Boosted {
			fmt.Println("Request is boosted by HTMX")
		}
		w.Write([]byte("Clicked!"))
	})

	s.Router.Post("/example", func(w http.ResponseWriter, r *http.Request) {
		htmxReq := htmx.Request(r)
		if htmxReq != nil && htmxReq.Boosted {
			fmt.Println("Request is boosted by HTMX")
		}
		w.Write([]byte("Clicked!"))
	})
}

func (s *Server) StartServer() {
	ctx := context.Background()
	s.Logger.Log(ctx, slog.LevelInfo, "Server started on port 3000")

	err := http.ListenAndServe(":3000", s.Router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
