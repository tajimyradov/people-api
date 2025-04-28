package service

import (
	"people-api/internal/domain"
	"people-api/internal/repository"
)

type PersonService struct {
	repo     *repository.PersonRepository
	enricher *Enricher
}

func NewPersonService(repo *repository.PersonRepository, enricher *Enricher) *PersonService {
	return &PersonService{repo: repo, enricher: enricher}
}

func (s *PersonService) Create(input *domain.Person) (int, error) {
	enriched, err := s.enricher.Enrich(input.Name)
	if err != nil {
		return 0, err
	}

	input.Age = enriched.Age
	input.Gender = enriched.Gender
	input.Nationality = enriched.Nationality

	return s.repo.Create(input)
}

func (s *PersonService) List(name, surname, nationality string, page, limit int) ([]domain.Person, error) {
	return s.repo.List(name, surname, nationality, page, limit)
}

func (s *PersonService) GetByID(id string) (*domain.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) Update(id string, person *domain.Person) error {
	return s.repo.Update(id, person)
}

func (s *PersonService) Delete(id string) error {
	return s.repo.Delete(id)
}
