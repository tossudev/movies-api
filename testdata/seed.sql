/*
Editor's note:
This is all AI-generated and uses IDMB top 20 movies (supposedly)
*/

-- Sample movie database data
-- 20 highly rated films appearing in IMDb's Top 250.
-- duration is stored as minutes.
-- birth_date uses ISO 8601 YYYY-MM-DD.
-- IDs are explicit so the relationship rows are deterministic.

-- Optional cleanup for repeatable testing:
DELETE FROM movie_actors;
DELETE FROM movie_genres;
DELETE FROM actor;
DELETE FROM movie;
DELETE FROM genre;

INSERT INTO genre (id, name) VALUES
(1, 'Drama'),
(2, 'Crime'),
(3, 'Action'),
(4, 'Adventure'),
(5, 'Fantasy'),
(6, 'Sci-Fi'),
(7, 'Thriller'),
(8, 'Biography'),
(9, 'History'),
(10, 'War'),
(11, 'Western'),
(12, 'Mystery');

INSERT INTO movie (id, title, release_year, duration) VALUES
(1, 'The Shawshank Redemption', '1994', '142'),
(2, 'The Godfather', '1972', '175'),
(3, 'The Dark Knight', '2008', '152'),
(4, 'The Godfather Part II', '1974', '202'),
(5, 'The Lord of the Rings: The Return of the King', '2003', '201'),
(6, '12 Angry Men', '1957', '96'),
(7, 'Schindler''s List', '1993', '195'),
(8, 'The Lord of the Rings: The Fellowship of the Ring', '2001', '178'),
(9, 'Pulp Fiction', '1994', '154'),
(10, 'The Good, the Bad and the Ugly', '1966', '178'),
(11, 'The Lord of the Rings: The Two Towers', '2002', '179'),
(12, 'Forrest Gump', '1994', '142'),
(13, 'Fight Club', '1999', '139'),
(14, 'Inception', '2010', '148'),
(15, 'Star Wars: Episode V - The Empire Strikes Back', '1980', '124'),
(16, 'The Matrix', '1999', '136'),
(17, 'Goodfellas', '1990', '145'),
(18, 'Interstellar', '2014', '169'),
(19, 'One Flew Over the Cuckoo''s Nest', '1975', '133'),
(20, 'Seven', '1995', '127');

INSERT INTO actor (id, name, birth_date) VALUES
(1, 'Tim Robbins', '1958-10-16'),
(2, 'Morgan Freeman', '1937-06-01'),
(3, 'Marlon Brando', '1924-04-03'),
(4, 'Al Pacino', '1940-04-25'),
(5, 'James Caan', '1940-03-26'),
(6, 'Christian Bale', '1974-01-30'),
(7, 'Heath Ledger', '1979-04-04'),
(8, 'Aaron Eckhart', '1968-03-12'),
(9, 'Robert De Niro', '1943-08-17'),
(10, 'Robert Duvall', '1931-01-05'),
(11, 'Elijah Wood', '1981-01-28'),
(12, 'Viggo Mortensen', '1958-10-20'),
(13, 'Ian McKellen', '1939-05-25'),
(14, 'Henry Fonda', '1905-05-16'),
(15, 'Lee J. Cobb', '1911-12-08'),
(16, 'Martin Balsam', '1919-11-04'),
(17, 'Liam Neeson', '1952-06-07'),
(18, 'Ben Kingsley', '1943-12-31'),
(19, 'Ralph Fiennes', '1962-12-22'),
(20, 'Sean Astin', '1971-02-25'),
(21, 'Uma Thurman', '1970-04-29'),
(22, 'John Travolta', '1954-02-18'),
(23, 'Samuel L. Jackson', '1948-12-21'),
(24, 'Clint Eastwood', '1930-05-31'),
(25, 'Eli Wallach', '1915-12-07'),
(26, 'Lee Van Cleef', '1925-01-09'),
(27, 'Andy Serkis', '1964-04-20'),
(28, 'Tom Hanks', '1956-07-09'),
(29, 'Robin Wright', '1966-04-08'),
(30, 'Gary Sinise', '1955-03-17'),
(31, 'Brad Pitt', '1963-12-18'),
(32, 'Edward Norton', '1969-08-18'),
(33, 'Helena Bonham Carter', '1966-05-26'),
(34, 'Leonardo DiCaprio', '1974-11-11'),
(35, 'Joseph Gordon-Levitt', '1981-02-17'),
(36, 'Elliot Page', '1987-02-21'),
(37, 'Tom Hardy', '1977-09-15'),
(38, 'Mark Hamill', '1951-09-25'),
(39, 'Harrison Ford', '1942-07-13'),
(40, 'Carrie Fisher', '1956-10-21'),
(41, 'Keanu Reeves', '1964-09-02'),
(42, 'Laurence Fishburne', '1961-07-30'),
(43, 'Carrie-Anne Moss', '1967-08-21'),
(44, 'Ray Liotta', '1954-12-18'),
(45, 'Joe Pesci', '1943-02-09'),
(46, 'Matthew McConaughey', '1969-11-04'),
(47, 'Anne Hathaway', '1982-11-12'),
(48, 'Jessica Chastain', '1977-03-24'),
(49, 'Jack Nicholson', '1937-04-22'),
(50, 'Louise Fletcher', '1934-07-22'),
(51, 'Will Sampson', '1933-09-27'),
(52, 'Kevin Spacey', '1959-07-26');

INSERT INTO movie_genres (movie_id, genre_id) VALUES
(1,1),
(2,1),(2,2),
(3,1),(3,2),(3,3),(3,7),
(4,1),(4,2),
(5,1),(5,3),(5,4),(5,5),
(6,1),
(7,1),(7,8),(7,9),
(8,1),(8,3),(8,4),(8,5),
(9,1),(9,2),
(10,4),(10,11),
(11,1),(11,3),(11,4),(11,5),
(12,1),
(13,1),
(14,3),(14,4),(14,6),(14,7),
(15,3),(15,4),(15,5),(15,6),
(16,3),(16,6),
(17,1),(17,2),(17,8),
(18,1),(18,4),(18,6),
(19,1),
(20,1),(20,2),(20,12),(20,7);

INSERT INTO movie_actors (movie_id, actor_id) VALUES
(1,1),(1,2),
(2,3),(2,4),(2,5),(2,10),
(3,6),(3,7),(3,8),(3,2),
(4,4),(4,9),(4,10),
(5,11),(5,12),(5,13),(5,20),(5,27),
(6,14),(6,15),(6,16),
(7,17),(7,18),(7,19),
(8,11),(8,12),(8,13),(8,20),
(9,21),(9,22),(9,23),
(10,24),(10,25),(10,26),
(11,11),(11,12),(11,13),(11,20),(11,27),
(12,28),(12,29),(12,30),
(13,31),(13,32),(13,33),
(14,34),(14,35),(14,36),(14,37),
(15,38),(15,39),(15,40),
(16,41),(16,42),(16,43),
(17,9),(17,44),(17,45),
(18,46),(18,47),(18,48),
(19,49),(19,50),(19,51),
(20,2),(20,31),(20,52);
