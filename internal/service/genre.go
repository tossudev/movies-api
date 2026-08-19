package service

import (
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

func (s *GenreService) GetAll() ([]models.Genre, error) {
	return s.repo.GetAll()
}

func (s *GenreService) GetByID(id int) (models.Genre, error) {
	return s.repo.GetByID(id)
}

func (s *GenreService) Create(req dto.CreateGenreRequest) error {
	return s.repo.Create(req)
}

func (s *GenreService) Update(id int, req dto.UpdateGenreRequest) error {
	return s.repo.Update(id, req)
}

func (s *GenreService) Delete(id int) error {
	return s.repo.Delete(id)
}
