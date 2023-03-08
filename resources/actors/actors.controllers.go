package actors

// comment for commit
import (
	"log"
	"net/http"
	"strconv"

	db "example.com/sql-chi/database"
	e "example.com/sql-chi/error"
	"example.com/sql-chi/resources/filmactor"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)

	render.RenderList(w, r, NewActorListResponse(actors))
}

func ListActorWithID(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	var actors []*Actor
	db.DB.Where("actor_id = ?", IdParam).Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func ListActorWithName(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("aq")
	var actors []*Actor
	db.DB.Where("first_name LIKE ? OR last_name LIKE ?", "%"+q+"%", "%"+q+"%").Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func ListFilmsWithActorId(w http.ResponseWriter, r *http.Request) {
	IdParam := chi.URLParam(r, "id")
	var actors []*Actor
	db.DB.Where("actor_id = ?", IdParam).Find(&actors)
	var filmActors []*filmactor.FilmActor
	db.DB.Where("actor_id = ?", IdParam).Find(&filmActors)
	db.DB.Table("film_actor").Select("film_actor.film_id, film.title").Joins("left join film on film.film_id = film_actor.film_id").Scan(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

// no actor ID needed when creating. can remove this param in postman
func CreateActor(w http.ResponseWriter, r *http.Request) {

	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor

	db.DB.Create(actor)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	IdParam, _ := strconv.Atoi(chi.URLParam(r, "id"))
	db.DB.Delete(&Actor{}, IdParam)
	log.Printf("Actor with ID %d was deleted.", IdParam)
}

// actor ID needed when updating, maybe fix in future
func UpdateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor

	db.DB.Save(actor)

	log.Printf("Actor with ID %d was updated.", actor.ActorId)
	render.Render(w, r, NewActorResponse(actor))
}
