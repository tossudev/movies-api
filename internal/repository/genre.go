package repository

import (
	"database/sql"
	"strings"

	"movies-api/internal/dto"
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

func (r *GenreRepository) Create(req dto.CreateGenreRequest) (int, error) {
	query := "INSERT INTO genre (name) VALUES (?)"
	result, err := r.db.Exec(query, req.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *GenreRepository) Update(id int, req dto.UpdateGenreRequest) error {
	var sets []string
	var args []any

	if req.Name != nil {
		sets = append(sets, "name = ?")
		args = append(args, *req.Name)
	}

	if len(sets) == 0 {
		return nil // or return an error
	}

	args = append(args, id)

	query := "UPDATE genre SET " +
		strings.Join(sets, ", ") +
		" WHERE id = ?"

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *GenreRepository) Delete(id int) error {
	query := "DELETE FROM genre WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
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
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
			return nil, err // TODO: concretize error
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err // TODO: concretize error
	}

	return movies, nil
}
