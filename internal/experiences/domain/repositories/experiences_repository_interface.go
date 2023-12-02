package repositories

import domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"

type ExperiencesRepositoryInterface interface {
	Save(experience *domain.Experience) error
	ListAll() ([]domain.Experience, error)
	FindById(id string) (domain.Experience, error)
	Update(experience *domain.Experience) error
}
