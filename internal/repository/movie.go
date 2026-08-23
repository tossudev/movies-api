package repository

import (
	"database/sql"

	"movies-api/internal/dto"
	"movies-api/internal/models"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetAll() ([]models.Movie, error) {
	rows, err := r.db.Query("SELECT * FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
			return nil, err // TODO: concretize error
		}

		if movie.Actors, err = r.getActors(movie.ID); err != nil {
			return nil, err
		}
		if movie.Genres, err = r.getGenres(movie.ID); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return movies, nil
}

func (r *MovieRepository) Exists(id int) bool {
	_, err := r.GetByID(id)
	return err != sql.ErrNoRows
}

func (r *MovieRepository) GetByID(id int) (models.Movie, error) {
	var movie models.Movie
	query := "SELECT * FROM movie WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
		return movie, err // TODO: concretize error
	}

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
		if _, err := tx.Exec("INSERT INTO movie_genres (movie_id, genre_id) VALUES (?, ?)", movieID, genreID); err != nil {
			return 0, err
		}
	}

	for _, actorID := range req.ActorIDs {
		if _, err := tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES (?, ?)", movieID, actorID); err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return movieID, nil
}

func (r *MovieRepository) Update(movie models.Movie) error {
	query := "UPDATE movie SET title = ?, release_year = ?, duration = ? WHERE id = ?"
	_, err := r.db.Exec(query, movie.Title, movie.ReleaseYear, movie.Duration, movie.ID)

	return err
}

func (r *MovieRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM movie WHERE id = ?", id)
	return err
}

func (r *MovieRepository) getActors(movieID int) ([]models.Actor, error) {
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

func (r *MovieRepository) getGenres(movieID int) ([]models.Genre, error) {
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
