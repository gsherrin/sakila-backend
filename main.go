package main

// comment for commit
import (
	"example.com/sql-chi/database"
	"example.com/sql-chi/resources/actors"
	"example.com/sql-chi/resources/filmactor"
	"example.com/sql-chi/resources/films"
	"example.com/sql-chi/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&filmactor.FilmActor{})
	database.DB.AutoMigrate(&actors.Actor{})
	database.DB.AutoMigrate(&films.Film{})
	server.Init()

}
