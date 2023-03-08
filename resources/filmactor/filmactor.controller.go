package filmactor

// comment for commit
import (
	"errors"
	"net/http"

	db "example.com/sql-chi/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListFilmsWithActors(w http.ResponseWriter, r *http.Request) {
	var filmsActors []*FilmActor
	db.DB.Find(&filmsActors)
	render.RenderList(w, r, NewFilmActorListResponse(filmsActors))
}

func ListFilmActorWithID(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	var filmsactors []*FilmActor
	db.DB.Where("actor_id = ?", IdParam).Find(&filmsactors)
	render.RenderList(w, r, NewFilmActorListResponse(filmsactors))
}

type FilmActorRequest struct {
	*FilmActor
}

func (a *FilmActorRequest) Bind(r *http.Request) error {
	if a.FilmActor == nil {
		return errors.New("missing required Actor fields")
	}

	return nil
}

type FilmActorResponse struct {
	*FilmActor
}

func NewFilmActorResponse(filmactor *FilmActor) *FilmActorResponse {
	return &FilmActorResponse{filmactor}
}

func NewFilmActorListResponse(filmactors []*FilmActor) []render.Renderer {
	list := []render.Renderer{}
	for _, filmactor := range filmactors {
		list = append(list, NewFilmActorResponse(filmactor))
	}
	return list
}

func (a *FilmActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
