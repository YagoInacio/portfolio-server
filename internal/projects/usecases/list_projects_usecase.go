package usecases

import (
	"github.com/yagoinacio/portfolio-server/internal/projects/domain/repositories"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

type ListProjectsOutput struct {
	ID          string                           `json:"id" example:"6568ee3e7bbf5a6160f444f4"`
	Title       string                           `json:"title" example:"App 1"`
	Description []string                         `json:"description" example:"lets the user register projects"`
	Display     bool                             `json:"display" example:"true"`
	Techs       []technologies_domain.Technology `json:"techs"`
}

type ListProjectsUseCase struct {
	ProjectsRepository repositories.ProjectsRepositoryInterface
}

func NewListProjectsUseCase(projectsRepository repositories.ProjectsRepositoryInterface) *ListProjectsUseCase {
	return &ListProjectsUseCase{
		ProjectsRepository: projectsRepository,
	}
}

func (u *ListProjectsUseCase) Execute() (*[]ListProjectsOutput, error) {
	result, err := u.ProjectsRepository.ListDisplayable()
	if err != nil {
		return nil, err
	}

	var outputs []ListProjectsOutput

	for _, project := range result {
		toAppend := ListProjectsOutput{
			ID:          project.ID.Hex(),
			Title:       project.Title,
			Description: project.Description,
			Techs:       project.Techs,
			Display:     project.Display,
		}

		outputs = append(outputs, toAppend)
	}

	return &outputs, nil
}
