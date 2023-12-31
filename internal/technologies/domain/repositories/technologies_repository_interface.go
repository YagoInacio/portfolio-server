package repositories

import domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"

type TechnologiesRepositoryInterface interface {
	Save(technology *domain.Technology) error
	FindDisplayable() ([]domain.Technology, error)
	FindByName(name string) (domain.Technology, error)
	FindByid(id string) (domain.Technology, error)
	SetDisplay(id string, display bool) error
}
