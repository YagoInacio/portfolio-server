package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/experiences/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

func TestShouldBeAbleToListExperiences(t *testing.T) {
	xpRepo := repositories.NewInMemoryExperiencesRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	listExperiencesUseCase := NewListExperiencesUseCase(xpRepo)

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
	xp2, _ := domain.NewExperience(
		"POS_TEST_2",
		"COMPANY_TEST_2",
		"LOGO_TEST_2.SVG",
		"01/2020",
		"07/2021",
		[]string{"SUMMARY_TEST_1", "SUMMARY_TEST_2", "SUMMARY_TEST_3"},
		[]technologies_domain.Technology{*tech2},
	)
	xpRepo.Save(xp1)
	xpRepo.Save(xp2)

	result, _ := listExperiencesUseCase.Execute()

	assert.Len(t, *result, 2)
	assert.Len(t, (*result)[0].Techs, 2)
	assert.Len(t, (*result)[1].Techs, 1)
	assert.Equal(t, (*result)[0].Period.End, "")
}
