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
		return nil, fmt.Errorf("query movies: %w", err)
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
		return nil, fmt.Errorf("search movies: %w", err)
	}
	defer rows.Close()
	return r.getMoviesFromRows(rows)
}

func (r *MovieRepository) GetByID(id int) (models.Movie, error) {
	var movie models.Movie
	var err error
	query := "SELECT * FROM movie WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
		return movie, fmt.Errorf("scan movie: %w", err)
	}

	movie.Actors, err = r.GetActors(movie.ID)
	if err != nil {
		return movie, fmt.Errorf("get movie actors: %w", err)
	}
	movie.Genres, err = r.GetGenres(movie.ID)
	if err != nil {
		return movie, fmt.Errorf("get movie genres: %w", err)
	}

	return movie, nil
}

func (r *MovieRepository) Create(req dto.CreateMovieRequest) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin create movie transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO movie (title, release_year, duration) VALUES (?, ?, ?)", req.Title, req.ReleaseYear, req.Duration)
	if err != nil {
		return 0, fmt.Errorf("insert movie: %w", err)
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get created movie id: %w", err)
	}
	movieID := int(id64)

	for _, genreID := range req.GenreIDs {
		if !r.GenreExists(genreID) {
			return 0, fmt.Errorf(IDNotFound, "genre", genreID)
		}

		if _, err := tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)", movieID, genreID); err != nil {
			return 0, fmt.Errorf("create movie-genre relationship: %w", err)
		}
	}

	for _, actorID := range req.ActorIDs {
		if !r.ActorExists(actorID) {
			return 0, fmt.Errorf(IDNotFound, "actor", actorID)
		}

		if _, err := tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)", movieID, actorID); err != nil {
			return 0, fmt.Errorf("create movie-actor relationship: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("commit movie creation: %w", err)
	}

	return movieID, nil
}

func (r *MovieRepository) Update(id int, req dto.UpdateMovieRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin update movie transaction: %w", err)
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
		return fmt.Errorf("update movie: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated movie rows: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	movieID := id

	for _, genreID := range req.GenreIDs {
		if !r.GenreExists(genreID) {
			return fmt.Errorf(IDNotFound, "genre", genreID)
		}

		if _, err := tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)", movieID, genreID); err != nil {
			return fmt.Errorf("create movie-genre relationship: %w", err)
		}
	}

	for _, actorID := range req.ActorIDs {
		if !r.ActorExists(actorID) {
			return fmt.Errorf(IDNotFound, "actor", actorID)
		}

		if _, err := tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)", movieID, actorID); err != nil {
			return fmt.Errorf("create movie-actor relationship: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit movie update: %w", err)
	}
	return nil
}

func (r *MovieRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM movie WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete movie: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted movie rows: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *MovieRepository) GetActors(movieID int) ([]models.Actor, error) {
	rows, err := r.db.Query("SELECT * FROM actor WHERE id in (SELECT actor_id FROM movie_actors WHERE movie_id = ?)", movieID)
	if err != nil {
		return nil, fmt.Errorf("query movie actors: %w", err)
	}
	defer rows.Close()

	actors := []models.Actor{}
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
			return nil, fmt.Errorf("scan movie actor row: %w", err)
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate movie actors: %w", err)
	}

	return actors, nil
}

func (r *MovieRepository) GetGenres(movieID int) ([]models.Genre, error) {
	rows, err := r.db.Query("SELECT * FROM genre WHERE id in (SELECT genre_id FROM movie_genres WHERE movie_id = ?)", movieID)
	if err != nil {
		return nil, fmt.Errorf("query movie genres: %w", err)
	}
	defer rows.Close()

	genres := []models.Genre{}
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, fmt.Errorf("scan movie genre row: %w", err)
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate movie genres: %w", err)
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
			return nil, fmt.Errorf("scan movie row: %w", err)
		}

		if movie.Actors, err = r.GetActors(movie.ID); err != nil {
			return nil, fmt.Errorf("get movie actors: %w", err)
		}
		if movie.Genres, err = r.GetGenres(movie.ID); err != nil {
			return nil, fmt.Errorf("get movie genres: %w", err)
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate movies: %w", err)
	}

	return movies, nil
}
