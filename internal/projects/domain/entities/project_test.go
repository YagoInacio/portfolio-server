package projects_domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
)

func TestShouldBeAbleTOCreateProject(t *testing.T) {
	tech1, _ := technologies_domain.NewTechnology("TECH_TEST_1", "TECH_1.PNG", true)
	tech2, _ := technologies_domain.NewTechnology("TECH_TEST_2", "TECH_2.PNG", true)

	project, _ := NewProject(
		"TITLE_TEST",
		[]string{"DESCRIPTION_TEST_1", "DESCRIPTION_TEST_2", "DESCRIPTION_TEST_3"},
		[]technologies_domain.Technology{*tech1, *tech2},
	)

	assert.NotEmpty(t, project.ID)
	assert.Equal(t, len(project.Description), 3)
	assert.Equal(t, len(project.Techs), 2)
}

func TestTitleShouldNotBeEmpty(t *testing.T) {
	project := Project{}

	assert.EqualError(t, project.Validate(), "title is required")
}

func TestDescriptionShouldNotBeEmpty(t *testing.T) {
	project := Project{
		Title: "TITLE_TEST",
	}

	assert.EqualError(t, project.Validate(), "description is required")
}

func TestTechsShouldNotBeEmpty(t *testing.T) {
	project := Project{
		Title:       "TITLE_TEST",
		Description: []string{"SUM_1", "SUM_2"},
	}

	assert.EqualError(t, project.Validate(), "techs are required")
}
