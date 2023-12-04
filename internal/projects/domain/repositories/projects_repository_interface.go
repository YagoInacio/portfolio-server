package repositories

import domain "github.com/yagoinacio/portfolio-server/internal/projects/domain/entities"

type ProjectsRepositoryInterface interface {
	Save(project *domain.Project) error
	ListDisplayable() ([]domain.Project, error)
	FindById(id string) (domain.Project, error)
	Update(project *domain.Project) error
}
