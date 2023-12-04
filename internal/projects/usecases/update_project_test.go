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

func TestShouldBeAbleToUpdateProject(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateProjectUseCase := NewUpdateProjectUseCase(projectsRepo, techRepo)

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
	projectsRepo.Save(proj1)

	result, _ := updateProjectUseCase.Execute(UpdateProjectInput{
		ID:          proj1.ID.Hex(),
		Title:       "UPDATED_TITLE_TEST_1",
		Description: []string{"UPDATED_DESC_TEST_1", "UPDATED_DESC_TEST_2", "UPDATED_DESC_TEST_3", "UPDATED_DESC_TEST_4"},
		Techs:       []string{tech1.ID.Hex()},
	})

	var updated domain.Project

	for _, project := range repositories.Projects {
		if project.ID.Hex() == result.ID {
			updated = project
			break
		}
	}

	assert.Equal(t, len(updated.Description), 4)
	assert.Equal(t, len(updated.Techs), 1)
	assert.Equal(t, updated.Title, "UPDATED_TITLE_TEST_1")
}

func TestShouldNotBeAbleToUpdateNonExistentProject(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateProjectUseCase := NewUpdateProjectUseCase(projectsRepo, techRepo)

	_, err := updateProjectUseCase.Execute(UpdateProjectInput{
		ID: primitive.NewObjectID().Hex(),
	})

	assert.EqualError(t, err, "project not found")
}

func TestShouldNotBeAbleToUpdateProjectWithInvalidTechId(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	updateProjectUseCase := NewUpdateProjectUseCase(projectsRepo, techRepo)

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
	projectsRepo.Save(proj1)

	_, err := updateProjectUseCase.Execute(UpdateProjectInput{
		ID:    proj1.ID.Hex(),
		Techs: []string{primitive.NewObjectID().Hex()},
	})

	assert.EqualError(t, err, "technology not found")
}
