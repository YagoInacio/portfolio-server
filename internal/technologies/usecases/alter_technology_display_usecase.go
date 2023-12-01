package usecases

import (
	"errors"

	"github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"
)

type AlterTechnologyDisplayInput struct {
	ID      string `json:"id"`
	Display bool   `json:"display"`
}

type AlterTechnologyDisplayOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Src     string `json:"src"`
	Display bool   `json:"display"`
}

type AlterTechnologyDisplayUseCase struct {
	TechnologyRepository repositories.TechnologiesRepositoryInterface
}

func NewAlterTechnologyDisplayUseCase(technologyRepository repositories.TechnologiesRepositoryInterface) *AlterTechnologyDisplayUseCase {
	return &AlterTechnologyDisplayUseCase{
		TechnologyRepository: technologyRepository,
	}
}

func (a *AlterTechnologyDisplayUseCase) Execute(input AlterTechnologyDisplayInput) (*AlterTechnologyDisplayOutput, error) {
	tech, err := a.TechnologyRepository.FindByid(input.ID)
	if err != nil {
		return nil, errors.New("technology not found")
	}

	tech.Display = input.Display
	err = a.TechnologyRepository.SetDisplay(input.ID, input.Display)
	if err != nil {
		return nil, err
	}

	return &AlterTechnologyDisplayOutput{
		ID:      tech.ID.Hex(),
		Name:    tech.Name,
		Src:     tech.Src,
		Display: tech.Display,
	}, nil
}
