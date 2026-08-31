package routes

import (
	"database/sql"
	"net/http"

	"movies-api/internal/handlers"
	"movies-api/internal/middleware"
	"movies-api/internal/repository"
	"movies-api/internal/service"
	"movies-api/internal/validator"
)

func New(database *sql.DB) http.Handler {
	// Dependencies
	validator.Init()

	// Repositories
	actorRepo := repository.NewActorRepository(database)
	movieRepo := repository.NewMovieRepository(database)
	genreRepo := repository.NewGenreRepository(database)

	// Services
	actorService := service.NewActorService(actorRepo)
	movieService := service.NewMovieService(movieRepo)
	genreService := service.NewGenreService(genreRepo)

	// Handlers
	actorHandler := handlers.NewActorHandler(actorService, validator.Validate)
	movieHandler := handlers.NewMovieHandler(movieService, validator.Validate)
	genreHandler := handlers.NewGenreHandler(genreService, validator.Validate)

	// Routes
	mux := http.NewServeMux()

	// Actors
	mux.HandleFunc("GET /api/actors", actorHandler.GetAll)                // Retrieve all actors.
	mux.HandleFunc("GET /api/actors/{id}", actorHandler.GetByID)          // Retrieve actor by id.
	mux.HandleFunc("GET /api/actors/{id}/movies", actorHandler.GetMovies) // Retrieve movies from actor by id.
	mux.HandleFunc("POST /api/actors", actorHandler.Create)               // Create new actor.
	mux.HandleFunc("PATCH /api/actors/{id}", actorHandler.Update)         // Partially update actor by id.
	mux.HandleFunc("DELETE /api/actors/{id}", actorHandler.Delete)        // Delete actor by id.

	// Movies
	mux.HandleFunc("GET /api/movies", movieHandler.GetAll)                // Retrieve all movies.
	mux.HandleFunc("GET /api/movies/search", movieHandler.Search)         // Search movies
	mux.HandleFunc("GET /api/movies/{id}", movieHandler.GetByID)          // Retrieve movie by id.
	mux.HandleFunc("GET /api/movies/{id}/actors", movieHandler.GetActors) // Retrieve actors in movie by id.
	mux.HandleFunc("GET /api/movies/{id}/genres", movieHandler.GetGenres) // Retrieve genres in movie by id.
	mux.HandleFunc("POST /api/movies", movieHandler.Create)               // Create new movie.
	mux.HandleFunc("PATCH /api/movies/{id}", movieHandler.Update)         // Partially update movie by id.
	mux.HandleFunc("DELETE /api/movies/{id}", movieHandler.Delete)        // Delete movie by id.

	// Genres
	mux.HandleFunc("GET /api/genres", genreHandler.GetAll)                // Retrieve all genres.
	mux.HandleFunc("GET /api/genres/{id}", genreHandler.GetByID)          // Retrieve genre by id.
	mux.HandleFunc("GET /api/genres/{id}/movies", genreHandler.GetMovies) // Retrieve actors in movie by id.
	mux.HandleFunc("POST /api/genres", genreHandler.Create)               // Create new genre.
	mux.HandleFunc("PATCH /api/genres/{id}", genreHandler.Update)         // Partially update genre by id.
	mux.HandleFunc("DELETE /api/genres/{id}", genreHandler.Delete)        // Delete genre by id.

	// Middleware
	handler := middleware.Logger(mux)
	handler = middleware.Recovery(handler)

	return handler
}
