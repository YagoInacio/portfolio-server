package usecases

import (
	"errors"

	"github.com/yagoinacio/portfolio-server/internal/projects/domain/repositories"
)

type AlterProjectDisplayInput struct {
	ID      string `json:"id" example:"6568ed9d59e4487ccb66c757"`
	Display bool   `json:"display" example:"true"`
}

type AlterProjectDisplayOutput struct {
	ID          string   `json:"id" example:"6568ed9d59e4487ccb66c757"`
	Title       string   `json:"title" example:"Portfolio App"`
	Description []string `json:"description" example:"lets the user register projects"`
	Techs       []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
	Display     bool     `json:"display" example:"true"`
}

type AlterProjectDisplayUseCase struct {
	ProjectsRepository repositories.ProjectsRepositoryInterface
}

func NewAlterProjectDisplayUseCase(projectsRepository repositories.ProjectsRepositoryInterface) *AlterProjectDisplayUseCase {
	return &AlterProjectDisplayUseCase{
		ProjectsRepository: projectsRepository,
	}
}

func (u *AlterProjectDisplayUseCase) Execute(input AlterProjectDisplayInput) (*AlterProjectDisplayOutput, error) {
	project, err := u.ProjectsRepository.FindById(input.ID)
	if err != nil {
		return nil, errors.New("project not found")
	}

	project.Display = input.Display
	err = u.ProjectsRepository.Update(&project)
	if err != nil {
		return nil, err
	}

	var techs []string

	for _, tech := range project.Techs {
		techs = append(techs, tech.ID.Hex())
	}

	return &AlterProjectDisplayOutput{
		ID:          project.ID.Hex(),
		Title:       project.Title,
		Description: project.Description,
		Techs:       techs,
		Display:     project.Display,
	}, nil
}
