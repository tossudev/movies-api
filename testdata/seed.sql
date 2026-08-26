/*
Editor's note:
This is all AI-generated and uses IDMB top 250 movies (supposedly)

Expected:
8 genres, with every genre linked to at least 2 movies
24 movies
45 actors
Release years 1954–2014
Movies with 1–3 genres
Movies with 1–5 actors
11 actors appear in multiple movies
Multiple shared casts, including the LOTR trilogy, Godfather films, etc.
Foreign/older titles are included to make the date and relationship ranges more useful for testing
*/

-- Sample seed data based on representative movies from IMDb's current Top 250.
-- IMDb Top 250: https://www.imdb.com/chart/top/
-- Accessed for verification: 2026-08-26

BEGIN TRANSACTION;

INSERT INTO genre (id, name) VALUES
(1, 'Action'),
(2, 'Adventure'),
(3, 'Biography'),
(4, 'Crime'),
(5, 'Drama'),
(6, 'Fantasy'),
(7, 'Sci-Fi'),
(8, 'Thriller');

INSERT INTO movie (id, title, release_year, duration) VALUES
(1, 'The Shawshank Redemption', 1994, 142),
(2, 'The Godfather', 1972, 175),
(3, 'The Dark Knight', 2008, 152),
(4, 'The Godfather Part II', 1974, 202),
(5, 'The Lord of the Rings: The Return of the King', 2003, 201),
(6, '12 Angry Men', 1957, 96),
(7, 'Schindler''s List', 1993, 195),
(8, 'The Lord of the Rings: The Fellowship of the Ring', 2001, 178),
(9, 'Pulp Fiction', 1994, 154),
(10, 'The Good, the Bad and the Ugly', 1966, 178),
(11, 'The Lord of the Rings: The Two Towers', 2002, 179),
(12, 'Forrest Gump', 1994, 142),
(13, 'Fight Club', 1999, 139),
(14, 'Inception', 2010, 148),
(15, 'The Empire Strikes Back', 1980, 124),
(16, 'The Matrix', 1999, 136),
(17, 'Interstellar', 2014, 169),
(18, 'Goodfellas', 1990, 145),
(19, 'One Flew Over the Cuckoo''s Nest', 1975, 133),
(20, 'Seven Samurai', 1954, 207),
(21, 'Se7en', 1995, 127),
(22, 'The Silence of the Lambs', 1991, 118),
(23, 'City of God', 2002, 130),
(24, 'Saving Private Ryan', 1998, 169);

INSERT INTO actor (id, name, birth_date) VALUES
(1, 'Tim Robbins', '1958-10-16'),
(2, 'Morgan Freeman', '1937-06-01'),
(3, 'Marlon Brando', '1924-04-03'),
(4, 'Al Pacino', '1940-04-25'),
(5, 'James Caan', '1940-03-26'),
(6, 'Christian Bale', '1974-01-30'),
(7, 'Heath Ledger', '1979-04-04'),
(8, 'Aaron Eckhart', '1968-03-12'),
(9, 'Elijah Wood', '1981-01-17'),
(10, 'Viggo Mortensen', '1958-10-20'),
(11, 'Ian McKellen', '1939-05-25'),
(12, 'Sean Astin', '1971-02-25'),
(13, 'Uma Thurman', '1970-04-29'),
(14, 'Samuel L. Jackson', '1948-12-21'),
(15, 'John Travolta', '1954-02-18'),
(16, 'Clint Eastwood', '1930-05-31'),
(17, 'Eli Wallach', '1915-12-07'),
(18, 'Lee Van Cleef', '1925-01-09'),
(19, 'Tom Hanks', '1956-07-09'),
(20, 'Gary Sinise', '1955-03-17'),
(21, 'Brad Pitt', '1963-12-18'),
(22, 'Edward Norton', '1969-08-18'),
(23, 'Leonardo DiCaprio', '1974-11-11'),
(24, 'Joseph Gordon-Levitt', '1981-02-17'),
(25, 'Tom Hardy', '1977-09-15'),
(26, 'Ken Watanabe', '1959-10-21'),
(27, 'Carrie Fisher', '1956-10-21'),
(28, 'Mark Hamill', '1951-09-25'),
(29, 'Harrison Ford', '1942-07-13'),
(30, 'Keanu Reeves', '1964-09-02'),
(31, 'Laurence Fishburne', '1961-07-30'),
(32, 'Matthew McConaughey', '1969-11-04'),
(33, 'Robert De Niro', '1943-08-17'),
(34, 'Ray Liotta', '1954-12-18'),
(35, 'Joe Pesci', '1943-02-09'),
(36, 'Jodie Foster', '1962-11-19'),
(37, 'Anthony Hopkins', '1937-12-31'),
(38, 'Matt Damon', '1970-10-08'),
(39, 'Vincent D''Onofrio', '1959-06-30'),
(40, 'Martin Balsam', '1919-11-04'),
(41, 'Jack Nicholson', '1937-04-22'),
(42, 'Takashi Shimura', '1905-03-12'),
(43, 'Liam Neeson', '1952-06-07'),
(44, 'Toshiro Mifune', '1920-04-01'),
(45, 'Alexandre Rodrigues', '1983-05-21');

INSERT INTO movie_genres (movie_id, genre_id) VALUES
-- Single genre
(1, 5),
(13, 5),
(19, 5),
-- Multiple genres
(2, 4), (2, 5),
(3, 1), (3, 4), (3, 5),
(4, 4), (4, 5),
(5, 2), (5, 5), (5, 6),
(6, 4), (6, 5),
(7, 3), (7, 5),
(8, 2), (8, 5), (8, 6),
(9, 4), (9, 5),
(10, 1), (10, 5),
(11, 2), (11, 5), (11, 6),
(12, 5),
(14, 1), (14, 2), (14, 7),
(15, 1), (15, 2), (15, 6),
(16, 1), (16, 7),
(17, 2), (17, 5), (17, 7),
(18, 3), (18, 4), (18, 5),
(20, 1), (20, 5),
(21, 4), (21, 5), (21, 8),
(22, 4), (22, 5), (22, 8),
(23, 4), (23, 5),
(24, 5), (24, 8);

INSERT INTO movie_actors (movie_id, actor_id) VALUES
-- The Shawshank Redemption
(1, 1), (1, 2),
-- The Godfather
(2, 3), (2, 4), (2, 5),
-- The Dark Knight
(3, 6), (3, 7), (3, 8),
-- The Godfather Part II
(4, 4), (4, 33),
-- The Lord of the Rings: The Return of the King
(5, 9), (5, 10), (5, 11), (5, 12),
-- 12 Angry Men
(6, 40),
-- Schindler's List
(7, 43),
-- The Lord of the Rings: The Fellowship of the Ring
(8, 9), (8, 10), (8, 11), (8, 12),
-- Pulp Fiction
(9, 13), (9, 14), (9, 15),
-- The Good, the Bad and the Ugly
(10, 16), (10, 17), (10, 18),
-- The Lord of the Rings: The Two Towers
(11, 9), (11, 10), (11, 11), (11, 12),
-- Forrest Gump
(12, 19), (12, 20),
-- Fight Club
(13, 21), (13, 22),
-- Inception
(14, 23), (14, 24), (14, 25), (14, 26), (14, 38),
-- The Empire Strikes Back
(15, 27), (15, 28), (15, 29),
-- The Matrix
(16, 30), (16, 31),
-- Interstellar
(17, 32), (17, 38),
-- Goodfellas
(18, 33), (18, 34), (18, 35),
-- One Flew Over the Cuckoo's Nest
(19, 41),
-- Seven Samurai
(20, 42), (20, 44),
-- Se7en
(21, 21), (21, 22),
-- The Silence of the Lambs
(22, 36), (22, 37),
-- City of God
(23, 45),
-- Saving Private Ryan
(24, 19), (24, 20);

COMMIT;
