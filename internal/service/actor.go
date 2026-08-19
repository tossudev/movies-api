package service

import (
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

func (s *ActorService) Create(actor models.Actor) error {
	return s.repo.Create(actor)
}

func (s *ActorService) Update(actor models.Actor) error {
	return s.repo.Update(actor)
}

func (s *ActorService) Delete(id int) error {
	return s.repo.Delete(id)
}
