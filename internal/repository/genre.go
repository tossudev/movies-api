package repository

import (
	"database/sql"

	"movies-api/internal/models"
)

type GenreRepository struct {
	db *sql.DB
}

func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{db: db}
}

func (r *GenreRepository) GetAll() ([]models.Genre, error) {
	rows, err := r.db.Query("SELECT * FROM genre")
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

func (r *GenreRepository) Exists(id int) bool {
	_, err := r.GetByID(id)
	return err != sql.ErrNoRows
}

func (r *GenreRepository) GetByID(id int) (models.Genre, error) {
	var genre models.Genre
	query := "SELECT * FROM genre WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&genre.ID, &genre.Name); err != nil {
		return genre, err // TODO: concretize error
	}

	return genre, nil
}

func (r *GenreRepository) Create(genre models.Genre) error {
	query := "INSERT INTO genre (name) VALUES (?)"
	_, err := r.db.Exec(query, genre.Name)

	return err
}

func (r *GenreRepository) Update(genre models.Genre) error {
	query := "UPDATE genre SET name = ? WHERE id = ?"
	_, err := r.db.Exec(query, genre.Name, genre.ID)

	return err
}

func (r *GenreRepository) Delete(id int) error {
	query := "DELETE FROM genre WHERE id = ?"
	_, err := r.db.Exec(query, id)

	return err
}

func (r *GenreRepository) CreateRelationship(genreID, movieID int) error {
	query := "INSERT INTO movie_genres VALUES (?, ?)"
	_, err := r.db.Exec(query, movieID, genreID)

	return err
}

func (r *GenreRepository) DeleteRelationship(genreID, movieID int) error {
	query := "DELETE FROM movie_genres WHERE movie_id = ? AND genre_id = ?"
	_, err := r.db.Exec(query, movieID, genreID)

	return err
}

func (r *GenreRepository) GetMovies(genreID int) ([]models.Movie, error) {
	rows, err := r.db.Query("SELECT * FROM movie WHERE id in (SELECT movie_id FROM movie_genres WHERE genre_id = ?)", genreID)
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
