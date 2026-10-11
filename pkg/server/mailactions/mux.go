package mailactions

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux() *chi.Mux {
	r := chi.NewMux()

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, http.StatusOK, page{
			Title: "hello world",
			Body:  "Click the button below to confirm that you want to receive updates.",
		})
	})

	return r
}
