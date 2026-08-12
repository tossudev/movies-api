package main

import (
	"log"
	"fmt"
	"net/http"

	"movies-api/internal/config"
	"movies-api/internal/db"
	"movies-api/internal/routes"
	"movies-api/internal/repository"
)

func main() {
	cfg := config.Load()
	database, err := db.Open(cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := db.Migrate(database); err != nil {
		log.Fatal(err)
	}

	handler := routes.New(database)

	// ! [TESTING FUNCTIONALITY]:

	repository.InitActorRepository(database)
	repository.InitMovieRepository(database)
	repository.InitGenreRepository(database)
	
	fmt.Println(repository.Movie.Exists(1))		// true
	fmt.Println(repository.Genre.Exists(1))		// true
	fmt.Println(repository.Actor.Exists(1))		// true
	fmt.Println(repository.Movie.Exists(-1))	// false

	movie, err := repository.Movie.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movie)

	log.Printf("Server listening on %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, handler))
}
