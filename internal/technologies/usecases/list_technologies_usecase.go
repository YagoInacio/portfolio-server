package usecases

import "github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"

type ListTechnologiesOutput struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Src     string `json:"src"`
	Display bool   `json:"display"`
}

type ListTechnologiesUseCase struct {
	TechnologyRepository repositories.TechnologiesRepositoryInterface
}

func NewListTechnologiesUseCase(technologyRepository repositories.TechnologiesRepositoryInterface) *ListTechnologiesUseCase {
	return &ListTechnologiesUseCase{
		TechnologyRepository: technologyRepository,
	}
}

func (l *ListTechnologiesUseCase) Execute() (*[]ListTechnologiesOutput, error) {
	result, err := l.TechnologyRepository.FindDisplayable()
	if err != nil {
		return nil, err
	}

	var outputs []ListTechnologiesOutput

	for _, tech := range result {
		toAppend := ListTechnologiesOutput{
			ID:      tech.ID.Hex(),
			Name:    tech.Name,
			Src:     tech.Src,
			Display: tech.Display,
		}
		outputs = append(outputs, toAppend)
	}

	return &outputs, nil
}
