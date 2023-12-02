package usecases

import (
	experiences_domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	"github.com/yagoinacio/portfolio-server/internal/experiences/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateExperienceInput struct {
	Position string   `json:"position" example:"Developer"`
	Company  string   `json:"company" example:"Metacortex"`
	Logo     string   `json:"logo" example:"metacortex.svg"`
	Start    string   `json:"start" example:"01/2023"`
	End      string   `json:"end" example:"10/2023"`
	Summary  []string `json:"summary" example:"Developed 3 apps"`
	Techs    []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type Period struct {
	Start string `json:"start" example:"01/2023"`
	End   string `json:"end" example:"10/2023"`
}

type CreateExperienceOutput struct {
	ID       string   `json:"id" example:"6568ee3e7bbf5a6160f444f4"`
	Position string   `json:"position" example:"Developer"`
	Company  string   `json:"company" example:"Metacortex"`
	Logo     string   `json:"logo" example:"metacortex.svg"`
	Period   Period   `json:"period"`
	Summary  []string `json:"summary" example:"Developed 3 apps"`
	Techs    []string `json:"techs" example:"6568ed9d59e4487ccb66c757"`
}

type CreateExperienceUseCase struct {
	ExperiencesRepository repositories.ExperiencesRepositoryInterface
}

func NewCreateExperienceUseCase(experiencesRepository repositories.ExperiencesRepositoryInterface) *CreateExperienceUseCase {
	return &CreateExperienceUseCase{
		ExperiencesRepository: experiencesRepository,
	}
}

func (u *CreateExperienceUseCase) Execute(input CreateExperienceInput) (*CreateExperienceOutput, error) {
	var techs []primitive.ObjectID

	for _, tech := range input.Techs {
		techId, err := primitive.ObjectIDFromHex(tech)
		if err != nil {
			return nil, err
		}
		techs = append(techs, techId)
	}

	xp, err := experiences_domain.NewExperience(
		input.Position,
		input.Company,
		input.Logo,
		input.Start,
		input.End,
		input.Summary,
		techs,
	)
	if err != nil {
		return nil, err
	}

	err = u.ExperiencesRepository.Save(xp)
	if err != nil {
		return nil, err
	}

	var insertedTechs []string

	for _, tech := range xp.Techs {
		insertedTechs = append(insertedTechs, tech.Hex())
	}

	return &CreateExperienceOutput{
		ID:       xp.ID.Hex(),
		Position: xp.Position,
		Company:  xp.Company,
		Logo:     xp.Logo,
		Period: Period{
			Start: xp.Period.Start,
			End:   xp.Period.End,
		},
		Summary: xp.Summary,
		Techs:   insertedTechs,
	}, nil
}
