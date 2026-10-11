package server

import (
	"email-sender/pkg/server/mailactions"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router            *chi.Mux
	mailActionHandler http.Handler
}

func NewServer() (*Server, error) {
	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	server := &Server{
		router:            router,
		mailActionHandler: mailactions.NewMux(),
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	s.router.Mount("/mail-actions", http.StripPrefix("/mail-actions", s.mailActionHandler))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
