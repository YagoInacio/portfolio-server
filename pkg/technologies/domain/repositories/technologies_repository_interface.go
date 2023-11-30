package repositories

import domain "github.com/yagoinacio/portfolio-server/pkg/technologies/domain/entities"

type TechnologiesRepositoryInterface interface {
	Save(technology *domain.Technology) error
	FindDisplayable() ([]domain.Technology, error)
	FindByName(name string) (domain.Technology, error)
	SetDisplay(id string, display bool) error
}
