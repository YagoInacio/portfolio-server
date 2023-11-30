package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/pkg/technologies/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/pkg/technologies/domain/entities"
)

func TestShouldBeAbleToCreateTechnology(t *testing.T) {
	techRepo := repositories.NewInMemoryTechnologiesRepository()
	createTechUseCase := NewCreateTechnologyUseCase(techRepo)

	input := TechnologyInput{
		Name:    "NAME_TEST",
		Src:     "TEST.PNG",
		Display: true,
	}

	result, _ := createTechUseCase.Execute(input)

	var inserted domain.Technology

	for _, tech := range repositories.Technologies {
		if tech.ID.Hex() == result.ID {
			inserted = tech
			break
		}
	}

	assert.NotEqual(t, inserted, domain.Technology{})
	assert.NotEmpty(t, inserted.ID)
	assert.Equal(t, inserted.Name, "NAME_TEST")
	assert.Equal(t, inserted.Src, "TEST.PNG")
	assert.Equal(t, inserted.Display, true)
}

func TestShouldNotBeAbleToCreateExistingTechnology(t *testing.T) {
	techRepo := repositories.NewInMemoryTechnologiesRepository()
	createTechUseCase := NewCreateTechnologyUseCase(techRepo)

	input := TechnologyInput{
		Name:    "NAME_TEST",
		Src:     "TEST.PNG",
		Display: true,
	}

	createTechUseCase.Execute(input)
	_, err := createTechUseCase.Execute(input)

	expectedErr := errors.New("technology already exists")
	assert.EqualError(t, err, expectedErr.Error(), "technology already exists")
}
