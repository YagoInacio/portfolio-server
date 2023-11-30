package repositories

import (
	"errors"

	domain "github.com/yagoinacio/portfolio-server/pkg/technologies/domain/entities"
)

type InMemoryTechnologiesRepository struct{}

func NewInMemoryTechnologiesRepository() *InMemoryTechnologiesRepository {
	return &InMemoryTechnologiesRepository{}
}

var Technologies []domain.Technology

func (r *InMemoryTechnologiesRepository) Save(technology *domain.Technology) error {
	Technologies = append(Technologies, *technology)

	return nil
}

func (r *InMemoryTechnologiesRepository) FindDisplayable() ([]domain.Technology, error) {
	var displayable []domain.Technology

	for _, tech := range Technologies {
		if tech.Display {
			displayable = append(displayable, tech)
		}
	}

	return displayable, nil
}

func (r *InMemoryTechnologiesRepository) FindByName(name string) (domain.Technology, error) {
	var technology domain.Technology

	for _, tech := range Technologies {
		if tech.Name == name {
			technology = tech
			break
		}
	}

	if technology == (domain.Technology{}) {
		return technology, errors.New("technology not found")
	}

	return technology, nil
}

func (r *InMemoryTechnologiesRepository) SetDisplay(id string, display bool) error {
	for i, tech := range Technologies {
		if tech.ID.Hex() == id {
			Technologies[i].Display = display
			break
		}
	}

	return nil
}
