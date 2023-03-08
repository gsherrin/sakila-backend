package films

// comment for commit
import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Get("/{id}", ListFilmWithID)
	router.Get("/search", ListFilmWithName)
	router.Post("/", CreateFilm)
	router.Delete("/{id}", DeleteFilm)
	router.Patch("/", UpdateFilm)
	return router
}
