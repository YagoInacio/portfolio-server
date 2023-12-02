package repositories

import (
	"errors"

	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
)

type InMemoryExperiencesRepository struct{}

func NewInMemoryExperiencesRepository() *InMemoryExperiencesRepository {
	return &InMemoryExperiencesRepository{}
}

var Experiences []domain.Experience

func (repo *InMemoryExperiencesRepository) Save(experience *domain.Experience) error {
	Experiences = append(Experiences, *experience)

	return nil
}

func (repo *InMemoryExperiencesRepository) ListAll() ([]domain.Experience, error) {
	return Experiences, nil
}

func (repo *InMemoryExperiencesRepository) FindById(id string) (domain.Experience, error) {
	var experience *domain.Experience

	for _, xp := range Experiences {
		if xp.ID.Hex() == id {
			experience = &xp
			break
		}
	}

	if experience == nil {
		return domain.Experience{}, errors.New("experience not found")
	}

	return *experience, nil
}

func (repo *InMemoryExperiencesRepository) Update(experience *domain.Experience) error {
	for i, xp := range Experiences {
		if xp.ID.Hex() == experience.ID.Hex() {
			Experiences[i] = *experience
			break
		}
	}

	return nil
}
