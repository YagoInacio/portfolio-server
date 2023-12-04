package usecases

import (
	"errors"

	"github.com/yagoinacio/portfolio-server/internal/projects/domain/repositories"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"
)

type UpdateProjectInput struct {
	ID          string   `json:"id" example:"6568ed9d59e4487ccb66c757"`
	Title       string   `json:"title" example:"Portfolio App"`
	Description []string `json:"description" example:"lets the user register projects"`
	Techs       []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type UpdateProjectOutput struct {
	ID          string   `json:"id" example:"6568ed9d59e4487ccb66c757"`
	Title       string   `json:"title" example:"Portfolio App"`
	Description []string `json:"description" example:"lets the user register projects"`
	Techs       []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
	Display     bool     `json:"display" example:"true"`
}

type UpdateProjectUseCase struct {
	ProjectsRepository   repositories.ProjectsRepositoryInterface
	TechnologyRepository tech_repositories.TechnologiesRepositoryInterface
}

func NewUpdateProjectUseCase(
	projectsRepository repositories.ProjectsRepositoryInterface,
	technologyRepository tech_repositories.TechnologiesRepositoryInterface,
) *UpdateProjectUseCase {
	return &UpdateProjectUseCase{
		ProjectsRepository:   projectsRepository,
		TechnologyRepository: technologyRepository,
	}
}

func (u *UpdateProjectUseCase) Execute(input UpdateProjectInput) (*UpdateProjectOutput, error) {
	project, err := u.ProjectsRepository.FindById(input.ID)
	if err != nil {
		return nil, errors.New("project not found")
	}

	if len(input.Techs) > 0 {
		var techs []technologies_domain.Technology

		for _, tech := range input.Techs {
			technology, err := u.TechnologyRepository.FindByid(tech)
			if err != nil {
				return nil, errors.New("technology not found")
			}
			techs = append(techs, technology)
		}

		project.Techs = techs
	}

	if len(input.Description) > 0 {
		project.Description = input.Description
	}

	if input.Title != "" {
		project.Title = input.Title
	}

	err = u.ProjectsRepository.Update(&project)
	if err != nil {
		return nil, err
	}

	var updatedTechs []string

	for _, tech := range project.Techs {
		updatedTechs = append(updatedTechs, tech.ID.Hex())
	}

	return &UpdateProjectOutput{
		ID:          project.ID.Hex(),
		Title:       project.Title,
		Description: project.Description,
		Techs:       updatedTechs,
		Display:     project.Display,
	}, nil
}
