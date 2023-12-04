package usecases

import (
	"errors"

	domain "github.com/yagoinacio/portfolio-server/internal/projects/domain/entities"
	"github.com/yagoinacio/portfolio-server/internal/projects/domain/repositories"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"
)

type CreateProjectInput struct {
	Title       string   `json:"title" example:"Portfolio App"`
	Description []string `json:"description" example:"lets the user register projects"`
	Techs       []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type CreateProjectOutput struct {
	ID          string   `json:"id" example:"6568ee3e7bbf5a6160f444f4"`
	Title       string   `json:"title" example:"Portfolio App"`
	Description []string `json:"description" example:"lets the user register projects"`
	Techs       []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
	Display     bool     `json:"display" example:"true"`
}

type CreateProjectUseCase struct {
	ProjectsRepository   repositories.ProjectsRepositoryInterface
	TechnologyRepository tech_repositories.TechnologiesRepositoryInterface
}

func NewCreateProjectUseCase(
	projectsRepository repositories.ProjectsRepositoryInterface,
	technologyRepository tech_repositories.TechnologiesRepositoryInterface,
) *CreateProjectUseCase {
	return &CreateProjectUseCase{
		ProjectsRepository:   projectsRepository,
		TechnologyRepository: technologyRepository,
	}
}

func (u *CreateProjectUseCase) Execute(input CreateProjectInput) (*CreateProjectOutput, error) {
	var techs []technologies_domain.Technology

	for _, tech := range input.Techs {
		technology, err := u.TechnologyRepository.FindByid(tech)
		if err != nil {
			return nil, errors.New("technology not found")
		}
		techs = append(techs, technology)
	}

	project, err := domain.NewProject(
		input.Title,
		input.Description,
		techs,
	)
	if err != nil {
		return nil, err
	}

	err = u.ProjectsRepository.Save(project)
	if err != nil {
		return nil, err
	}

	var insertedTechs []string

	for _, tech := range project.Techs {
		insertedTechs = append(insertedTechs, tech.ID.Hex())
	}

	return &CreateProjectOutput{
		ID:          project.ID.Hex(),
		Title:       project.Title,
		Description: project.Description,
		Display:     project.Display,
		Techs:       insertedTechs,
	}, nil
}
