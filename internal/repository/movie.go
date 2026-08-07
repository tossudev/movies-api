package repository

import (
	"database/sql"
	"movies-api/internal/models"
)

func GetAllMovies(db *sql.DB) ([]models.Movie, error) {
	rows, err := db.Query("SELECT * FROM movie")
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
		return movies, err
	}

	return movies, nil
}
