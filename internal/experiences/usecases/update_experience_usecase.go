package usecases

import (
	"errors"

	"github.com/yagoinacio/portfolio-server/internal/experiences/domain/repositories"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/domain/repositories"
)

type UpdateExperienceInput struct {
	ID       string   `json:"id" example:"6568ed9d59e4487ccb66c757"`
	Position string   `json:"position" example:"Developer"`
	Company  string   `json:"company" example:"Metacortex"`
	Logo     string   `json:"logo" example:"metacortex.svg"`
	Start    string   `json:"start" example:"01/2023"`
	End      string   `json:"end" example:"10/2023"`
	Summary  []string `json:"summary" example:"Developed 3 apps"`
	Techs    []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type UpdateExperienceOutput struct {
	ID       string   `json:"id" example:"6568ee3e7bbf5a6160f444f4"`
	Position string   `json:"position" example:"Developer"`
	Company  string   `json:"company" example:"Metacortex"`
	Logo     string   `json:"logo" example:"metacortex.svg"`
	Period   Period   `json:"period"`
	Summary  []string `json:"summary" example:"Developed 3 apps"`
	Techs    []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type UpdateExperienceUseCase struct {
	ExperiencesRepository repositories.ExperiencesRepositoryInterface
	TechnologyRepository  tech_repositories.TechnologiesRepositoryInterface
}

func NewUpdateExperienceUseCase(
	experiencesRepository repositories.ExperiencesRepositoryInterface,
	technologyRepository tech_repositories.TechnologiesRepositoryInterface,
) *UpdateExperienceUseCase {
	return &UpdateExperienceUseCase{
		ExperiencesRepository: experiencesRepository,
		TechnologyRepository:  technologyRepository,
	}
}

func (u *UpdateExperienceUseCase) Execute(input UpdateExperienceInput) (*UpdateExperienceOutput, error) {
	experience, err := u.ExperiencesRepository.FindById(input.ID)
	if err != nil {
		return nil, errors.New("experience not found")
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

		experience.Techs = techs
	}

	if len(input.Summary) > 0 {
		experience.Summary = input.Summary
	}

	if input.Position != "" {
		experience.Position = input.Position
	}

	if input.Company != "" {
		experience.Company = input.Company
	}

	if input.Logo != "" {
		experience.Logo = input.Logo
	}

	if input.Start != "" {
		experience.Period.Start = input.Start
	}
	if input.End != "" {
		experience.Period.End = input.End
	}

	err = u.ExperiencesRepository.Update(&experience)
	if err != nil {
		return nil, err
	}

	var updatedTechs []string

	for _, tech := range experience.Techs {
		updatedTechs = append(updatedTechs, tech.ID.Hex())
	}

	return &UpdateExperienceOutput{
		ID:       experience.ID.Hex(),
		Position: experience.Position,
		Company:  experience.Company,
		Logo:     experience.Logo,
		Period: Period{
			Start: experience.Period.Start,
			End:   experience.Period.End,
		},
		Summary: experience.Summary,
		Techs:   updatedTechs,
	}, nil
}
