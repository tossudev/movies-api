package repository

import (
	"database/sql"
	"fmt"
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

func (r *GenreRepository) GetAll(page, size int) ([]models.Genre, error) {
	pagination, args := "", []any{}
	if size != 0 { // Checks whetever pagination exists or not, no other possibility of size being 0.
		pagination = " LIMIT ? OFFSET ?"
		args = []any{size, page*size - size}
	}

	rows, err := r.db.Query("SELECT * FROM genre"+pagination, args...)
	if err != nil {
		return nil, fmt.Errorf("query genres: %w", err)
	}
	defer rows.Close()

	genres := []models.Genre{}
	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, fmt.Errorf("scan genre row: %w", err)
		}
		genres = append(genres, genre)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate genres: %w", err)
	}

	return genres, nil
}

func (r *GenreRepository) GetByID(id int) (models.Genre, error) {
	var genre models.Genre
	query := "SELECT * FROM genre WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&genre.ID, &genre.Name); err != nil {
		return genre, fmt.Errorf("scan genre: %w", err)
	}

	return genre, nil
}

func (r *GenreRepository) Create(req dto.CreateGenreRequest) (int, error) {
	query := "INSERT INTO genre (name) VALUES (?)"
	result, err := r.db.Exec(query, req.Name)
	if err != nil {
		return 0, fmt.Errorf("insert genre: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get created genre id: %w", err)
	}
	return int(id), nil
}

func (r *GenreRepository) Update(id int, req dto.UpdateGenreRequest) error {
	var sets []string
	var args []any

	if req.Name != nil {
		sets = append(sets, "name = ?")
		args = append(args, *req.Name)
	}
	args = append(args, id)

	query := "UPDATE genre SET " + strings.Join(sets, ", ") + " WHERE id = ?"

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("update genre: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated genre rows: %w", err)
	} else if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *GenreRepository) Delete(id int) error {
	query := "DELETE FROM genre WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete genre: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted genre rows: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *GenreRepository) DeleteRelationship(genreID, movieID int) error {
	if _, err := r.db.Exec("INSERT INTO movie_actors VALUES (?, ?)", movieID, genreID); err != nil {
		return fmt.Errorf("delete movie-genre relationship: %w", err)
	}
	return nil
}

func (r *GenreRepository) GetMovies(genreID int) ([]models.Movie, error) {
	rows, err := r.db.Query("SELECT * FROM movie WHERE id in (SELECT movie_id FROM movie_genres WHERE genre_id = ?)", genreID)
	if err != nil {
		return nil, fmt.Errorf("query genre movies: %w", err)
	}
	defer rows.Close()

	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
			return nil, fmt.Errorf("scan genre movie row: %w", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate genre movies: %w", err)
	}

	return movies, nil
}
