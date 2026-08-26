package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"movies-api/internal/dto"
	"movies-api/internal/models"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (r *ActorRepository) GetAll(page, size int, filters map[string]string) ([]models.Actor, error) {
	clauses, args := "", []any{}
	if size != 0 { // Checks whetever pagination exists or not, no other possibility of size being 0.
		clauses = " LIMIT ? OFFSET ?"
		args = []any{size, (page - 1) * size}
	}

	if value, ok := filters["name"]; ok {
		clauses += " AND name = ? COLLATE NOCASE"
		args = append(args, value)
	}
	if value, ok := filters["birth_date"]; ok {
		clauses += " AND birth_date = ?"
		args = append(args, value)
	}

	clauses = strings.Replace(clauses, "AND", "WHERE", 1)

	rows, err := r.db.Query("SELECT * FROM actor"+clauses, args...)

	if err != nil {
		return nil, fmt.Errorf("query actors: %w", err)
	}
	defer rows.Close()

	actors := []models.Actor{}
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
			return nil, fmt.Errorf("scan actor row: %w", err)
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate actors: %w", err)
	}

	return actors, nil
}

func (r *ActorRepository) GetByID(id int) (models.Actor, error) {
	var actor models.Actor
	query := "SELECT * FROM actor WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
		return actor, fmt.Errorf("scan actor: %w", err)
	}

	return actor, nil
}

func (r *ActorRepository) Create(req dto.CreateActorRequest) (int, error) {
	query := "INSERT INTO actor (name, birth_date) VALUES (?, ?)"
	result, err := r.db.Exec(query, req.Name, req.BirthDate)
	if err != nil {
		return 0, fmt.Errorf("insert actor: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get created actor id: %w", err)
	}

	return int(id), nil
}

func (r *ActorRepository) Update(id int, req dto.UpdateActorRequest) error {
	var sets []string
	var args []any

	if req.Name != nil {
		sets = append(sets, "name = ?")
		args = append(args, *req.Name)
	}

	if req.BirthDate != nil {
		sets = append(sets, "birth_date = ?")
		args = append(args, *req.BirthDate)
	}

	if len(sets) == 0 {
		return errors.New("no fields to update")
	}

	args = append(args, id)

	query := "UPDATE actor SET " + strings.Join(sets, ", ") + " WHERE id = ?"
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("update actor: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated actor rows: %w", err)
	} else if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ActorRepository) Delete(id int) error {
	query := "DELETE FROM actor WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete actor: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted actor rows: %w", err)
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ActorRepository) CreateRelationship(actorID, movieID int) error {
	if _, err := r.db.Exec("INSERT INTO movie_actors VALUES (?, ?)", movieID, actorID); err != nil {
		return fmt.Errorf("create movie-actor relationship: %w", err)
	}
	return nil
}

func (r *ActorRepository) DeleteRelationship(actorID, movieID int) error {
	if _, err := r.db.Exec("DELETE FROM movie_actors WHERE movie_id = ? AND actor_id = ?", movieID, actorID); err != nil {
		return fmt.Errorf("delete movie-actor relationship: %w", err)
	}
	return nil
}

func (r *ActorRepository) GetMovies(actorID int) ([]models.Movie, error) {
	rows, err := r.db.Query("SELECT * FROM movie WHERE id in (SELECT movie_id FROM movie_actors WHERE actor_id = ?)", actorID)
	if err != nil {
		return nil, fmt.Errorf("query actor movies: %w", err)
	}
	defer rows.Close()

	movies := []models.Movie{}
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseYear, &movie.Duration); err != nil {
			return nil, fmt.Errorf("scan actor movie row: %w", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate actor movies: %w", err)
	}

	return movies, nil
}
