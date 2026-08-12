package routes

import (
	"database/sql"
	"net/http"

	"movies-api/internal/handlers"
	"movies-api/internal/middleware"
	"movies-api/internal/repository"
	"movies-api/internal/service"
)

func New(database *sql.DB) http.Handler {
	// Repositories
	actorRepo := repository.NewActorRepository(database)
	movieRepo := repository.NewMovieRepository(database)
	genreRepo := repository.NewGenreRepository(database)

	// Services
	actorService := service.NewActorService(actorRepo)
	movieService := service.NewMovieService(movieRepo)
	genreService := service.NewGenreService(genreRepo)

	// Handlers
	actorHandler := handlers.NewActorHandler(actorService)
	movieHandler := handlers.NewMovieHandler(movieService)
	genreHandler := handlers.NewGenreHandler(genreService)

	// Routes
	mux := http.NewServeMux()

	// Actors
	mux.HandleFunc("GET /api/actor", actorHandler.GetAll) // Retrieve all actors.

	// Movies
	mux.HandleFunc("GET /api/movie", movieHandler.GetAll) // Retrieve all actors.

	// Genres
	mux.HandleFunc("GET /api/genre", genreHandler.GetAll) // Retrieve all actors.

	// Middleware
	handler := middleware.Logger(mux)
	handler = middleware.Recovery(handler)

	return handler
}
