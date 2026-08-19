package service

import (
	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

type ActorService struct {
	repo *repository.ActorRepository
}

func NewActorService(repo *repository.ActorRepository) *ActorService {
	return &ActorService{
		repo: repo,
	}
}

func (s *ActorService) GetAll() ([]models.Actor, error) {
	return s.repo.GetAll()
}

func (s *ActorService) GetByID(id int) (models.Actor, error) {
	return s.repo.GetByID(id)
}

func (s *ActorService) Create(req dto.CreateActorRequest) error {
	return s.repo.Create(req)
}

func (s *ActorService) Update(id int, req dto.UpdateActorRequest) error {
	return s.repo.Update(id, req)
}

func (s *ActorService) Delete(id int) error {
	return s.repo.Delete(id)
}
