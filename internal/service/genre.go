package service

import (
	"fmt"
	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

type GenreService struct {
	repo *repository.GenreRepository
}

func NewGenreService(repo *repository.GenreRepository) *GenreService {
	return &GenreService{
		repo: repo,
	}
}

func (s *GenreService) GetAll(page, size int) ([]models.Genre, error) {
	return s.repo.GetAll(page, size)
}

func (s *GenreService) GetByID(id int) (models.Genre, error) {
	return s.repo.GetByID(id)
}

func (s *GenreService) Create(req dto.CreateGenreRequest) (models.Genre, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return models.Genre{}, err
	}

	return models.Genre{ID: id, Name: req.Name}, nil
}

func (s *GenreService) Update(id int, req dto.UpdateGenreRequest) error {
	return s.repo.Update(id, req)
}

func (s *GenreService) Delete(id int, force bool) error {
	if !force {
		if movies, err := s.repo.GetMovies(id); err != nil {
			return fmt.Errorf("query deleted genre's movies: %w", err)
		} else if len(movies) > 0 {
			return fmt.Errorf("cannot delete genre because it has %d associated movies", len(movies))
		}
	}

	return s.repo.Delete(id)
}
