/*
Editor's note:
This is all AI-generated and uses IDMB top 20 movies (supposedly)
*/

-- SQLite movie DB: 5 genres, ONLY 20 movies, 15 actors
-- Many-to-Many:
--   Genre <-> Movie via MovieGenre
--   Actor <-> Movie via MovieActor

PRAGMA foreign_keys = ON;

BEGIN;

-- 5 Genres
INSERT INTO genre (id, name) VALUES
  (1, 'Drama'),
  (2, 'Crime'),
  (3, 'Action'),
  (4, 'Sci-Fi'),
  (5, 'Adventure');

-- 20 Movies (IMDb Top 250 entries; restricted to 20 titles total in this DB)
-- (Durations are minutes; adjust as desired.)
INSERT INTO movie (id, title, release_year, duration) VALUES
  (1,  'The Shawshank Redemption', 1994, 142),
  (2,  'The Godfather', 1972, 175),
  (3,  'The Dark Knight', 2008, 152),
  (4,  'The Godfather Part II', 1974, 202),
  (5,  '12 Angry Men', 1957, 96),
  (6,  'Schindler''s List', 1993, 195),
  (7,  'The Lord of the Rings: The Return of the King', 2003, 201),
  (8,  'Pulp Fiction', 1994, 154),
  (9,  'The Lord of the Rings: The Fellowship of the Ring', 2001, 178),
  (10, 'The Good, the Bad and the Ugly', 1966, 161),
  (11, 'Fight Club', 1999, 139),
  (12, 'Inception', 2010, 148),
  (13, 'The Lord of the Rings: The Two Towers', 2002, 179),
  (14, 'Forrest Gump', 1994, 142),
  (15, 'The Empire Strikes Back', 1980, 124),
  (16, 'The Matrix', 1999, 136),
  (17, 'The Silence of the Lambs', 1991, 118),
  (18, 'Se7en', 1995, 127),
  (19, 'The Prestige', 2006, 130),
  (20, 'Saving Private Ryan', 1998, 169);

-- Genre <-> Movie (many-to-many)
INSERT INTO movie_genres (movie_id, genre_id) VALUES
  -- Drama
  (1, 1),(5, 1),(6, 1),(8, 1),(11, 1),(12, 1),(14, 1),(17, 1),(20, 1),
  -- Crime
  (2, 2),(4, 2),(8, 2),(10, 2),(15, 2),(17, 2),(18, 2),
  -- Action
  (3, 3),(7, 3),(9, 3),(11, 3),(12, 3),(18, 3),(20, 3),
  -- Sci-Fi
  (3, 4),(12, 4),(15, 4),(16, 4),(19, 4),
  -- Adventure
  (7, 5),(9, 5),(13, 5),(10, 5),(20, 5);

-- 15 Actors
INSERT INTO actor (id, name, birth_date) VALUES
  (1,  'Tim Robbins', '1958-10-16'),
  (2,  'Morgan Freeman', '1937-06-01'),
  (3,  'Marlon Brando', '1924-04-03'),
  (4,  'Al Pacino', '1940-04-25'),
  (5,  'James Caan', '1940-03-26'),
  (6,  'Christian Bale', '1974-01-30'),
  (7,  'Heath Ledger', '1979-04-04'),
  (8,  'Robert De Niro', '1943-08-17'),
  (9,  'Joe Pesci', '1943-02-09'),
  (10, 'Leonardo DiCaprio', '1974-11-11'),
  (11, 'Joseph Gordon-Levitt', '1981-02-17'),
  (12, 'Keanu Reeves', '1964-09-02'),
  (13, 'Carrie-Anne Moss', '1967-08-21'),
  (14, 'Anthony Hopkins', '1937-12-31'),
  (15, 'Samuel L. Jackson', '1948-12-21');

-- Actor <-> Movie (many-to-many)
INSERT INTO movie_actors (movie_id, actor_id) VALUES
  -- The Shawshank Redemption
  (1, 1),
  (1, 2),

  -- The Godfather
  (2, 3),
  (2, 4),
  (2, 5),

  -- The Dark Knight
  (3, 6),
  (3, 7),

  -- The Godfather Part II
  (4, 8),
  (4, 4),
  (4, 3),

  -- 12 Angry Men
  (5, 2),

  -- Schindler's List
  (6, 2),
  (6, 1),

  -- LOTR: Return of the King
  (7, 2),

  -- Pulp Fiction
  (8, 15),
  (8, 2),

  -- LOTR: Fellowship of the Ring
  (9, 2),

  -- The Good, the Bad and the Ugly
  (10, 2),

  -- Fight Club
  (11, 6),
  (11, 10),

  -- Inception
  (12, 10),
  (12, 11),
  (12, 2),

  -- LOTR: Two Towers
  (13, 2),

  -- Forrest Gump
  (14, 2),

  -- Empire Strikes Back
  (15, 12),

  -- The Matrix
  (16, 12),
  (16, 13),
  (16, 2),

  -- Silence of the Lambs
  (17, 14),
  (17, 1),

  -- Se7en
  (18, 4),
  (18, 2),

  -- The Prestige
  (19, 10),
  (19, 6),

  -- Saving Private Ryan
  (20, 1),
  (20, 8);

COMMIT;

