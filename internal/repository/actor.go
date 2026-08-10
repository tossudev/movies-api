package repository

import (
	"database/sql"
	"movies-api/internal/models"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (r *ActorRepository) GetAll() ([]models.Actor, error) {
	rows, err := r.db.Query("SELECT * FROM actor")
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

func (r *ActorRepository) GetByID(id int) (models.Actor, error) {
	var actor models.Actor
	// TODO: should use indexed variables or '?'
	query := "SELECT * FROM actor WHERE id = ?;"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&actor.ID, &actor.Name, &actor.BirthDate); err != nil {
		return actor, err // TODO: concretize error
	}

	return actor, nil
}

func (r *ActorRepository) Create(actor models.Actor) error {
	query := "INSERT INTO actor (name, birth_date) VALUES (?, ?)"
	_, err := r.db.Exec(query, actor.Name, actor.BirthDate)

	return err
}

func (r *ActorRepository) Update(actor models.Actor) error {
	query := "UPDATE actor SET name = ?, birth_date = ? WHERE id = ?"
	_, err := r.db.Exec(query, actor.Name, actor.BirthDate, actor.ID)

	return err
}

func (r *ActorRepository) Delete(id int) error {
	query := "DELETE FROM actor WHERE id = ?"
	_, err := r.db.Exec(query, id)

	return err
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
