package repository

import (
	"database/sql"
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

func (r *ActorRepository) GetAll(page, size int) ([]models.Actor, error) {
	pagination, args := "", []any{}
	if size != 0 { // Checks whetever pagination exists or not, no other possibility of size being 0.
		pagination = " LIMIT ? OFFSET ?"
		args = []any{size, page*size - size}
	}

	rows, err := r.db.Query("SELECT * FROM actor"+pagination, args...)

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

func (r *ActorRepository) Exists(id int) bool {
	_, err := r.GetByID(id)
	return err != sql.ErrNoRows
}

func (r *ActorRepository) GetByID(id int) (models.Actor, error) {
	var actor models.Actor
	query := "SELECT * FROM actor WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
		return actor, err // TODO: concretize error
	}

	return actor, nil
}

func (r *ActorRepository) Create(req dto.CreateActorRequest) (int, error) {
	query := "INSERT INTO actor (name, birth_date) VALUES (?, ?)"
	result, err := r.db.Exec(query, req.Name, req.BirthDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
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
		return nil // or return an error
	}

	args = append(args, id)

	query := "UPDATE actor SET " + strings.Join(sets, ", ") + " WHERE id = ?"
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *ActorRepository) Delete(id int) error {
	query := "DELETE FROM actor WHERE id = ?"
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

func (r *ActorRepository) CreateRelationship(actorID, movieID int) error {
	query := "INSERT INTO movie_actors VALUES (?, ?)"
	_, err := r.db.Exec(query, movieID, actorID)

	return err
}

func (r *ActorRepository) DeleteRelationship(actorID, movieID int) error {
	query := "DELETE FROM movie_actors WHERE movie_id = ? AND actor_id = ?"
	_, err := r.db.Exec(query, movieID, actorID)

	return err
}

func (r *ActorRepository) GetMovies(actorID int) ([]models.Movie, error) {
	rows, err := r.db.Query("SELECT * FROM movie WHERE id in (SELECT movie_id FROM movie_actors WHERE actor_id = ?)", actorID)
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
