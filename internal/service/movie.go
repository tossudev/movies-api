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

func (s *MovieService) GetAll() ([]models.Movie, error) {
	return s.repo.GetAll()
}

func (s *MovieService) Search(query string) ([]models.Movie, error) {
	if len(query) == 0 {
		return nil, errors.New("Empty search")
	}
	return s.repo.Search(query)
}

func (s *MovieService) GetByID(id int) (models.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *MovieService) Create(req dto.CreateMovieRequest) (models.Movie, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{ID: id, Title: req.Title, ReleaseYear: req.ReleaseYear, Duration: req.Duration}, nil
}

func (s *MovieService) Update(id int, req dto.UpdateMovieRequest) error {
	return s.repo.Update(id, req)
}

func (s *MovieService) Delete(id int, force bool) error {
	if !force {
		if _, err := s.repo.GetActors(id); err == nil {
			return errors.New("Couldn't delete movie")
		}
		if _, err := s.repo.GetGenres(id); err == nil {
			return errors.New("Couldn't delete movie")
		}
	}

	return s.repo.Delete(id)
}
