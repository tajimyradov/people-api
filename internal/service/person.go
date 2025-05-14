package service

import (
	"go.uber.org/zap"
	"people-api/internal/domain"
	"people-api/internal/repository"
)

type PersonService struct {
	repo     *repository.PersonRepository
	logger   *zap.Logger
	enricher *Enricher
}

func NewPersonService(repo *repository.PersonRepository, enricher *Enricher, logger *zap.Logger) *PersonService {
	return &PersonService{
		repo:     repo,
		enricher: enricher,
		logger:   logger,
	}
}

func (s *PersonService) Create(input *domain.Person) (int, error) {
	enriched, err := s.enricher.Enrich(input.Name)
	if err != nil {
		s.logger.Error("enrich error", zap.Error(err))
		return 0, err
	}

	input.Age = enriched.Age
	input.Gender = enriched.Gender
	input.Nationality = enriched.Nationality

	return s.repo.Create(input)
}

func (s *PersonService) List(name, surname, nationality string, page, limit int) ([]domain.Person, error) {
	result, err := s.repo.List(name, surname, nationality, page, limit)
	if err != nil {
		s.logger.Error("list error", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (s *PersonService) GetByID(id string) (*domain.Person, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("get by id error", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (s *PersonService) Update(id string, person *domain.Person) error {
	err := s.repo.Update(id, person)
	if err != nil {
		s.logger.Error("update error", zap.Error(err))
		return err
	}
	return nil
}

func (s *PersonService) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("delete error", zap.Error(err))
		return err
	}
	return nil
}
