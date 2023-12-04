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

func TestShouldBeAbleToAlterProjectDisplay(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	alterProjectDisplayUseCase := NewAlterProjectDisplayUseCase(projectsRepo)

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

	result, _ := alterProjectDisplayUseCase.Execute(AlterProjectDisplayInput{
		ID:      proj1.ID.Hex(),
		Display: false,
	})

	var altered domain.Project

	for _, project := range repositories.Projects {
		if project.ID.Hex() == result.ID {
			altered = project
			break
		}
	}

	assert.Equal(t, result.Display, false)
	assert.Equal(t, altered.Display, false)
}

func TestShouldNotBeAbleToAlterNonExistentProject(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	alterProjectDisplayUseCase := NewAlterProjectDisplayUseCase(projectsRepo)

	_, err := alterProjectDisplayUseCase.Execute(AlterProjectDisplayInput{
		ID:      primitive.NewObjectID().Hex(),
		Display: false,
	})

	assert.EqualError(t, err, "project not found")
}
