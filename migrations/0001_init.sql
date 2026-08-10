
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

CREATE TABLE IF NOT EXISTS movie_genres (
	movie_id INT NOT NULL,
	genre_id INT NOT NULL,
	PRIMARY KEY (movie_id, genre_id)
);

CREATE TABLE IF NOT EXISTS movie_actors (
	movie_id INT NOT NULL,
	actor_id INT NOT NULL,
	PRIMARY KEY (movie_id, actor_id)
);
