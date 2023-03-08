package filmactor

// comment for commit
import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilmsWithActors)
	router.Get("/{id}", ListFilmActorWithID)
	return router
}
