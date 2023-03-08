package actors

// comment for commit
import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListActors)
	router.Get("/{id}", ListActorWithID)
	router.Get("/search", ListActorWithName)
	// router.Get("/{id}films", ListFilmsWithActorId)
	router.Post("/", CreateActor)
	router.Delete("/{id}", DeleteActor)
	router.Patch("/", UpdateActor)
	return router
}
