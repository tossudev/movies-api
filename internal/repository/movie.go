package repository

import (
	"database/sql"

	"movies-api/internal/models"
)

// In go, there's no possibility to creating a method to other package's structs.
// Some handshaker must be created :(
type MovieRepository struct {
	db *sql.DB
}

var Movie *MovieRepository

func InitMovieRepository(db *sql.DB) {
	Movie = &MovieRepository{db: db}
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
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Releaseyear, &movie.Duration); err != nil {
			return nil, err // TODO: concretize error
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return movies, nil
}

func (r *MovieRepository) GetByID(id int) (models.Movie, error) {
	var movie models.Movie
	query := "SELECT * FROM movie WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&movie.ID, &movie.Title, &movie.Releaseyear, &movie.Duration); err != nil {
		return movie, err // TODO: concretize error
	}

	return movie, nil
}

func (r *MovieRepository) Create(movie models.Movie) error {
	query := "INSERT INTO movie (title, release_year, duration) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, movie.Title, movie.Releaseyear, movie.Duration)

	return err
}

func (r *MovieRepository) Update(movie models.Movie) error {
	query := "UPDATE movie SET title = ?, release_year = ?, duration = ? WHERE id = ?"
	_, err := r.db.Exec(query, movie.Title, movie.Releaseyear, movie.Duration, movie.ID)

	return err
}

func (r *MovieRepository) Delete(id int) error {
	query := "DELETE FROM movie WHERE id = ?"
	_, err := r.db.Exec(query, id)

	return err
}
