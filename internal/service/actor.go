package service

import (
	"errors"
	"fmt"
	"movies-api/internal/dto"
	"movies-api/internal/models"
	"movies-api/internal/repository"
)

var ErrAssociatedMoviesActor = errors.New("actor has associated movies")

type ActorService struct {
	repo *repository.ActorRepository
}

func NewActorService(repo *repository.ActorRepository) *ActorService {
	return &ActorService{
		repo: repo,
	}
}

func (s *ActorService) GetAll(page, size int) ([]models.Actor, error) {
	return s.repo.GetAll(page, size)
}

func (s *ActorService) GetByID(id int) (models.Actor, error) {
	return s.repo.GetByID(id)
}

func (s *ActorService) Create(req dto.CreateActorRequest) (models.Actor, error) {
	id, err := s.repo.Create(req)
	if err != nil {
		return models.Actor{}, err
	}

	return models.Actor{ID: id, Name: req.Name, BirthDate: req.BirthDate}, nil
}

func (s *ActorService) Update(id int, req dto.UpdateActorRequest) error {
	return s.repo.Update(id, req)
}

func (s *ActorService) Delete(id int, force bool) error {
	if !force {
		if movies, err := s.repo.GetMovies(id); err != nil {
			return fmt.Errorf("query deleted actor's movies: %w", err)
		} else if len(movies) > 0 {
			return ErrAssociatedMoviesActor
		}
	}
	return s.repo.Delete(id)
}
