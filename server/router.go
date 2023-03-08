package server

// comment for commit
import (
	"example.com/sql-chi/resources/actors"
	"example.com/sql-chi/resources/filmactor"
	"example.com/sql-chi/resources/films"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Router() chi.Router {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Mount("/filmactor", filmactor.Routes())
	router.Mount("/films", films.Routes())
	router.Mount("/actors", actors.Routes())

	return router
}
