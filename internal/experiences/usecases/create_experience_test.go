package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
)

func TestShouldBeAbleToCreateExperience(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	createExperienceUseCase := NewCreateExperienceUseCase(xpRepo)

	input := CreateExperienceInput{
		Position: "POS_TEST",
		Company:  "COMPANY_TEST",
		Logo:     "LOGO_TEST.SVG",
		Start:    "01/2020",
		End:      "01/2022",
		Summary:  []string{"SUM_1", "SUM_2", "SUM_3"},
		Techs:    []string{"6568ed9d59e4487ccb66c757", "6568ee3e7bbf5a6160f444f4"},
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
	createExperienceUseCase := NewCreateExperienceUseCase(xpRepo)

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

	assert.EqualError(t, err, "the provided hex string is not a valid ObjectID")
}
