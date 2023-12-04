package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/projects/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/projects/domain/entities"
	tech_repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldBeAbleToListProjects(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	listProjectsUseCase := NewListProjectsUseCase(projectsRepo)

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

	proj1, _ := domain.NewProject(
		"TITLE_TEST_1",
		[]string{"DESCRIPTION_TEST_1", "DESCRIPTION_TEST_2", "DESCRIPTION_TEST_3"},
		[]technologies_domain.Technology{*tech1, *tech2},
	)
	proj2 := domain.Project{
		ID:          primitive.NewObjectID(),
		Title:       "TITLE_TEST_2",
		Description: []string{"DESCRIPTION_TEST_1", "DESCRIPTION_TEST_2", "DESCRIPTION_TEST_3"},
		Techs:       []technologies_domain.Technology{*tech1},
		Display:     false,
	}
	proj3, _ := domain.NewProject(
		"TITLE_TEST_3",
		[]string{"DESCRIPTION_TEST_1", "DESCRIPTION_TEST_2", "DESCRIPTION_TEST_3"},
		[]technologies_domain.Technology{*tech2},
	)

	projectsRepo.Save(proj1)
	projectsRepo.Save(&proj2)
	projectsRepo.Save(proj3)

	result, _ := listProjectsUseCase.Execute()

	assert.Len(t, *result, 2)
	assert.Len(t, (*result)[0].Techs, 2)
	assert.Len(t, (*result)[1].Techs, 1)
	assert.Equal(t, (*result)[0].Display, true)
	assert.Equal(t, (*result)[1].Display, true)
}
