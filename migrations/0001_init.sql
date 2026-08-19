
/*
Genre (fields: id, name)
Movie (fields: id, title, release_year, duration)
Actor (fields: id, name, birth_date)
*/

CREATE TABLE IF NOT EXISTS genre (
	id INTEGER PRIMARY KEY,
	name TEXT							NOT NULL
);

CREATE TABLE IF NOT EXISTS movie (
	id INTEGER PRIMARY KEY,
	title TEXT							NOT NULL,
	release_year INT					NOT NULL,
	duration INT        NOT NULL
);

CREATE TABLE IF NOT EXISTS actor (
	id INTEGER PRIMARY KEY,
	name TEXT							NOT NULL,
	birth_date TEXT						NOT NULL
);

CREATE TABLE IF NOT EXISTS movie_actors (
    movie_id INTEGER NOT NULL,
    actor_id INTEGER NOT NULL,
    PRIMARY KEY (movie_id, actor_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    FOREIGN KEY (actor_id) REFERENCES actor(id) ON DELETE CASCADE
  );

  CREATE TABLE IF NOT EXISTS movie_genres (
    movie_id INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    PRIMARY KEY (movie_id, genre_id),
    FOREIGN KEY (movie_id) REFERENCES movie(id) ON DELETE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES genre(id) ON DELETE CASCADE
  );
