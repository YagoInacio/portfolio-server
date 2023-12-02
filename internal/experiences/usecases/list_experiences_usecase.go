package usecases

import (
	"github.com/yagoinacio/portfolio-server/internal/experiences/domain/repositories"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

type ListExperiencesOutput struct {
	ID       string                           `json:"id" example:"6568ee3e7bbf5a6160f444f4"`
	Position string                           `json:"position" example:"Developer"`
	Company  string                           `json:"company" example:"Metacortex"`
	Logo     string                           `json:"logo" example:"metacortex.svg"`
	Period   Period                           `json:"period"`
	Summary  []string                         `json:"summary" example:"Developed 3 apps"`
	Techs    []technologies_domain.Technology `json:"techs"`
}

type ListExperiencesUseCase struct {
	ExperiencesRepository repositories.ExperiencesRepositoryInterface
}

func NewListExperiencesUseCase(experiencesRepository repositories.ExperiencesRepositoryInterface) *ListExperiencesUseCase {
	return &ListExperiencesUseCase{
		ExperiencesRepository: experiencesRepository,
	}
}

func (u *ListExperiencesUseCase) Execute() (*[]ListExperiencesOutput, error) {
	result, err := u.ExperiencesRepository.ListAll()
	if err != nil {
		return nil, err
	}

	var outputs []ListExperiencesOutput

	for _, xp := range result {
		toAppend := ListExperiencesOutput{
			ID:       xp.ID.Hex(),
			Position: xp.Position,
			Company:  xp.Company,
			Logo:     xp.Logo,
			Period: Period{
				Start: xp.Period.Start,
				End:   xp.Period.End,
			},
			Summary: xp.Summary,
			Techs:   xp.Techs,
		}
		outputs = append(outputs, toAppend)
	}

	return &outputs, nil
}
