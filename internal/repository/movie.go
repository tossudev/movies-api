package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"movies-api/internal/dto"
	"movies-api/internal/models"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetAll(page, size int) ([]models.Movie, error) {
	pagination, args := "", []any{}
	if size != 0 { // Checks whetever pagination exists or not, no other possibility of size being 0.
		pagination = " LIMIT ? OFFSET ?"
		args = []any{size, page*size - size}
	}

	rows, err := r.db.Query("SELECT * FROM movie"+pagination, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.getMoviesFromRows(rows)
}

func (r *MovieRepository) Search(query string, page, size int) ([]models.Movie, error) {
	pagination, args := "", []any{"%" + query + "%"}
	if size != 0 { // Checks whetever pagination exists or not, no other possibility of size being 0.
		pagination = " LIMIT ? OFFSET ?"
		args = append(args, size, page*size-size)
	}

	rows, err := r.db.Query("SELECT * FROM movie WHERE title LIKE ?"+pagination, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.getMoviesFromRows(rows)
}

func (r *MovieRepository) GetByID(id int) (models.Movie, error) {
	var movie models.Movie
	query := "SELECT * FROM movie WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
		return movie, err // TODO: concretize error
	}

	movie.Actors, _ = r.GetActors(movie.ID)
	movie.Genres, _ = r.GetGenres(movie.ID)

	return movie, nil
}

func (r *MovieRepository) Create(req dto.CreateMovieRequest) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO movie (title, release_year, duration) VALUES (?, ?, ?)", req.Title, req.ReleaseYear, req.Duration)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	movieID := int(id64)

	for _, genreID := range req.GenreIDs {
		if !r.GenreExists(genreID) {
			return 0, fmt.Errorf(IDNotFound, "genre", genreID)
		}

		if _, err := tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)", movieID, genreID); err != nil {
			return 0, err
		}
	}

	for _, actorID := range req.ActorIDs {
		if !r.ActorExists(actorID) {
			return 0, fmt.Errorf(IDNotFound, "actor", actorID)
		}

		if _, err := tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)", movieID, actorID); err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return movieID, nil
}

func (r *MovieRepository) Update(id int, req dto.UpdateMovieRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	var sets []string
	var args []any

	if req.Title != nil {
		sets = append(sets, "title = ?")
		args = append(args, *req.Title)
	}

	if req.ReleaseYear != nil {
		sets = append(sets, "release_year = ?")
		args = append(args, *req.ReleaseYear)
	}

	if req.Duration != nil {
		sets = append(sets, "duration = ?")
		args = append(args, *req.Duration)
	}

	if len(sets) == 0 {
		return nil // or return an error
	}

	args = append(args, id)
	query := "UPDATE movie SET " + strings.Join(sets, ", ") + " WHERE id = ?"

	result, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return err
	}
	movieID := int(id64)

	for _, genreID := range req.GenreIDs {
		if !r.GenreExists(genreID) {
			return fmt.Errorf(IDNotFound, "genre", genreID)
		}

		if _, err := tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)", movieID, genreID); err != nil {
			return err
		}
	}

	for _, actorID := range req.ActorIDs {
		if !r.ActorExists(actorID) {
			return fmt.Errorf(IDNotFound, "actor", actorID)
		}

		if _, err := tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)", movieID, actorID); err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (r *MovieRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM movie WHERE id = ?", id)
	return err
}

func (r *MovieRepository) GetActors(movieID int) ([]models.Actor, error) {
	rows, err := r.db.Query("SELECT * FROM actor WHERE id in (SELECT actor_id FROM movie_actors WHERE movie_id = ?)", movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actors := []models.Actor{}
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
			return nil, err // TODO: concretize error
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return actors, nil
}

func (r *MovieRepository) GetGenres(movieID int) ([]models.Genre, error) {
	rows, err := r.db.Query("SELECT * FROM genre WHERE id in (SELECT genre_id FROM movie_genres WHERE movie_id = ?)", movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []models.Genre{}
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, err // TODO: concretize error
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return genres, nil
}

func (r *MovieRepository) GenreExists(id int) bool {
	return r.db.QueryRow("SELECT * FROM genre WHERE id = ?", id).Err() == nil
}

func (r *MovieRepository) ActorExists(id int) bool {
	return r.db.QueryRow("SELECT * FROM actor WHERE id = ?", id).Err() == nil
}

func (r *MovieRepository) getMoviesFromRows(rows *sql.Rows) ([]models.Movie, error) {
	var err error
	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
			return nil, err // TODO: concretize error
		}

		if movie.Actors, err = r.GetActors(movie.ID); err != nil {
			return nil, err
		}
		if movie.Genres, err = r.GetGenres(movie.ID); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return movies, nil
}
