package routes

import (
	"database/sql"
	"net/http"

	"movies-api/internal/handlers"
	"movies-api/internal/middleware"
	"movies-api/internal/repository"
	"movies-api/internal/service"

	"github.com/go-playground/validator/v10"
)

func New(database *sql.DB) http.Handler {
	// Dependencies
	validate := validator.New()

	// Repositories
	actorRepo := repository.NewActorRepository(database)
	movieRepo := repository.NewMovieRepository(database)
	genreRepo := repository.NewGenreRepository(database)

	// Services
	actorService := service.NewActorService(actorRepo)
	movieService := service.NewMovieService(movieRepo)
	genreService := service.NewGenreService(genreRepo)

	// Handlers
	actorHandler := handlers.NewActorHandler(actorService, validate)
	movieHandler := handlers.NewMovieHandler(movieService, validate)
	genreHandler := handlers.NewGenreHandler(genreService, validate)

	// Routes
	mux := http.NewServeMux()

	// Actors
	mux.HandleFunc("GET /api/actors", actorHandler.GetAll)         // Retrieve all actors.
	mux.HandleFunc("GET /api/actors/{id}", actorHandler.GetByID)   // Retrieve actor by id.
	mux.HandleFunc("POST /api/actors", actorHandler.Create)        // Create new actor.
	mux.HandleFunc("PATCH /api/actors/{id}", actorHandler.Update)  // Partially update actor by id.
	mux.HandleFunc("DELETE /api/actors/{id}", actorHandler.Delete) // Delete actor by id.

	// Movies
	mux.HandleFunc("GET /api/movies", movieHandler.GetAll)        // Retrieve all movies.
	mux.HandleFunc("GET /api/movies/{id}", movieHandler.GetByID)  // Retrieve movie by id.
	mux.HandleFunc("POST /api/movies", movieHandler.Create)       // Create new movie.
	mux.HandleFunc("PATCH /api/movies/{id}", movieHandler.Update) // Partially update movie by id.
	// mux.HandleFunc("DELETE /api/movies/", ) // Delete actor by id.

	// Genres
	mux.HandleFunc("GET /api/genres", genreHandler.GetAll)         // Retrieve all genres.
	mux.HandleFunc("GET /api/genres/{id}", genreHandler.GetByID)   // Retrieve genre by id.
	mux.HandleFunc("POST /api/genres", genreHandler.Create)        // Create new genre.
	mux.HandleFunc("PATCH /api/genres/{id}", genreHandler.Update)  // Partially update genre by id.
	mux.HandleFunc("DELETE /api/genres/{id}", genreHandler.Delete) // Delete genre by id.

	// Middleware
	handler := middleware.Logger(mux)
	handler = middleware.Recovery(handler)

	return handler
}
