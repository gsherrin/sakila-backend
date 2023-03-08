package filmactor

// comment for commit
import (
	"time"
)

type FilmActor struct {
	ActorId    int       `gorm:"many2many:actor;type:smallint;primaryKey"`
	FilmId     int       `gorm:"many2many:film;type:smallint;primaryKey"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}
