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

func TestShouldBeAbleToCreateProject(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	createProjectUseCase := NewCreateProjectUseCase(projectsRepo, techRepo)

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

	input := CreateProjectInput{
		Title:       "TITLE_TEST",
		Description: []string{"Description_1", "Description_2", "Description_3"},
		Techs:       []string{tech1.ID.Hex(), tech2.ID.Hex()},
	}

	result, _ := createProjectUseCase.Execute(input)

	var inserted domain.Project

	for _, project := range repositories.Projects {
		if project.ID.Hex() == result.ID {
			inserted = project
			break
		}
	}

	assert.NotEqual(t, inserted, domain.Project{})
	assert.NotEmpty(t, inserted.ID)
	assert.Equal(t, len(inserted.Description), 3)
	assert.Equal(t, len(inserted.Techs), 2)
}

func TestShouldNotBeAbleToCreateProjectWithInvalidTechId(t *testing.T) {
	projectsRepo := repositories.NewInMemoryProjectsRepository()
	techRepo := tech_repositories.NewInMemoryTechnologiesRepository()
	createProjectUseCase := NewCreateProjectUseCase(projectsRepo, techRepo)

	input := CreateProjectInput{
		Title:       "TITLE_TEST",
		Description: []string{"Description_1", "Description_2", "Description_3"},
		Techs:       []string{primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex()},
	}

	_, err := createProjectUseCase.Execute(input)

	assert.EqualError(t, err, "technology not found")
}
