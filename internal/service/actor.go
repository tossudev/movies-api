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
