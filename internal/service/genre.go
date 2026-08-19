package service

import (
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

func (s *GenreService) Create(genre models.Genre) error {
	return s.repo.Create(genre)
}

func (s *GenreService) Update(genre models.Genre) error {
	return s.repo.Update(genre)
}

func (s *GenreService) Delete(id int) error {
	return s.repo.Delete(id)
}
