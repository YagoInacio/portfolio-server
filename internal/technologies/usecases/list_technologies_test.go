package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	repositories "github.com/yagoinacio/portfolio-server/internal/technologies/adapters/persistance/in_memory"
)

func TestShouldBeAbleToListDisplayableTechnologies(t *testing.T) {
	techRepo := repositories.NewInMemoryTechnologiesRepository()
	createTechUseCase := NewCreateTechnologyUseCase(techRepo)
	listTechsUseCase := NewListTechnologiesUseCase(techRepo)

	createTechUseCase.Execute(TechnologyInput{
		Name:    "NAME_TEST_1",
		Src:     "TEST.PNG",
		Display: true,
	})
	createTechUseCase.Execute(TechnologyInput{
		Name:    "NAME_TEST_2",
		Src:     "TEST.PNG",
		Display: false,
	})
	createTechUseCase.Execute(TechnologyInput{
		Name:    "NAME_TEST_3",
		Src:     "TEST.PNG",
		Display: true,
	})

	result, _ := listTechsUseCase.Execute()

	assert.Len(t, *result, 2)
	assert.Equal(t, (*result)[0].Display, true)
	assert.Equal(t, (*result)[1].Display, true)
}
