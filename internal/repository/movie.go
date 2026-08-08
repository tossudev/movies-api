package repository

import (
	"database/sql"
	"movies-api/internal/models"
)

type MovieRepository struct { // In go, there's no possibility to creating a method to other
	db *sql.DB // package's structs. Some handshaker must be created :(
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
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// func (r *MovieRepository) GetByID(id int) (models.Movie, error) {
// 	// ...
// }

// func (r *MovieRepository) Create(movie models.Movie) error {
// 	// ...
// }

// func (r *MovieRepository) Update(movie models.Movie) error {
// 	// ...
// }

// func (r *MovieRepository) Delete(id int) error {
// 	// ...
// }
