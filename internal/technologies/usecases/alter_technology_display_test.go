package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
	domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldBeAbleToAlterTechnologyDisplay(t *testing.T) {
	techRepo := repositories.NewInMemoryTechnologiesRepository()
	createTechUseCase := NewCreateTechnologyUseCase(techRepo)
	alterTech := NewAlterTechnologyDisplayUseCase(techRepo)

	input := TechnologyInput{
		Name:    "NAME_TEST",
		Src:     "TEST.PNG",
		Display: true,
	}

	createResult, _ := createTechUseCase.Execute(input)
	result, _ := alterTech.Execute(AlterTechnologyDisplayInput{
		ID:      createResult.ID,
		Display: false,
	})

	var altered domain.Technology

	for _, tech := range repositories.Technologies {
		if tech.ID.Hex() == result.ID {
			altered = tech
			break
		}
	}

	assert.Equal(t, result.Display, false)
	assert.Equal(t, altered.Display, false)
}

func TestShouldNotBeAbleToAlterNonExistentTechnology(t *testing.T) {
	techRepo := repositories.NewInMemoryTechnologiesRepository()
	alterTech := NewAlterTechnologyDisplayUseCase(techRepo)

	_, err := alterTech.Execute(AlterTechnologyDisplayInput{
		ID:      primitive.NewObjectID().Hex(),
		Display: false,
	})

	expectedErr := errors.New("technology not found")
	assert.EqualError(t, err, expectedErr.Error(), "technology not found")
}
