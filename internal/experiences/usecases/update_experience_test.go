package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldBeAbleToUpdateExperience(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateExperienceUseCase := NewUpdateExperienceUseCase(xpRepo, techRepo)

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

	xp1, _ := domain.NewExperience(
		"POS_TEST",
		"COMPANY_TEST",
		"LOGO_TEST.SVG",
		"01/2023",
		"",
		[]string{"SUMMARY_TEST_1", "SUMMARY_TEST_2", "SUMMARY_TEST_3"},
		[]technologies_domain.Technology{*tech1, *tech2},
	)
	xpRepo.Save(xp1)

	result, _ := updateExperienceUseCase.Execute(UpdateExperienceInput{
		ID:       xp1.ID.Hex(),
		Position: "POS_TEST_UPDATED",
		Company:  "COMPANY_TEST_UPDATED",
		Logo:     "LOGO_TEST_UPDATED.SVG",
		Start:    "01/2020",
		End:      "01/2022",
		Summary:  []string{"SUM_1", "SUM_2", "SUM_3", "SUM_4"},
		Techs:    []string{tech1.ID.Hex()},
	})

	var updated domain.Experience

	for _, xp := range repositories.Experiences {
		if xp.ID.Hex() == result.ID {
			updated = xp
			break
		}
	}

	assert.Equal(t, len(updated.Summary), 4)
	assert.Equal(t, len(updated.Techs), 1)
	assert.Equal(t, updated.Position, "POS_TEST_UPDATED")
	assert.Equal(t, updated.Company, "COMPANY_TEST_UPDATED")
	assert.Equal(t, updated.Logo, "LOGO_TEST_UPDATED.SVG")
	assert.Equal(t, updated.Period.Start, "01/2020")
	assert.Equal(t, updated.Period.End, "01/2022")
}

func TestShouldNotBeAbleToUpdateNonExistentExperience(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateExperienceUseCase := NewUpdateExperienceUseCase(xpRepo, techRepo)

	_, err := updateExperienceUseCase.Execute(UpdateExperienceInput{
		ID:       primitive.NewObjectID().Hex(),
		Position: "POS_TEST_UPDATED",
		Company:  "COMPANY_TEST_UPDATED",
		Logo:     "LOGO_TEST_UPDATED.SVG",
		Start:    "01/2020",
		End:      "01/2022",
	})

	assert.EqualError(t, err, "experience not found")
}

func TestShouldNotBeAbleToUpdateExperienceWithInvalidTechId(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateExperienceUseCase := NewUpdateExperienceUseCase(xpRepo, techRepo)

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

	xp1, _ := domain.NewExperience(
		"POS_TEST",
		"COMPANY_TEST",
		"LOGO_TEST.SVG",
		"01/2023",
		"",
		[]string{"SUMMARY_TEST_1", "SUMMARY_TEST_2", "SUMMARY_TEST_3"},
		[]technologies_domain.Technology{*tech1, *tech2},
	)
	xpRepo.Save(xp1)

	_, err := updateExperienceUseCase.Execute(UpdateExperienceInput{
		ID:    xp1.ID.Hex(),
		Techs: []string{primitive.NewObjectID().Hex()},
	})

	assert.EqualError(t, err, "technology not found")
}
