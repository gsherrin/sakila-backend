package films

// comment for commit
import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId             int    `gorm:"many2many:filmactor;type:smallint;primaryKey"`
	Title              string `gorm:"type:varchar(128)"`
	Description        string `gorm:"type:text"`
	ReleaseYear        int    `gorm:"type:year"`
	OriginalLanguageId int    `gorm:"type:tinyint;default:null"`
	LanguageId         int    `gorm:"type:tinyint;not null"`
	Length             int    `gorm:"type:smallint"`
	// Rating             string    `gorm:"type:enum('G','PG','PG-13','R','NC-17');column:rating"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

func (a *FilmRequest) Bind(r *http.Request) error {
	if a.Film == nil {
		return errors.New("missing required Film fields")
	}

	a.Film.Title = strings.ToUpper(a.Film.Title)

	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(films *Film) *FilmResponse {
	return &FilmResponse{films}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (a *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
