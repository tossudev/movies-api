package repository

import (
	"database/sql"

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
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Releaseyear, &movie.Duration); err != nil {
			return nil, err // TODO: concretize error
		}

		if movie.Actors, err = r.getActors(movie.ID); err != nil {
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
