package service

import (
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

func (s *MovieService) GetByID(id int) (models.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *MovieService) Create(req dto.CreateMovieRequest) (models.Movie, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{ID: id, Title: req.Title, Releaseyear: req.ReleaseYear, Duration: req.Duration}, nil
}
