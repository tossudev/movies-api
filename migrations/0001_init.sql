
/*
Genre (fields: id, name)
Movie (fields: id, title, release_year, duration)
Actor (fields: id, name, birth_date)
*/

CREATE TABLE genre (
	id INT PRIMARY KEY	NOT NULL,
	name TEXT							NOT NULL
);

CREATE TABLE movie (
	id INT PRIMARY KEY	NOT NULL,
	title TEXT							NOT NULL,
	release_year TEXT					NOT NULL,
	duration TEXT						NOT NULL
);

CREATE TABLE actor (
	id INT PRIMARY KEY	NOT NULL,
	name TEXT							NOT NULL,
	birth_date TEXT						NOT NULL
);

CREATE TABLE movie_genres (
	movie_id INT NOT NULL,
	genre_id INT NOT NULL,
	PRIMARY KEY (movie_id, genre_id)
);

CREATE TABLE movie_actors (
	movie_id INT NOT NULL,
	actor_id INT NOT NULL,
	PRIMARY KEY (movie_id, actor_id)
);
