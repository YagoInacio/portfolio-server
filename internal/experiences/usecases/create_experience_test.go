package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

func TestShouldBeAbleToCreateExperience(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	createExperienceUseCase := NewCreateExperienceUseCase(xpRepo, techRepo)

	tech1, _ := technologies_domain.NewTechnology(
		"TECH_1",
		"TECH_1.png",
		true,
	)
	tech2, _ := technologies_domain.NewTechnology(
		"TECH_2",
		"TECH_2.png",
		true,
	)

	techRepo.Save(tech1)
	techRepo.Save(tech2)

	input := CreateExperienceInput{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
		Start:    "01/2020",
		End:      "01/2022",
		Summary:  []string{"SUM_1", "SUM_2", "SUM_3"},
		Techs:    []string{tech1.ID.Hex(), tech2.ID.Hex()},
	}

	result, _ := createExperienceUseCase.Execute(input)

	var inserted domain.Experience

	for _, xp := range repositories.Experiences {
		if xp.ID.Hex() == result.ID {
			inserted = xp
			break
		}
	}

	assert.NotEqual(t, inserted, domain.Experience{})
	assert.NotEmpty(t, inserted.ID)
	assert.Equal(t, len(inserted.Summary), 3)
	assert.Equal(t, len(inserted.Techs), 2)
}

func TestShouldNotBeAbleToCreateExperienceWithInvalidTechId(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	createExperienceUseCase := NewCreateExperienceUseCase(xpRepo, techRepo)

	input := CreateExperienceInput{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
		Start:    "01/2020",
		End:      "01/2022",
		Summary:  []string{"SUM_1", "SUM_2", "SUM_3"},
		Techs:    []string{"6568ed9d59e4487ccb66c757", "6568ee3e5a6160f444f4"},
	}

	_, err := createExperienceUseCase.Execute(input)

	assert.EqualError(t, err, "technology not found")
}
