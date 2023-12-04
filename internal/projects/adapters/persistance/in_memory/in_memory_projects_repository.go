package repositories

import (
	"errors"

	domain "github.com/yagoinacio/portfolio-server/internal/projects/domain/entities"
)

type InMemoryProjectsRepository struct{}

func NewInMemoryProjectsRepository() *InMemoryProjectsRepository {
	return &InMemoryProjectsRepository{}
}

var Projects []domain.Project

func (repo *InMemoryProjectsRepository) Save(project *domain.Project) error {
	Projects = append(Projects, *project)

	return nil
}

func (r *InMemoryProjectsRepository) ListDisplayable() ([]domain.Project, error) {
	var displayable []domain.Project

	for _, project := range Projects {
		if project.Display {
			displayable = append(displayable, project)
		}
	}

	return displayable, nil
}

func (repo *InMemoryProjectsRepository) FindById(id string) (domain.Project, error) {
	var project *domain.Project

	for _, proj := range Projects {
		if proj.ID.Hex() == id {
			project = &proj
			break
		}
	}

	if project == nil {
		return domain.Project{}, errors.New("experience not found")
	}

	return *project, nil
}

func (repo *InMemoryProjectsRepository) Update(project *domain.Project) error {
	for i, proj := range Projects {
		if proj.ID.Hex() == project.ID.Hex() {
			Projects[i] = *project
			break
		}
	}

	return nil
}
