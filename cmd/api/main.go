package main

import (
	"log"
	"net/http"

	"movies-api/internal/config"
	"movies-api/internal/db"
	"movies-api/internal/routes"
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
	// movieRepo := repository.NewMovieRepository(database)
	// movies, err := movieRepo.GetAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, m := range movies {
	// 	fmt.Println(m.Title)
	// }

	log.Printf("Server listening on %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, handler))
}
