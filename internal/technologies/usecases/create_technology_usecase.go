package usecases

import (
	"errors"

	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"
)

type TechnologyInput struct {
	Name    string `json:"name" example:"Go"`      // tech name
	Src     string `json:"src" example:"go.png"`   // tech icon file name
	Display bool   `json:"display" example:"true"` // whether it will be displayed on the skills page
}

type TechnologyOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Src     string `json:"src"`
	Display bool   `json:"display"`
}

type CreateTechnologyUseCase struct {
	TechnologyRepository repositories.TechnologiesRepositoryInterface
}

func NewCreateTechnologyUseCase(technologyRepository repositories.TechnologiesRepositoryInterface) *CreateTechnologyUseCase {
	return &CreateTechnologyUseCase{
		TechnologyRepository: technologyRepository,
	}
}

func (c *CreateTechnologyUseCase) Execute(input TechnologyInput) (*TechnologyOutput, error) {
	_, err := c.TechnologyRepository.FindByName(input.Name)
	if err == nil {
		return nil, errors.New("technology already exists")
	} else {
		tech, err := technologies_domain.NewTechnology(input.Name, input.Src, input.Display)
		if err != nil {
			return nil, err
		}

		err = c.TechnologyRepository.Save(tech)
		if err != nil {
			return nil, err
		}

		return &TechnologyOutput{
			ID:      tech.ID.Hex(),
			Name:    tech.Name,
			Src:     tech.Src,
			Display: tech.Display,
		}, nil
	}
}
