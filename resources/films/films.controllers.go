package films

// comment for commit
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "example.com/sql-chi/database"
	e "example.com/sql-chi/error"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func ListFilmWithID(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	var films []*Film
	db.DB.Where("film_id = ?", IdParam).Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func ListFilmWithName(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("aq")
	var films []*Film
	db.DB.Where("title LIKE ? OR release_year LIKE ?", "%"+q+"%", "%"+q+"%").Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {

	var data FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	film := data.Film

	db.DB.Create(film)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewFilmResponse(film))
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	IdParam, _ := strconv.Atoi(chi.URLParam(r, "id"))
	db.DB.Delete(&Film{}, IdParam)
	log.Printf("Film with ID %d was deleted.", IdParam)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var data FilmRequest
	// if err := render.Bind(r, &data); err != nil {
	// 	render.Render(w, r, e.ErrInvalidRequest(err))
	// }
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
		return
	}

	film := data.Film

	db.DB.Save(film)

	log.Printf("Film with ID %d was updated.", film.FilmId)
	render.Render(w, r, NewFilmResponse(film))
}
