package service

import (
	"errors"

	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

type MovieService struct {
	repo *repository.MovieRepository
}

func NewMovieService(repo *repository.MovieRepository) *MovieService {
	return &MovieService{
		repo: repo,
	}
}

func (s *MovieService) GetAll(page, size int, filters map[string]string) ([]models.Movie, error) {
	return s.repo.GetAll(page, size, filters)
}

func (s *MovieService) Search(query string, page, size int) ([]models.Movie, error) {
	if len(query) == 0 {
		return nil, errors.New("Empty search")
	}
	return s.repo.Search(query, page, size)
}

func (s *MovieService) GetByID(id int) (models.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *MovieService) GetActors(id int) ([]models.Actor, error) {
	return s.repo.GetActors(id)
}

func (s *MovieService) Create(req dto.CreateMovieRequest) (models.Movie, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return models.Movie{}, err
	}

	genres, err := s.repo.GetGenres(id)
	if err != nil {
		return models.Movie{}, err
	}
	actors, err := s.repo.GetActors(id)
	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{ID: id, Title: req.Title, ReleaseYear: req.ReleaseYear, Duration: req.Duration, Genres: genres, Actors: actors}, nil
}

func (s *MovieService) Update(id int, req dto.UpdateMovieRequest) error {
	return s.repo.Update(id, req)
}

func (s *MovieService) Delete(id int, force bool) error {
	return s.repo.Delete(id)
}
