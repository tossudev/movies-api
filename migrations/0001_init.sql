
/*
Genre (fields: id, name)
Movie (fields: id, title, release_year, duration)
Actor (fields: id, name, birth_date)
*/

CREATE TABLE IF NOT EXISTS genre (
	id INT PRIMARY KEY	NOT NULL,
	name TEXT							NOT NULL
);

CREATE TABLE IF NOT EXISTS movie (
	id INT PRIMARY KEY	NOT NULL,
	title TEXT							NOT NULL,
	release_year INT					NOT NULL,
	duration INT        NOT NULL
);

CREATE TABLE IF NOT EXISTS actor (
	id INT PRIMARY KEY	NOT NULL,
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
